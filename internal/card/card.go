package card

import (
	// "github.com/stripe/stripe-go/v76"
)


type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

// func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, error, string) {
// 	stripe.Key = c.Secret

// 	// create a payment intent
// 	params := &stripe.PaymentIntentParams{
		
// 	}
// }