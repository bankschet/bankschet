// package schetmoney maps currency codes to currency numeric codes https://en.wikipedia.org/wiki/ISO_4217
package schetmoney

var Codes = map[string]string{
	"810": "RUR",
	"643": "RUB",
	"840": "USD",
	"978": "EUR",
}

var Nums = map[string]string{
	"RUR": "810",
	"RUB": "643",
	"USD": "840",
	"EUR": "978",
}
