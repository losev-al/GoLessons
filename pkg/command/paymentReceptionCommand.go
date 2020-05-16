package command

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/logrusorgru/aurora"
)

type paymentReceptionCommand struct {
}

// NewPaymentReceptionCommand create command realization for payment reception
func NewPaymentReceptionCommand() Command {
	return &paymentReceptionCommand{}
}

func (c *paymentReceptionCommand) Execute() error {
	fmt.Println(aurora.Green("Try payment reception"))
	result := rand.Intn(100)
	if result < chanceOfSuccess {
		fmt.Println(aurora.Green("Payment reception success"))
		return nil
	}
	fmt.Println(aurora.Red("Payment reception failed"))
	return errors.New("Payment reception failed")
}

func (c *paymentReceptionCommand) Undo() error {
	fmt.Println(aurora.Yellow("Try cancel payment reception"))
	result := rand.Intn(100)
	if result < chanceOfSuccess {
		fmt.Println(aurora.Yellow("Cancel payment reception success"))
		return nil
	}
	fmt.Println(aurora.Red("Cancel payment reception failed"))
	return errors.New("Cancel payment reception failed")
}
