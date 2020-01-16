// Package filter provides tools for operating on slices of userdefined types
// that have implemented the Filterable interface.
package filter

// These flags will determine what kind of filter logic will be applied when
// comparing the "compare" value to the filterable slice. These flags also apply when
// comparing types aphabetically.
const (
	LT  = "less than"                // Filter out values less than the compare value.
	LTE = "less than or equal to"    // Filter out values less than or equal to the compare value.
	GT  = "greater than"             // Filter out values greater than the compare value.
	GTE = "greater than or equal to" // Filter out values greater than or equal to the compare value.
	EQL = "equal to"                 // Filter out values equal to the compare value.
)

// Filterable interface describes any user defined type that may need to be filter out of
// a slice. OrdinalVal and LexicalVal methods should return a value that defines behavior for how a
// type should be compared amongst other instances of the same type. In the event that there could be multiple
// ways that a type could compare instances of itself, whether numerically or alphabetically, the
// SetOrdinalOption and SetLexicalOption can be used to change what aspect of the type should be used for
// comparison.
type Filterable interface {
	SetOrdinalOption(option string)
	SetLexicalOption(option string)
	OrdinalVal() float64
	LexicalVal() string
}

// Ordinal filter will return an array of values that have been compared numberically
// against a "compare" value. The comparison is determined by which filter flag is
// pass into the function.
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

// Lexical filter will return an array of values that have been compared alphabetically
// against a "compare" value. The comparison is determined by which filter flag is
// pass into the function.
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
