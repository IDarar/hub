package domain

import (
	"time"

	_ "gorm.io/gorm"
)

type Hub struct {
	Treatises []*Treatise
	//Master    *Master
}

type Treatise struct {
	ID           string `gorm:"primaryKey"` //E (Ethics), L (letters), TTP (Tractatus Theologico-Politicus) etc
	Title        string
	Description  string
	Date         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Parts        []Part        `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`
	Propositions []Proposition `gorm:"many2many:treatise_propositions;constraint:OnDelete:CASCADE"`

	//TODO maybe not to add it to treatises
	Literature []*Literature `gorm:"-"` //`gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`

	Difficulty    int `gorm:"-"`
	Importance    int `gorm:"-"`
	Inconsistency int `gorm:"-"`

	RatesDifficulty    []*Rate `gorm:"-"`
	RatesImportance    []*Rate `gorm:"-"`
	RatesInconsistency []*Rate `gorm:"-"`

	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`
}

type Part struct {
	ID           string `gorm:"primaryKey"` //EV (Ethics 5 part), TPI (Tractatus Politicus, First chapter) etc
	TargetID     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	FullName     string
	Description  string
	Propositions []Proposition `gorm:"many2many:part_propositions;constraint:OnDelete:CASCADE"`
	Literature   []Literature  `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`

	Difficulty    int `gorm:"-"`
	Importance    int `gorm:"-"`
	Inconsistency int `gorm:"-"`

	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`
}

type Proposition struct {
	ID          string `gorm:"primaryKey"` //EVIX (... 9 proposition), TPIVII (... 7 statement) etc
	TargetID    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Text        string

	Explanation string

	References []*Reference `gorm:"foreignKey:Target;constraint:OnDelete:CASCADE"`

	Difficulty    int `gorm:"-"`
	Importance    int `gorm:"-"`
	Inconsistency int `gorm:"-"`

	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`
}
type Literature struct {
	ID       int `gorm:"primaryKey"`
	TargetID string
	Title    string //article or book
}
type Note struct {
	ID          string
	Treatise    *Treatise
	Proposition *Proposition

	Type string //original, publisher, my etc
	Text string
}
type Reference struct {
	ID                int    `gorm:"primaryKey"`
	Target            string //ID of proposition
	TargetProposition string //To which references
}
