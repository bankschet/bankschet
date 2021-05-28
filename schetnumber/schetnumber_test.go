package schetnumber_test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/bankschet/bankschet/schetnumber"
)

func TestAccountType(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)

	testCases := []struct {
		line        int
		name        string
		number      string
		expected    string
		expectedErr error
	}{
		{
			line:        line(),
			name:        "Счёт физ.лица резидента",
			number:      "40817810455556666666",
			expected:    "40817",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			number, err := schetnumber.New(tc.number)

			linkToToest := fmt.Sprintf("%s:%d", file, tc.line)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("\nexpected: %q\nrecieved: %q\ncase: %s", tc.expectedErr, err, linkToToest)
			}

			if number.Type() != tc.expected {
				t.Errorf("\nexpected: %q\nrecieved: %q\ncase: %s", tc.expected, number.Type(), linkToToest)
			}
		})
	}
}

func TestAccountCurrency(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)

	testCases := []struct {
		line        int
		name        string
		number      string
		expected    string
		expectedErr error
	}{
		{
			line:        line(),
			name:        "Рублевый счёт",
			number:      "40817810455556666666",
			expected:    "RUR",
			expectedErr: nil,
		},
		{
			line:        line(),
			name:        "Евровый счёт",
			number:      "40817978455556666666",
			expected:    "EUR",
			expectedErr: nil,
		},
		{
			line:        line(),
			name:        "Не умеем открывать счета в беллоруских рублях",
			number:      "40817906455556666666",
			expected:    "",
			expectedErr: schetnumber.ErrUnknownCurrency,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			number, err := schetnumber.New(tc.number)

			linkToToest := fmt.Sprintf("%s:%d", file, tc.line)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("\nexpected: %q, recieved: %q\ncase: %s", tc.expectedErr, err, linkToToest)
			}

			if number.CurrencyCode != tc.expected {
				t.Errorf("\nexpected: %q\nrecieved: %q\ncase: %s", tc.expected, number.CurrencyCode, linkToToest)
			}
		})
	}
}

func TestAccountTypeCode(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)

	testCases := []struct {
		line        int
		name        string
		number      string
		expected    string
		expectedErr error
	}{
		{
			line:        line(),
			name:        "Текущий счёт резидента",
			number:      "40817810455556666666",
			expected:    "0100",
			expectedErr: nil,
		},
		{
			line:        line(),
			name:        "Текущий счёт не резидента",
			number:      "40820978455556666666",
			expected:    "0100",
			expectedErr: nil,
		},
		{
			line:        line(),
			name:        "Электронный кошелёк, электронные средства платежа (ЭСП) для переводов электронных денежных средств (ЭДС)",
			number:      "40914978455556666666",
			expected:    "9902",
			expectedErr: nil,
		},
		{
			line:        line(),
			name:        "Неизвестный код для типа 00000",
			number:      "00000810455556666666",
			expected:    "",
			expectedErr: schetnumber.ErrUnknownTypeCode,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			number, err := schetnumber.New(tc.number)

			linkToToest := fmt.Sprintf("%s:%d", file, tc.line)

			if number.TypeCode != tc.expected {
				t.Errorf("\nexpected: %q\nrecieved: %q\ncase: %s", tc.expected, number.TypeCode, linkToToest)
			}

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("\nexpected: %q\nrecieved: %q\ncase: %s", tc.expectedErr, err, linkToToest)
			}
		})
	}
}

func line() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}
