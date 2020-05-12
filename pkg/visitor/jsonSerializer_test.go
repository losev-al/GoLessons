package visitor

import (
	"testing"

	"github.com/losev-al/GoLessons/pkg/person"
)

func TestJSONSerialize(t *testing.T) {
	t.Parallel()
	h := person.Human{Mana: 50}
	h.SetHP(50)
	o := person.Orc{Rage: 30}
	o.SetHP(70)
	tests := map[person.Serializeble]string{
		&h: "{ \"type\":\"human\", \"hp\":\"50\", \"mana\":\"50\" }",
		&o: "{ \"type\":\"orc\", \"hp\":\"70\", \"rage\":\"30\" }",
	}
	for p, want := range tests {
		got := p.Accept(JSONSerializer{})
		if got == want {
			t.Logf("JSONSerialize() success. Want = %v, Got = %v", want, got)
		} else {
			t.Errorf("JSONSerialize() failed. Want = %v, Got = %v", want, got)
		}
	}
}
