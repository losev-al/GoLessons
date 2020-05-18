// Эмулируется ситуация описанная здесь: https://youtu.be/mkZblULa60E?list=PLeIe7YX-CxPSiVznMeE0trL021W1kshx0&t=1290
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/losev-al/GoLessons/pkg/bookingsupervisor"
	"github.com/losev-al/GoLessons/pkg/factorymethod"
	"github.com/losev-al/GoLessons/pkg/payer"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	p := &payer.Payer{HasBankСard: false, HasPayPalWallet: true, HasYandexMoney: true}
	s := bookingsupervisor.NewBookingSupervisor(p, factorymethod.BookingTripCreator{}, factorymethod.PayingForTrip{}, "A", "B", "C", "D")
	message, err := s.Run()
	if err == nil {
		fmt.Println(message)
	}
}
