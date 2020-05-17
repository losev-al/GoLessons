package chainofresponsibility

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/losev-al/GoLessons/pkg/payer"
)

type yandexPaymentHandler struct {
	next PaymentHandler
}

// NewYandexPaymentHandler create yandexPaymentHandler
func NewYandexPaymentHandler(n PaymentHandler) PaymentHandler {
	return &yandexPaymentHandler{next: n}
}

func (h *yandexPaymentHandler) Pay(p *payer.Payer) {
	if p.HasPayPalWallet {
		fmt.Println(aurora.Magenta("Try payment reception by Yandex money"))
	} else {
		if h.next != nil {
			h.next.Pay(p)
		}
	}
}
