package id

type Issuer int

const (
	main Issuer = iota + 0
)

func (i Issuer) String() string {
	return [...]string{"DF 547"}[i-1]
}

func (i Issuer) EnumIndex() int {
	return int(i)
}
