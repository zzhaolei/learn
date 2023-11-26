package model

type User struct {
	Id       uint64 `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Gender   int    `json:"gender" gorm:"column:gender"`
}

func (u *User) TableName() string {
	return "user"
}
