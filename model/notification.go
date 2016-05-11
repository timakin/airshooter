package model

type Notification struct {
	Title     *string `json: "title"`
	Text      *string `json: "text"`
	From      *Communicator
	To        *Communicator
	CreatedAt *int64 `json: "created_at"`
	Expiry    *int64 `json: "expiry"`
}

type Notifications []*Notification

func NewNotification(args map[string]interface{}) *Notification {
	return &Notification{
		Title:     args["title"].(*string),
		Text:      args["text"].(*string),
		From:      args["from"].(*Communicator),
		To:        args["to"].(*Communicator),
		CreatedAt: args["created_at"].(*int64),
		Expiry:    args["expiry"].(*int64),
	}
}
