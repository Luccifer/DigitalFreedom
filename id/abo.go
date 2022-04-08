package id

type ABO int

const (
	Am  ABO = iota + 0
	Ap  ABO = iota + 1
	Bm  ABO = iota + 2
	Bp  ABO = iota + 3
	Om  ABO = iota + 4
	Op  ABO = iota + 5
	ABm ABO = iota + 6
	ABp ABO = iota + 7
)

func (abo ABO) String() string {
	return [...]string{"A-", "A+", "B-", "B+", "O-", "O+", "AB-", "AB+"}[abo-1]
}

func (abo ABO) EnumIndex() int {
	return int(abo)
}
