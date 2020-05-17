package payer

// Payer describe user's payment options
type Payer struct {
	HasBankСard     bool
	HasPayPalWallet bool
	HasYandexMoney  bool
}
