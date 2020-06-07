package person

import "fmt"

// Person have methods for work with main property of character
type Person interface {
	Name() string
	HP() int
	SetHP(hp int) error
}

type person struct {
	hp int
	name string
}

// HP get hp of person
func (p *person) HP() int {
	return p.hp
}

// SetHP set hp of person
func (p *person) SetHP(hp int) (err error) {
	if hp < 0 {
		return fmt.Errorf("HP cannot be less than 0. The passed value %v", hp)
	}
	if p.hp != hp {
		p.hp = hp
	}
	return
}

// GetName get name of person
func (p *person) Name() string {
	return p.name
}

// NewPerson create person
func NewPerson(name string, hp int) (Person, error) {
	p := &person{name: name}
	if err :=p.SetHP(hp);err != nil {
		return nil, err
	}
	return p, nil
}
