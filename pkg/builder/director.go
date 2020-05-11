package builder

// Director describes interface for create character
type Director interface {
	Construct()
}

// CreateHumanDirector creates a director implementation for human character
func CreateHumanDirector(b Builder) Director {
	return &humanDirector{b: b}
}

// CreateOrcDirector creates a director implementation for orc character
func CreateOrcDirector(b Builder) Director {
	return &orcDirector{b: b}
}
