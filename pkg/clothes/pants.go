package clothes

// Pants describes the clothing worn on legs
type Pants struct {
	armor int
}

// Armor return armor of current Pants
func (p *Pants) Armor() int {
	return p.armor
}

// SetArmor set armor of current Bodies Pants
func (p *Pants) SetArmor(armor int) {
	p.armor = armor
}
