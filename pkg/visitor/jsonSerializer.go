package visitor

import (
	"fmt"

	"github.com/losev-al/GoLessons/pkg/person"
)

// JSONSerializer supports methods for serialization in json
type JSONSerializer struct{}

// SerializeHuman return serialized string for human
func (s JSONSerializer) SerializeHuman(h *person.Human) string {
	return fmt.Sprintf("{ \"type\":\"human\", \"hp\":\"%v\", \"mana\":\"%v\" }", h.HP(), h.Mana)
}

// SerializeOrc return serialized string for orc
func (s JSONSerializer) SerializeOrc(o *person.Orc) string {
	return fmt.Sprintf("{ \"type\":\"orc\", \"hp\":\"%v\", \"rage\":\"%v\" }", o.HP(), o.Rage)
}
