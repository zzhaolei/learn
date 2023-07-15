package model

type TaskType string

const (
	Txt2Img TaskType = "TXT2IMG"
	Img2Img TaskType = "IMG2IMG"
)

type Message struct {
	Tasks []TaskMessage `json:"tasks"`
	Hash  string        `json:"hash,omitempty"`
}

type TaskMessage struct {
	TaskType    TaskType `json:"task_type"`
	StartTime   int      `json:"start_time"`
	CallBackUrl string   `json:"callback_url,omitempty"`
	Body        Task     `json:"body"`
}

type Task struct {
	Prompt string `json:"prompt"`
}
