package v2

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRomanNumerals(t *testing.T) {
	testcases := []struct {
		name         string
		num          int
		expectedResp string
	}{
		{
			name:         "1 converted to I",
			num:          1,
			expectedResp: "I",
		},
		{
			name:         "5 converted to V",
			num:          5,
			expectedResp: "V",
		},
		{
			name:         "9 converted to IX",
			num:          9,
			expectedResp: "IX",
		},
		{
			name:         "15 converted to XV",
			num:          15,
			expectedResp: "XV",
		},
		{
			name:         "1984 converted to MCMLXXXIV",
			num:          1984,
			expectedResp: "MCMLXXXIV",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actualResp := transformToRoman(tc.num)
			if !cmp.Equal(actualResp, tc.expectedResp) {
				t.Errorf("got %q but want %q", actualResp, tc.expectedResp)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	testcases := []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%s got converted to %d", tc.Roman, tc.Arabic), func(t *testing.T) {
			got := converToArabic(tc.Roman)
			if !cmp.Equal(got, tc.Arabic) {
				t.Errorf("got %q but want %q", got, tc.Arabic)
			}
		})
	}
}
