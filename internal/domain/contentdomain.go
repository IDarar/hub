package domain

import (
	"time"

	_ "gorm.io/gorm"
)

type Hub struct {
	Treatises []*Treatise
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
	Rates      []*Rate       `gorm:"many2many:treatise_rates;constraint:OnDelete:CASCADE"`

	Difficulty    int
	Importance    int
	Inconsistency int

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
	Rates        []*Rate       `gorm:"many2many:part_rates;constraint:OnDelete:CASCADE"`

	Difficulty    int
	Importance    int
	Inconsistency int

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

	Notes      []*Note      `gorm:"foreignKey:Target;constraint:OnDelete:CASCADE"`
	References []*Reference `gorm:"foreignKey:Target;constraint:OnDelete:CASCADE"`

	Rates         []*Rate `gorm:"many2many:proposition_rates;constraint:OnDelete:CASCADE"`
	Difficulty    int
	Importance    int
	Inconsistency int

	Articles []*Article `gorm:"-"`
	Comments []*Comment `gorm:"-"`
}

type Literature struct {
	ID       int `gorm:"primaryKey"`
	TargetID string
	Title    string //article or book
}
type Note struct {
	ID         int    `gorm:"primaryKey"`
	TreatiseID string `json:"treatise_id,omitempty"` //will be taken through prop/part-prop target
	Target     string `json:"target"`                //to which belongs
	Index      int    //palce of note

	Type string `json:"type"` //original, publisher, my etc
	Text string `json:"text,omitempty"`
}
type Reference struct {
	ID                int    `gorm:"primaryKey"`
	Target            string //ID of proposition
	TargetProposition string //To which references
}
