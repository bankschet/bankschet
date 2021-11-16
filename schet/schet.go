package schet

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bankschet/bankschet/schetmoney"
)

type Account struct {
	Number          string
	FirstType       string
	SecondType      string
	CurrencyCode    string
	CurrencyNumeric string
}

var ErrUnknownCurrency = errors.New("unknown currency")

func New(number string) (Account, error) {
	if len(number) != 20 {
		return Account{}, errors.New("account number mast be 20 digits long")
	}

	if strings.IndexFunc(number, func(c rune) bool { return c >= '0' || c <= '9' }) == -1 {
		return Account{}, errors.New("account numbers mast contains only digits")
	}

	account := Account{
		Number:     number,
		FirstType:  number[0:3],
		SecondType: number[3:5],
	}

	account.CurrencyNumeric = number[5:8]

	var ok bool
	account.CurrencyCode, ok = schetmoney.Codes[account.CurrencyNumeric]
	if !ok {
		return account, fmt.Errorf(
			"unexpected currency numeric code %q: %w", account.CurrencyNumeric, ErrUnknownCurrency,
		)
	}

	return account, nil
}

func (a Account) Type() string {
	return a.FirstType + a.SecondType
}
