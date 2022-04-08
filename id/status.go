package id

type Status int

const (
	Single      Status = iota + 0
	NonOfficial Status = iota + 1
	Married     Status = iota + 2
)

func (s Status) String() string {
	return [...]string{"Single", "NonOfficial", "Married"}[s-1]
}

func (s Status) EnumIndex() int {
	return int(s)
}
