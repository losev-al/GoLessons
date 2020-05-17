// Эмулируется ситуация описанная здесь: https://youtu.be/mkZblULa60E?list=PLeIe7YX-CxPSiVznMeE0trL021W1kshx0&t=1290
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/losev-al/GoLessons/pkg/payer"

	"github.com/losev-al/GoLessons/pkg/command"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	p := &payer.Payer{HasBankСard: false, HasPayPalWallet: true, HasYandexMoney: true}
	s := command.NewBookingSupervisor(p, "A", "B", "C", "D")
	message, err := s.Run()
	if err == nil {
		fmt.Println(message)
	}
}
