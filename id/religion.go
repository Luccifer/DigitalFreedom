package id

type Religion int

const (
	AtheismAgnostic Religion = iota + 0
	Christianity    Religion = iota + 1
	Islam           Religion = iota + 2
	Hinduism        Religion = iota + 3
	Folk            Religion = iota + 4
)

func (r Religion) String() string {
	return [...]string{"None", "Christianity", "Islam", "Hinduism", "Folk"}[r-1]
}

func (r Religion) EnumIndex() int {
	return int(r)
}
