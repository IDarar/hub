package domain

type Rate struct {
	ID int `gorm:"primaryKey"`
	//UserID int
	TargetID string
	Type     string
	Value    int
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
