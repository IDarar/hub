package domain

import (
	"time"

	_ "gorm.io/gorm"
)

/*user prop
status prop
mark of prop
rate of props
add to favs
underlines, notes*/

type User struct {
	//TODO last proposition opened
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Password string

	RegisteredAt time.Time
	LastVisitAt  time.Time
	Session      Session
	OnlineChan   chan int `gorm:"-"`
	IsOnline     bool     `gorm:"-"`

	UserLists *UserLists `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	Roles    []Role     `gorm:"many2many:user_role;"` //admin, SuperModerator, ContentModerator, ForumModerator
	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`

	Notifications []*Notification `gorm:"-"` //new articles, news, replyes etc
	Chats         []*Chat         `gorm:"-"`
}

type Role struct {
	Role  string `gorm:"primaryKey"`
	Users []User `gorm:"many2many:user_role;"`
}
