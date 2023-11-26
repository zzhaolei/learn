package user

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"simple-user-web/config"
	"simple-user-web/user/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type User struct {
	cfg config.Config
}

func New(config config.Config) *User {
	return &User{config}
}

// Get 获取指定用户的信息
func (u *User) Get(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "param invalid",
		})
		return
	}

	var detail Detail
	err = u.cfg.DB.Model(&model.User{}).Where("id = $1", id).First(&detail).Error
	if err != nil || detail.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "not found user",
		})
		return
	}
	c.JSON(http.StatusOK, detail)
}

// Create 创建用户
func (u *User) Create(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "params invalid",
		})
		return
	}

	var count int64
	if err := u.cfg.DB.Model(&model.User{}).
		Where("username = $1", req.Username).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("unkown err: %s", err),
		})
		return
	}
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "username exists",
		})
		return
	}

	// password md5
	req.Password = Md5String(req.Password, u.cfg.Salt)

	// create user
	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Gender:   req.Gender,
	}
	err := u.cfg.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&user).Error
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("create user err: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": user.Id,
	})
}

// Update 更新用户信息
func (u *User) Update(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "param invalid",
		})
		return
	}

	var count int64
	err = u.cfg.DB.Model(&model.User{}).Where("id = $1", id).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("unkown err: %s", err),
		})
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "param invalid",
		})
		return
	}

	var req UpdateReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "param invalid",
		})
		return
	}

	// md5
	req.Password = Md5String(req.Password, u.cfg.Salt)

	user := model.User{
		Id:       uint64(id),
		Username: req.Username,
		Password: req.Password,
		Gender:   req.Gender,
	}
	err = u.cfg.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Updates(&user).Error
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("update user err: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// Delete 删除用户
func (u *User) Delete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "param invalid",
		})
		return
	}

	err = u.cfg.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&model.User{}, id).Error
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("delete user err: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Md5String(s, salt string) string {
	m := md5.New()
	_, _ = io.WriteString(m, salt)
	_, _ = io.WriteString(m, s)
	return fmt.Sprintf("%x", m.Sum(nil))
}
