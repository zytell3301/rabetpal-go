package rabetpal

import (
	"time"
)

type BankAccounts struct {
	AccountNumber  string    `cql:"account_number"`
	Iban           string    `cql:"iban"`
	ExpirationDate time.Time `cql:"expiration_date"`
}
