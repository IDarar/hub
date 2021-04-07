package domain

import "time"

type Role string

type User struct {
	//TODO last proposition opened
	ID           uint
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
	LastVisitAt  time.Time
	Role         UserRole //admin, SuperModerator, ContentModerator, ForumModerator

	EncryptedPassword string
	OnlineChan        chan int
	IsOnline          bool

	UserListID int
	UserLists  *UserLists

	Articles []*Article
	Comments []*Comment

	Notifications []*Notification //new articles, news, replyes etc

	Chats []*Chat
}
type UserRole struct {
	UsersIDs string
	Role     Role
}
