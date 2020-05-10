package person

import "testing"

// Personmock is mock for Person interface
type Personmock struct {
	*testing.T
	hp int
}

// NewPersonmock create mock for Person Interface
func NewPersonmock(t *testing.T) *Personmock {
	return &Personmock{T: t, hp: 100}
}

// HP is mock function for get hp
func (p *Personmock) HP() int {
	return p.hp
}

// SetHP is mock function for set HP
func (p *Personmock) SetHP(hp int) {
	p.hp = hp
}
