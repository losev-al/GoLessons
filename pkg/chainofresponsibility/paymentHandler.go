package chainofresponsibility

import (
	"github.com/losev-al/GoLessons/pkg/payer"
)

// PaymentHandler describe base interface of chain of responsibility pattern
type PaymentHandler interface {
	Pay(p *payer.Payer)
}
