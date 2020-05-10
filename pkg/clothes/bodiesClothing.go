package clothes

// BodiesClothes describes the clothing worn on the body
type BodiesClothes struct {
	armor int
}

// Armor return armor of current Bodies Clothes
func (b *BodiesClothes) Armor() int {
	return b.armor
}

// SetArmor set armor of current Bodies Clothes
func (b *BodiesClothes) SetArmor(armor int) {
	b.armor = armor
}
