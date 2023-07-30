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

	client := asynq.NewClient(opt)

	t1, err := task.NewEmailNotifyTask(1, 0)
	if err != nil {
		log.Fatal(err)
	}

	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)

	t2, err := task.NewEmailNotifyTask(2, 0)
	if err != nil {
		log.Fatal(err)
	}

	info, err = client.Enqueue(t2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" [*] Successfully enqueued task: %+v", info)
}
