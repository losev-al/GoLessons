package person

// Orc describes a orc race character
type Orc struct {
	hp int
}

// NewOrc create a Orc race character
func NewOrc() Person {
	return &Orc{hp: 70}
}

// HP return current HP
func (o *Orc) HP() int {
	return o.hp
}

// SetHP set current HP
func (o *Orc) SetHP(hp int) {
	o.hp = hp
}
