package command

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/losev-al/GoLessons/pkg/chainofresponsibility"
	"github.com/losev-al/GoLessons/pkg/payer"

	"github.com/logrusorgru/aurora"
)

type paymentReceptionCommand struct {
	p *payer.Payer
	h chainofresponsibility.PaymentHandler
}

// NewPaymentReceptionCommand create command realization for payment reception
func NewPaymentReceptionCommand(p *payer.Payer) Command {
	pp := chainofresponsibility.NewPayPalPaymentHandler(nil)
	ym := chainofresponsibility.NewYandexPaymentHandler(pp)
	bc := chainofresponsibility.NewBankPaymentHandler(ym)
	return &paymentReceptionCommand{p: p, h: bc}
}

func (c *paymentReceptionCommand) Execute() error {
	c.h.Pay(c.p)
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
