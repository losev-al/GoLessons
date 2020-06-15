package v1

const (
	New       State = 0
	Confirmed State = 1
	Process   State = 2
	Delivered State = 3
	Closed    State = 4
)

// State defines the states of order
type State int

// String() get a string representation of the state
func (t State) String() string {
	names := [...]string{
		"New",
		"Confirmed",
		"Process",
		"Delivered",
		"Closed",
	}
	if t < 0 || int(t) >= len(names) {
		return "Unknown"
	}
	return names[t]
}
