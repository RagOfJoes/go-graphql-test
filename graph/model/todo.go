package model

// Todo is modified to allow for resolvers
// to lookup User type
type Todo struct {
	ID     string `json:"_id" bson:"_id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}
