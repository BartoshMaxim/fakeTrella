package db

type Task struct {
	Id       string `bson:"_id,omitempty"`
	Message  string
	StatusId string
}
