package domain

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
