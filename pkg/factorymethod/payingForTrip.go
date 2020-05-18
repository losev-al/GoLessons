package factorymethod

import (
	"github.com/losev-al/GoLessons/pkg/command"
)

// PayingForTrip is Creator realization for paying
type PayingForTrip struct{}

// Create is factory method
func (b PayingForTrip) Create(from string, to string) command.Command {
	return command.NewPayingForTripCommand(from, to)
}
