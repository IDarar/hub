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

	DifficultyRate    Rate `gorm:"-"` //`gorm:"many2many:usertreatise_rates;constraint:OnDelete:CASCADE"`
	ImportanceRate    Rate `gorm:"-"` //`gorm:"many2many:usertreatise_rates;constraint:OnDelete:CASCADE"`
	InconsistencyRate Rate `gorm:"-"` //`gorm:"many2many:usertreatise_rates;constraint:OnDelete:CASCADE"`
	Progress          int
}

type UserPart struct {
	ID     int `gorm:"primaryKey"`
	UserID int

	TargetPart string //EV (Ethics 5 part), TPI (Tractatus Politicus, First chapter) etc

	Status            string
	DifficultyRate    Rate `gorm:"-"` //`gorm:"many2many:userpart_rates;constraint:OnDelete:CASCADE"`
	ImportanceRate    Rate `gorm:"-"` //`gorm:"many2many:userpart_rates;constraint:OnDelete:CASCADE"`
	InconsistencyRate Rate `gorm:"-"` //`gorm:"many2many:userpart_rates;constraint:OnDelete:CASCADE"`

	Progress int //all props of part / completed
}

type UserProposition struct {
	ID int `gorm:"primaryKey"`

	UserID int

	LocalText string //with notes, underlines etc
	//Marks             [3]interface{} //first two are indexes, third is format type
	TargetProposition string //EVIX (... 9 proposition), TPIVII (... 7 statement) etc
	Status            string //complete, unknow, in proccess etc
	DifficultyRate    Rate   `gorm:"-"` //`gorm:"many2many:userproposition_rates;constraint:OnDelete:CASCADE"`
	ImportanceRate    Rate   `gorm:"-"` //`gorm:"many2many:userproposition_rates;constraint:OnDelete:CASCADE"`
	InconsistencyRate Rate   `gorm:"-"` //`gorm:"many2many:userproposition_rates;constraint:OnDelete:CASCADE"`

	UserNotes []*UserNote `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

//on get append it to original prop's notes
type UserNote struct {
	ID     int `gorm:"primaryKey"`
	UserID int

	Index  int    //palce of note
	Target string //to user prop
	Text   string
	Type   string `json:"type"` //usertype only. or later not only

}
