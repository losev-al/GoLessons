package visitor

import (
	"fmt"

	"github.com/losev-al/GoLessons/pkg/person"
)

// XMLSerializer supports methods for serialization in xml
type XMLSerializer struct{}

// SerializeHuman return serialized string for human
func (s XMLSerializer) SerializeHuman(h *person.Human) string {
	return fmt.Sprintf("<person type=\"human\" hp=\"%v\" mana=\"%v\" />", h.HP(), h.Mana)
}

// SerializeOrc return serialized string for orc
func (s XMLSerializer) SerializeOrc(o *person.Orc) string {
	return fmt.Sprintf("<person type=\"orc\" hp=\"%v\" rage=\"%v\" />", o.HP(), o.Rage)
}
