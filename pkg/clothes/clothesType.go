package clothes

const (
	BodyClothing Type = 0
	Pants        Type = 1
)

// Type defines the types of clothing
type Type int

// String() get a string representation of the type
func (t Type) String() string {
	names := [...]string{
		"BodyClothing",
		"Pants"}
	if t < 0 || int(t) >= len(names) {
		return "Unknown"
	}
	return names[t]
}
