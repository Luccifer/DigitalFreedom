package id

type Sex int

const (
	Male       Sex = iota + 0
	Female     Sex = iota + 1
	FemaleMale Sex = iota + 2
	MaleFemale Sex = iota + 3
	Both       Sex = iota + 4
)

func (sex Sex) String() string {
	return [...]string{"Male", "Female", "FemaleMale", "MaleFemale", "Both"}[sex-1]
}

func (sex Sex) EnumIndex() int {
	return int(sex)
}
