package main

import (
	"image"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang/image-info/config"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	app := gin.Default()

	sizeGroup := app.Group("/_size")
	sizeGroup.GET("/*path", func(ctx *gin.Context) {
		path := ctx.Param("path")
		path = strings.TrimLeft(path, "/")

		width, height, err := readSize(path)
		if err != nil {
			log.Error().Str("path", path).Str("err", err.Error()).Send()
			ctx.JSON(http.StatusOK, []int{})
			return
		}
		ctx.JSON(http.StatusOK, []int{width, height})
	})

	sizeGroup.POST("/images", func(ctx *gin.Context) {
		var images struct {
			Urls []string `json:"urls"`
		}
		if err := ctx.ShouldBindJSON(&images); err != nil {
			log.Error().Str("err", err.Error()).Msg("bind body err")
			ctx.JSON(http.StatusOK, []map[string]any{})
			return
		}

		type _resp struct {
			Url  string `json:"url"`
			Size []int  `json:"size"`
		}

		resp := make([]_resp, 0, len(images.Urls))
		for _, v := range images.Urls {
			u, err := url.Parse(v)
			if err != nil {
				log.Error().Str("url", v).Str("err", err.Error()).Msg("url parse err")
				continue
			}
			path := u.Path
			path = strings.TrimLeft(path, "/")
			width, height, err := readSize(path)
			if err != nil {
				log.Error().Str("path", path).Str("err", err.Error()).Send()
				continue
			}
			resp = append(resp, _resp{
				Url:  v,
				Size: []int{width, height},
			})
		}
		ctx.JSON(http.StatusOK, resp)
	})

	err := app.Run(":8080")
	log.Error().Str("err", err.Error()).Msg("server shutdown by err")
}

func readSize(path string) (int, int, error) {
	fullpath := filepath.Join(config.MountPath, path)

	f, err := os.Open(fullpath)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return 0, 0, err
	}
	bounds := img.Bounds()
	return bounds.Max.X, bounds.Max.Y, nil
}
