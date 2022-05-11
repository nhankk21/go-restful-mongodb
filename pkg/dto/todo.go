package dto

type Todo struct {
	ID      int64  `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Content string `json:"content" bson:"content"`
	Status  string `json:"status" bson:"status"`
}
