package chainofresponsibility

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/losev-al/GoLessons/pkg/payer"
)

type payPalPaymentHandler struct {
	next PaymentHandler
}

// NewPayPalPaymentHandler create payPalPaymentHandler
func NewPayPalPaymentHandler(n PaymentHandler) PaymentHandler {
	return &payPalPaymentHandler{next: n}
}

func (h *payPalPaymentHandler) Pay(p *payer.Payer) {
	if p.HasPayPalWallet {
		fmt.Println(aurora.Magenta("Try payment reception by PayPal wallet"))
	} else {
		if h.next != nil {
			h.next.Pay(p)
		}
	}
}
