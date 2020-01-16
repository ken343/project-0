package filter

const (
	LT  = "less than"
	LTE = "less than or equal to"
	GT  = "greter than"
	GTE = "greter than or equal to"
	EQL = "equal to"
)

type Filterable interface {
	OrdinalVal() float64
	LexicalVal() string
}

func Ordinal(f []Filterable, flag string, compare float64) []Filterable {
	surrogateSlice := make([]Filterable, 0)
	var isFiltered bool

	for i := 0; i < len(f); i++ {
		isFiltered = compareFloat64(f[i].OrdinalVal(), compare, flag)

		if isFiltered {
			surrogateSlice = append(surrogateSlice, f[i])
		}
	}
	return surrogateSlice
}

func Lexical(f []Filterable, flag string, compare string) []Filterable {
	surrogateSlice := make([]Filterable, 0)
	var isFiltered bool

	for i := 0; i < len(f); i++ {
		isFiltered = compareString(f[i].LexicalVal(), compare, flag)

		if isFiltered {
			surrogateSlice = append(surrogateSlice, f[i])
		}
	}
	return surrogateSlice
}
