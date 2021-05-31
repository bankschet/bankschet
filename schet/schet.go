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
	TypeCode        string
	CurrencyCode    string
	CurrencyNumeric string
}

var (
	ErrInvalidLength         = errors.New("account number mast be 20 digits long")
	ErrNumberContainsLetters = errors.New("account numbers mast contains only digits")
	ErrUnknownCurrency       = errors.New("unknown currency")
	ErrUnknownTypeCode       = errors.New("unknown type code")
)

func New(number string) (Account, error) {
	if len(number) != 20 {
		return Account{}, ErrInvalidLength
	}

	if strings.IndexFunc(number, func(c rune) bool { return c >= '0' || c <= '9' }) == -1 {
		return Account{}, ErrNumberContainsLetters
	}

	firstType := number[0:3]
	secondType := number[3:5]

	typeCode, ok := TypeCodes[firstType+secondType]
	if !ok {
		return Account{}, fmt.Errorf("unknown code for type %q: %w", firstType+secondType, ErrUnknownTypeCode)
	}

	currencyNumeric := number[5:8]

	currencyCode, ok := schetmoney.Codes[currencyNumeric]
	if !ok {
		return Account{}, fmt.Errorf("unexpected currency numeric code %q: %w", currencyNumeric, ErrUnknownCurrency)
	}

	return Account{
		Number:          number,
		FirstType:       firstType,
		SecondType:      secondType,
		TypeCode:        typeCode,
		CurrencyNumeric: currencyNumeric,
		CurrencyCode:    currencyCode,
	}, nil
}

func (a Account) Type() string {
	return a.FirstType + a.SecondType
}

var TypeCodes = map[string]string{
	"40817": "0100", // Текущий счет резидента
	"40820": "0100", // Текущий счет не резидента
	// "":      "0200", // Расчетный счет
	// "":      "0300", // Бюджетный счет
	// "":      "0400", // Корреспондентский счет
	// "":      "0500", // Корреспондентский субсчет
	// "":      "0600", // Счет доверительного управления
	// "":      "0700", // Специальный банковский счет
	// "":      "0701", // Специальный банковский счет банковского платежного агента
	// "":      "0702", // Специальный банковский счет банковского платежного субагента
	// "":      "0703", // Специальный банковский счет платежного агента
	// "":      "0704", // Специальный банковский счет поставщика
	// "":      "0705", // Торговый банковский счет
	// "":      "0706", // Клиринговый банковский счет
	// "":      "0707", // Счет гарантийного фонда платежной системы
	// "":      "0708", // Номинальный счет
	// "":      "0709", // Счет эскроу
	// "":      "0710", // Залоговый счет
	// "":      "0711", // Специальный банковский счет должника
	// "":      "0800", // Депозитный счет суда
	// "":      "0900", // Депозитный счет подразделения службы судебных приставов
	// "":      "1000", // Депозитный счет правоохранительных органов
	// "":      "1100", // Депозитный счет нотариуса
	// "":      "1200", // Счет по вкладу
	// "":      "1300", // Счет по депозиту
	// "":      "1400", // Счет по государственному оборонному заказу
	// "":      "1500", // Специальный избирательный счет
	// "":      "1600", // Специальный счет фонда референдума
	// "":      "1700", // Расчетный счет застройщика, предусмотренный федеральным законом от 30.12.2004 N 214-ФЗ
	// "":      "1800", // Счет организаций, находящихся в федеральной собственности. Финансовые организации (40501)
	// "":      "1900", // Счет организаций, находящихся в государственной (кроме федеральной) собственности. Финансовые организации (40601)
	// "":      "2000", // Счет негосударственных организаций. Финансовые организации (40701)
	// "":      "2100", // Счет, предусмотренный положениями статьи 175 Жилищного кодекса Российской Федерации от 29.12.2004 N 188-ФЗ (40604, 40705)
	// "":      "2200", // Публичный депозитный счет
	// "":      "2300", // Брокерский счет и специальные депозитарные счета, предусмотренные федеральным законом от 22.04.1996 N 39-ФЗ
	// "":      "9901", // Корпоративное ЭСП
	"40903": "9902", // ЭСП, не являющееся корпоративным, аноним
	"40914": "9902", // ЭСП, не являющееся корпоративным, упрощённая идентификация
	// "":      "9999", // Иной счет
}
