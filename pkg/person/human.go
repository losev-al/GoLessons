package person

// Human describes a human race character
type Human struct {
	hp int
}

// NewHuman create a human race character
func NewHuman() Person {
	return &Human{hp: 50}
}

// HP return current HP
func (h *Human) HP() int {
	return h.hp
}

// SetHP set current HP
func (h *Human) SetHP(hp int) {
	h.hp = hp
}
