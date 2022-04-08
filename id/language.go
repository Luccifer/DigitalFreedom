package id

type Language struct {
	name   string
	origin bool
	level  LanguageLevel
}

type LanguageLevel int

const (
	Unconformed  LanguageLevel = iota + 0
	Elemetary    LanguageLevel = iota + 1
	Intermediate LanguageLevel = iota + 2
	Advanced     LanguageLevel = iota + 3
	Proficient   LanguageLevel = iota + 4
)

func (ll LanguageLevel) String() string {
	return [...]string{"Unconformed", "Elemetary", "Intermediate", "Advanced", "Proficient"}[ll-1]
}

func (ll LanguageLevel) EnumIndex() int {
	return int(ll)
}
