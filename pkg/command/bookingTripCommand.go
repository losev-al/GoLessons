package command

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/logrusorgru/aurora"
)

type bookingTripCommand struct {
	from string
	to   string
}

// NewBookingTripCommand create command realization for booking trip
func NewBookingTripCommand(from string, to string) Command {
	return &bookingTripCommand{from: from, to: to}
}

func (c *bookingTripCommand) Execute() error {
	fmt.Println(aurora.Sprintf(aurora.Green("Try booking trip from %v to %v"), c.from, c.to))
	result := rand.Intn(100)
	if result < chanceOfSuccess {
		fmt.Println(aurora.Sprintf(aurora.Green("Booking trip from %v to %v success"), c.from, c.to))
		return nil
	}
	fmt.Println(aurora.Sprintf(aurora.Red("Booking trip from %v to %v failed"), c.from, c.to))
	return errors.New("Booking failed")
}

func (c *bookingTripCommand) Undo() error {
	fmt.Println(aurora.Sprintf(aurora.Yellow("Try cancel booking trip from %v to %v"), c.from, c.to))
	result := rand.Intn(100)
	if result < chanceOfSuccess {
		fmt.Println(aurora.Sprintf(aurora.Yellow("Cancel booking trip from %v to %v success"), c.from, c.to))
		return nil
	}
	fmt.Println(aurora.Sprintf(aurora.Red("Cancel booking trip from %v to %v failed"), c.from, c.to))
	return errors.New("Cancel booking failed")
}
