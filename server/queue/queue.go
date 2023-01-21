package queue

type Message struct {
	Id        string `json:"id"`
	Data      []byte `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

type AccountDeletionMessage struct {
	AccountId string `json:"accountId"`
}

type NotificationMessage struct {
	AccountEmail string `json:"acountEmail"`
}

type Queue interface {
	Push(message Message) error
	Pull() ([]Message, error)
}
