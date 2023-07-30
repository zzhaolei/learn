package task

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

const (
	TypeNotifyEmail = "email:notify"
	TypeDelayEmail  = "email:delay"
)

type EmailTaskPayload struct {
	UserID int
	Num    int
}

var client = asynq.NewClient(asynq.RedisClientOpt{
	Addr: "localhost:6379",
	DB:   0,
})

func NewEmailNotifyTask(id, num int) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailTaskPayload{UserID: id, Num: num})
	if err != nil {
		return nil, errors.Wrap(err, "marshal email task payload failed")
	}
	return asynq.NewTask(TypeNotifyEmail, payload), nil
}

func NewEmailDelayTask(id, num int) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailTaskPayload{UserID: id, Num: num})
	if err != nil {
		return nil, errors.Wrap(err, "marshal email task payload failed")
	}
	return asynq.NewTask(TypeDelayEmail, payload), nil
}

func HandleEmailNofity(ctx context.Context, t *asynq.Task) error {
	var p EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Println("notify:", err)
		return err
	}
	log.Printf(" [*] Recv Notify Email to User %d Num %d", p.UserID, p.Num)

	time.Sleep(time.Second * 20)
	return nil
}

func HandleEmailDelay(ctx context.Context, t *asynq.Task) error {
	var p EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Println("delay:", err)
		return err
	}
	log.Printf(" [*] Send Delay Email to User %d Num %d", p.UserID, p.Num)
	return nil
}
