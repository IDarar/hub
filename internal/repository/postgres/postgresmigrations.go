package postgres

import (
	"time"
)

type Hub struct {
	Treatises []*Treatise
	Master    *Master
}
type Master struct {
	Name        string
	Description string
	LifeSpans   []*Part
	Literature  []string

	Articles []*Article
	Comments []*Comment
}
type Treatise struct {
	ID           string //E (Ethics), L (letters), TTP (Tractatus Theologico-Politicus) etc
	Name         string
	Description  string
	Date         string
	Parts        []*Part
	Propositions []*Proposition

	Literature []string

	Difficulty    int
	Importance    int
	Inconsistency int

	RatesDifficulty    []*Rate
	RatesImportance    []*Rate
	RatesInconsistency []*Rate

	Articles []*Article
	Comments []*Comment
}

type Part struct {
	ID           string //EV (Ethics 5 part), TPI (Tractatus Politicus, First chapter) etc
	Name         string
	FullName     string
	Description  string
	Propositions []*Proposition
	Literature   []string

	Difficulty    int
	Importance    int
	Inconsistency int

	Articles []*Article
	Comments []*Comment
}

type Proposition struct {
	ID          string //EVIX (... 9 proposition), TPIVII (... 7 statement) etc
	Name        string
	Description string
	Text        string

	Explanation string

	References []*Reference

	Difficulty    int
	Importance    int
	Inconsistency int

	Articles []*Article
	Comments []*Comment
}

type Note struct {
	ID          string
	Treatise    *Treatise
	Proposition *Proposition

	Type string //original, publisher
	Text string
}
type Reference struct {
	Target            string
	TargetProposition string
	Text              string
}

type Notification struct {
	UserID        int
	Target        string //which topic
	Type          string
	ReplyerUserID string
}
type Chat struct {
	Users    []string
	Messages []*Message
}
type Message struct {
	ID        string //TODO yeah, rather make it common
	UserID    int
	Text      string
	CreatedAt time.Time
	DeletedTo []string //To which users do not show message
}

type Rate struct {
	UserID int
	Target string
	Type   string
}

type Comment struct {
	Target  string
	UserID  string
	Replyes []*Comment
	Text    string
}
type Article struct {
	Target string //Maybe anything can be a target
	UserID string
	Name   string
	Text   string

	Comments []*Comment
	Replyes  []*Comment
}

type Role string

type User struct {
	//TODO last proposition opened
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"uniqueIndex"`
	Email        string `gorm:"uniqueIndex"`
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
	Articles   []*Article `gorm:"-"`
	Comments   []*Comment `gorm:"-"`

	Notifications []*Notification `gorm:"-"` //new articles, news, replyes etc

	Chats []*Chat `gorm:"-"`
}

type Session struct {
	UserID       int       `gorm:"primaryKey"`
	RefreshToken string    `json:"refreshToken" bson:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt" bson:"expiresAt"`
}

type UserRole struct {
	UsersIDs string
	Role     Role
}

type UserLists struct {
	UserID int

	UserTreatises    []*UserTreatise
	UserParts        []*UserPart
	UserPropositions []*UserProposition
}
type UserTreatise struct {
	UserID int

	TargetTreatise string //E (Ethics), L (letters), TTP (Tractatus Theologico-Politicus) etc
	Status         string

	DifficultyRate    int8
	ImportanceRate    int8
	InconsistencyRate int8
	Progress          int
}
type UserPart struct {
	UserID int

	TargetPart        string //EV (Ethics 5 part), TPI (Tractatus Politicus, First chapter) etc
	Status            string
	DifficultyRate    int8
	ImportanceRate    int8
	InconsistencyRate int8
	Progress          int
}
type UserProposition struct {
	UserID int

	LocalText         string         //with notes, underlines etc
	Marks             [3]interface{} //first two are indexes, third is format type
	TargetProposition string         //EVIX (... 9 proposition), TPIVII (... 7 statement) etc
	Status            string
	DifficultyRate    int8
	ImportanceRate    int8
	InconsistencyRate int8

	UserNotes *UserNote
}

type UserNote struct {
	UserID  int
	LocalID int //number of note in particular proposition
	Target  string
	Text    string
}
