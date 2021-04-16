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
//db searches target in db and changes values

type UserTreatise struct {
	ID int `gorm:"primaryKey"`

	UserID int

	TargetTreatise string //E (Ethics), L (letters), TTP (Tractatus Theologico-Politicus) etc
	Status         string

	DifficultyRate    int
	ImportanceRate    int
	InconsistencyRate int

	Progress    int
	IsCompleted bool
}

type UserPart struct {
	ID     int `gorm:"primaryKey"`
	UserID int

	TargetPart string //EV (Ethics 5 part), TPI (Tractatus Politicus, First chapter) etc

	Status      string
	IsCompleted bool

	DifficultyRate    int
	ImportanceRate    int
	InconsistencyRate int

	Progress int //all props of part / completed
}

type UserProposition struct {
	ID int `gorm:"primaryKey"`

	UserID int

	LocalText         string  //with notes, underlines etc
	Marks             []*Mark `gorm:"-"` //first two are indexes, third is format type
	TargetProposition string  //EVIX (... 9 proposition), TPIVII (... 7 statement) etc
	Status            string  //complete, unknow, in proccess etc
	IsCompleted       bool
	//Difficulty, Importance, Inconsistency
	//There can be only 3 rates (one for one type)  foreignKey:Refer;joinForeignKey:UserReferID;
	DifficultyRate    int
	ImportanceRate    int
	InconsistencyRate int
	UserNotes         []*UserNote `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

//on get append it to original prop's notes
type UserNote struct {
	ID     int `gorm:"primaryKey"`
	UserID int

	Index  int    //place of note
	Target string //to user prop
	Text   string
	Type   string `json:"type"` //usertype only. or later not only

}

//TODO
type Mark struct {
}
