package model

// from, to = id, type
// CreatedAt, Expiry = unixtimestamp
type Notification struct {
	Title     *string `json: "title"`
	Text      *string `json: "text"`
	From      *Communicator
	To        *Communicator
	CreatedAt *int `json: "created_at"`
	Expiry    *int `json: "expiry"`
}

type Notifications []Notification
