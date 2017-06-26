package db

type Status struct {
	Id      string `bson:"_id,omitempty"`
	Name    string
	BoardId string
}
