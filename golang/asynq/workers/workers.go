package main

import (
	"log"

	"golang/asynq/task"

	"github.com/hibiken/asynq"
)

type EmailTaskPayload struct {
	UserID int
}

func main() {
	opt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
		DB:   0,
	}
	cfg := asynq.Config{
		Concurrency: 10,
	}

	srv := asynq.NewServer(opt, cfg)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeNotifyEmail, task.HandleEmailNofity)
	mux.HandleFunc(task.TypeDelayEmail, task.HandleEmailDelay)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
