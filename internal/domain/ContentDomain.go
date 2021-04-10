package domain

import (
	"time"

	_ "gorm.io/gorm"
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
	ID           string `gorm:"primaryKey"` //E (Ethics), L (letters), TTP (Tractatus Theologico-Politicus) etc
	Title        string
	Description  string
	Date         string
	CreatedAt    time.Time
	Parts        []*Part        `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`
	Propositions []*Proposition `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`

	Literature []*Literature `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`

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
	ID       string `gorm:"primaryKey"` //EV (Ethics 5 part), TPI (Tractatus Politicus, First chapter) etc
	TargetID string

	Name         string
	FullName     string
	Description  string
	Propositions []*Proposition `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`
	Literature   []*Literature  `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`

	Difficulty    int `gorm:"-"`
	Importance    int `gorm:"-"`
	Inconsistency int `gorm:"-"`

	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`
}

type Proposition struct {
	ID          string `gorm:"primaryKey"` //EVIX (... 9 proposition), TPIVII (... 7 statement) etc
	TargetID    string
	Name        string
	Description string
	Text        string

	Explanation string

	References []*Reference `gorm:"-"`

	Difficulty    int `gorm:"-"`
	Importance    int `gorm:"-"`
	Inconsistency int `gorm:"-"`

	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`
}
type Literature struct {
	TargetID string `gorm:"primaryKey"`
	Title    string
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