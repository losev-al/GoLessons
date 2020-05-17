package chainofresponsibility

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/losev-al/GoLessons/pkg/payer"
)

type bankPaymentHandler struct {
	next PaymentHandler
}

// NewBankPaymentHandler create bankPaymentHandler
func NewBankPaymentHandler(n PaymentHandler) PaymentHandler {
	return &bankPaymentHandler{next: n}
}

func (h *bankPaymentHandler) Pay(p *payer.Payer) {
	if p.HasBankСard {
		fmt.Println(aurora.Magenta("Try payment reception by Bank card"))
	} else {
		if h.next != nil {
			h.next.Pay(p)
		}
	}
}
