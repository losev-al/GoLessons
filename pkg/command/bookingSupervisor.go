package command

import (
	"errors"
	"fmt"

	"github.com/losev-al/GoLessons/pkg/payer"

	"github.com/logrusorgru/aurora"
)

const repeatCount int = 3

// BookingSupervisor implement Supervisor for booking and paying for tickets along the route (include сompensating transaction)
type BookingSupervisor struct {
	commands []Command
}

// NewBookingSupervisor create Supervisor for booking and paying for tickets along the route
func NewBookingSupervisor(p *payer.Payer, cities ...string) *BookingSupervisor {
	result := BookingSupervisor{}
	citiesCount := len(cities)
	result.commands = make([]Command, 2*(citiesCount-1)+1)
	for i := 0; i < citiesCount-1; i++ {
		result.commands[i] = NewBookingTripCommand(cities[i], cities[i+1])
		result.commands[citiesCount+i] = NewPayingForTripCommand(cities[i], cities[i+1])
	}
	result.commands[citiesCount-1] = NewPaymentReceptionCommand(p)
	return &result
}

// Run run all commans in BookingSupervisor
func (s *BookingSupervisor) Run() (string, error) {
	for i := 0; i < len(s.commands); i++ {
		if runWithRepeat(s.commands[i], true) != nil {
			if s.RunUndo(i-1) != nil {
				fmt.Println(aurora.Cyan("Very bad situation. Technical support specialist assistance is required."))
				return "", errors.New("Compensating Transaction failed")
			}
			return "Problems occurred when ordering. Try choosing a different route or try ordering again later.", nil
		}
	}
	return "The order is executed", nil
}

// RunUndo run undo commans in BookingSupervisor
func (s *BookingSupervisor) RunUndo(step int) error {
	for i := 0; i < step+1; i++ {
		if runWithRepeat(s.commands[step-i], false) != nil {
			return errors.New("Undo failed")
		}
	}
	return nil
}

func runWithRepeat(c Command, doRunForward bool) error {
	var method func() error
	if doRunForward {
		method = c.Execute
	} else {
		method = c.Undo
	}
	for i := 0; i < repeatCount; i++ {
		err := method()
		if err == nil {
			return nil
		}
	}
	return errors.New("Run command failed")
}
