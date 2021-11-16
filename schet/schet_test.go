package schet_test

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/bankschet/bankschet/schet"
)

func TestAccountType(t *testing.T) {
	tests := []struct {
		line    string
		name    string
		number  string
		want    string
		wantErr error
	}{
		{
			line:    line(),
			name:    "Счёт физ.лица резидента",
			number:  "40817810455556666666",
			want:    "40817",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			number, err := schet.New(tt.number)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("\nwant: %q\nget: %q\ntest: %s", tt.wantErr, err, tt.line)
			}

			if number.Type() != tt.want {
				t.Errorf("\nwant: %q\nget: %q\ntest: %s", tt.want, number.Type(), tt.line)
			}
		})
	}
}

func TestAccountCurrency(t *testing.T) {
	tests := []struct {
		line    string
		name    string
		number  string
		want    string
		wantErr error
	}{
		{
			line:    line(),
			name:    "Рублевый счёт",
			number:  "40817810455556666666",
			want:    "RUR",
			wantErr: nil,
		},
		{
			line:    line(),
			name:    "Евровый счёт",
			number:  "40817978455556666666",
			want:    "EUR",
			wantErr: nil,
		},
		{
			line:    line(),
			name:    "Не умеем открывать счета в беллоруских рублях",
			number:  "40817906455556666666",
			want:    "",
			wantErr: schet.ErrUnknownCurrency,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			number, err := schet.New(tt.number)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("\nwant: %q, get: %q\ntest: %s", tt.wantErr, err, tt.line)
			}

			if number.CurrencyCode != tt.want {
				t.Errorf("\nwant: %q\nget: %q\ntest: %s", tt.want, number.CurrencyCode, tt.line)
			}
		})
	}
}

func line() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return "It was not possible to recover file and line number information about function invocations!"
}
