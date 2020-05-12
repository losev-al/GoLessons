package person

// SerializeVisitor describe interface for serialize info about person
type SerializeVisitor interface {
	SerializeHuman(h *Human) string
	SerializeOrc(o *Orc) string
}
