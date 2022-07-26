package account

import (
	"finbot/olddomain/account/currency"
	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID
	AccNumber string
	Balance   float64
	Currency  currency.Currency
}
