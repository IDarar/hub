package domain

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
