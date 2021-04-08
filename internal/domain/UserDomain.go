package domain

import "time"

type Role string

type User struct {
	//TODO last proposition opened
	ID           uint
	Name         string
	Email        string
	Password     string `gorm:"-"`
	RegisteredAt time.Time
	LastVisitAt  time.Time
	Role         UserRole `gorm:"-"` //admin, SuperModerator, ContentModerator, ForumModerator

	EncryptedPassword string   `gorm:"-"`
	OnlineChan        chan int `gorm:"-"`
	IsOnline          bool     `gorm:"-"`

	UserListID int        `gorm:"-"`
	UserLists  *UserLists `gorm:"-"`

	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`

	Notifications []*Notification `gorm:"-"` //new articles, news, replyes etc

	Chats []*Chat `gorm:"-"`
}
type UserRole struct {
	UsersIDs string
	Role     Role
}
