package domain

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
