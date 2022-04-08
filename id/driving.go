package id

/* TODO: - Add Code Comments from https://en.wikipedia.org/wiki/International_Driving_Permit */

type Driving int

const (
	Unlearned Driving = iota + 0
	A1        Driving = iota + 1
	A2        Driving = iota + 2
	B1        Driving = iota + 3
	B2        Driving = iota + 4
	C1        Driving = iota + 5
	C1E       Driving = iota + 6
	C         Driving = iota + 7
	D1        Driving = iota + 8
	D1E       Driving = iota + 9
	D         Driving = iota + 10
	DE        Driving = iota + 11
)

func (d Driving) String() string {
	return [...]string{"Unlearned", "A1", "A2", "B1", "B2", "C1", "C1E", "C", "D1", "D1E", "D", "DE"}[d-1]
}

func (d Driving) EnumIndex() int {
	return int(d)
}
