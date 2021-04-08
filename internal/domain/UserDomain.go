package domain

import (
	"time"

	_ "gorm.io/gorm"
)

type Role string

type User struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
	LastVisitAt  time.Time
	Session      Session
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
