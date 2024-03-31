package v2

import (
	"strings"
)

type RomanNumeral struct {
	value  int
	symbol string
}

var allRomanNumerals = []RomanNumeral{
	{
		value:  1000,
		symbol: "M",
	},
	{
		value:  900,
		symbol: "CM",
	},
	{
		value:  500,
		symbol: "D",
	},
	{
		value:  400,
		symbol: "CD",
	},
	{
		value:  100,
		symbol: "C",
	},
	{
		value:  90,
		symbol: "XC",
	},
	{
		value:  50,
		symbol: "L",
	},
	{
		value:  40,
		symbol: "XL",
	},
	{
		value:  10,
		symbol: "X",
	},
	{
		value:  9,
		symbol: "IX",
	},
	{
		value:  5,
		symbol: "V",
	},
	{
		value:  4,
		symbol: "IV",
	},
	{
		value:  1,
		symbol: "I",
	},
}

func convertToRoman(num int) string {

	var result strings.Builder

	for num > 0 {
		switch {
		case num > 9:
			result.WriteString("X")
			num -= 10
		case num > 8:
			result.WriteString("IX")
			num -= 9
		case num > 4:
			result.WriteString("V")
			num -= 5
		case num > 3:
			result.WriteString("IV")
			num -= 4
		default:
			result.WriteString("I")
			num -= 1
		}
	}

	return result.String()
}

func transformToRoman(num int) string {
	var result strings.Builder
	for _, roman := range allRomanNumerals {
		for num >= roman.value {
			result.WriteString(roman.symbol)
			num -= roman.value
		}
	}
	return result.String()
}

func converToArabic(roman string) int {

	arabic := 0
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.symbol) {
			arabic += numeral.value
			roman = strings.TrimPrefix(roman, numeral.symbol)
		}
	}
	return arabic
}
