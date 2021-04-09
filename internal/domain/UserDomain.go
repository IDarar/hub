package domain

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	//TODO last proposition opened
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"uniqueIndex"`
	Email        string `gorm:"uniqueIndex"`
	Password     string
	RegisteredAt time.Time
	LastVisitAt  time.Time
	Session      Session
	Roles        []Role `gorm:"many2many:user_role;"` //admin, SuperModerator, ContentModerator, ForumModerator

	EncryptedPassword string   `gorm:"-"`
	OnlineChan        chan int `gorm:"-"`
	IsOnline          bool     `gorm:"-"`

	UserListID int        `gorm:"-"`
	UserLists  *UserLists `gorm:"-"`
	Articles   []*Article `gorm:"-"`
	Comments   []*Comment `gorm:"-"`

	Notifications []*Notification `gorm:"-"` //new articles, news, replyes etc

	Chats []*Chat `gorm:"-"`
}

type Role struct {
	Role  string `gorm:"primaryKey"`
	Users []User `gorm:"many2many:user_role;"`
}
