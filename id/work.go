package id

type Work struct {
	country     string
	companyName string
	position    string
	salary      int
	vacation    bool
	WorkTime    WorkTime
	current     bool
	from        string
	to          string
}

type WorkTime int

const (
	PartTime WorkTime = iota + 0
	FullTime WorkTime = iota + 1
)

func (wt WorkTime) String() string {
	return [...]string{"PartTime", "FullTime"}[wt-1]
}

func (wt WorkTime) EnumIndex() int {
	return int(wt)
}
