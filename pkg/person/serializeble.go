package person

// Serializeble is interface that the structure should implement for serialization
type Serializeble interface {
	Accept(v SerializeVisitor) string
}
