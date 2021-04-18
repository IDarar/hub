package domain

import (
	_ "gorm.io/gorm"
)

type UserLists struct {
	UserID int `gorm:"primaryKey"`

	UserTreatises    []*UserTreatise    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	UserParts        []*UserPart        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	UserPropositions []*UserProposition `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Rates            []*Rate            `gorm:"many2many:userlists_rates;constraint:OnDelete:CASCADE"`
}

//how will it work
//user sends request with his ID and target of req
//db searches target in db and changes value
type UserTreatise struct {
	UserID int `gorm:"primaryKey"`

	TargetTreatise string `gorm:"primaryKey"` //E (Ethics), L (letters), TTP (Tractatus Theologico-Politicus) etc
	//TODO types of status
	Status string

	DifficultyRate    int
	ImportanceRate    int
	InconsistencyRate int

	Progress    int
	IsCompleted *bool
}

type UserPart struct {
	UserID int `gorm:"primaryKey"`

	TargetPart string `gorm:"primaryKey"` //EV (Ethics 5 part), TPI (Tractatus Politicus, First chapter) etc

	Status      string
	IsCompleted *bool

	DifficultyRate    int
	ImportanceRate    int
	InconsistencyRate int

	Progress int //all props of part / completed
}

type UserProposition struct {
	UserID int `gorm:"primaryKey"`

	LocalText         string //with notes, underlines etc
	TargetProposition string `gorm:"primaryKey"` //EVIX (... 9 proposition), TPIVII (... 7 statement) etc
	Status            string //complete, unknow, in proccess etc
	IsCompleted       *bool

	DifficultyRate    int
	ImportanceRate    int
	InconsistencyRate int
	UserNotes         []*UserNote `gorm:"-"`
}

//on get append it to original prop's notes
type UserNote struct {
	//TODO because of FKeys restrictions should check ex of target
	ID     int `gorm:"primaryKey"`
	UserID int

	Index  int    //place of note
	Target string //to user prop
	Text   string
	Type   string `json:"type"` //usertype only. or later not only
}
