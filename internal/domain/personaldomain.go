package domain

//is in microservice
type Notification struct {
	UserID        int
	Target        string //which topic
	Type          string
	ReplyerUserID string
}

//will be in microservice
type Chat struct {
	OwnerID      int //TODO maybe make chat as common object to each users, not as own copy of each user
	TargetUserID string
	Messages     []*Message
}
type Message struct {
	ID           string //TODO yeah, rather make it common
	UserID       int
	TargetUserID string
}
