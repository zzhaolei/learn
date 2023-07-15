package main

import (
	"encoding/json"
	"net/http"

	"golang/pipeline-manager/services/http/model"
	"golang/pipeline-manager/utils/kafka"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	app := gin.Default()
	router := app.Group("/")
	producer := kafka.NewProducer()

	router.POST("/tasks", func(ctx *gin.Context) {
		var (
			err error
			msg model.Message
		)

		if err = ctx.ShouldBindJSON(&msg); err != nil {
			log.Error().Str("err", err.Error()).Msg("bind json to struct fail")
			ctx.Status(http.StatusBadRequest)
			return
		}

		value, _ := json.Marshal(msg)
		err = producer.SendMessage(ctx.Request.Context(), []byte("test-key"), value)
		if err != nil {
			log.Error().Str("err", err.Error()).Msg("send message fail")
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.Status(http.StatusOK)
	})

	_ = app.Run(":8080")
}
