package factorymethod

import (
	"github.com/losev-al/GoLessons/pkg/command"
)

// BookingTripCreator is Creator realization for booking
type BookingTripCreator struct{}

// Create is factory method
func (b BookingTripCreator) Create(from string, to string) command.Command {
	return command.NewBookingTripCommand(from, to)
}
