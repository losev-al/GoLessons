package command

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/logrusorgru/aurora"
)

type payingForTripCommand struct {
	from string
	to   string
}

// NewPayingForTripCommand create command realization for paying for a trip
func NewPayingForTripCommand(from string, to string) Command {
	return &payingForTripCommand{from: from, to: to}
}

func (c *payingForTripCommand) Execute() error {
	fmt.Println(aurora.Sprintf(aurora.Green("Try paying for a trip from %v to %v"), c.from, c.to))
	result := rand.Intn(100)
	if result < chanceOfSuccess {
		fmt.Println(aurora.Sprintf(aurora.Green("Paying for a trip from %v to %v success"), c.from, c.to))
		return nil
	}
	fmt.Println(aurora.Sprintf(aurora.Red("Paying for a trip from %v to %v failed"), c.from, c.to))
	return errors.New("Booking failed")
}

func (c *payingForTripCommand) Undo() error {
	fmt.Println(aurora.Sprintf(aurora.Yellow("Try cancel paying for a trip from %v to %v"), c.from, c.to))
	result := rand.Intn(100)
	if result < chanceOfSuccess {
		fmt.Println(aurora.Sprintf(aurora.Yellow("Cancel paying for a trip from %v to %v success"), c.from, c.to))
		return nil
	}
	fmt.Println(aurora.Sprintf(aurora.Red("Cancel paying for a trip from %v to %v failed"), c.from, c.to))
	return errors.New("Cancel booking failed")
}
