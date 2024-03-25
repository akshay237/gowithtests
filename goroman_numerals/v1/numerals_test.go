package v1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRomanNumerals(t *testing.T) {
	got := convertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestRomanNumeral(t *testing.T) {
	t.Run("1 converted to I", func(t *testing.T) {
		got := convertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("2 got converted to II", func(t *testing.T) {
		got := convertToRoman(2)
		want := "II"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}

func TestRomanNumerics(t *testing.T) {
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
			name:         "2 converted to II",
			num:          2,
			expectedResp: "II",
		},
		{
			name:         "3 converted to III",
			num:          3,
			expectedResp: "III",
		},
		{
			name:         "4 converted to IV",
			num:          4,
			expectedResp: "IV",
		},
		{
			name:         "5 converted to V",
			num:          5,
			expectedResp: "V",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actualResp := convertToRoman(tc.num)
			if !cmp.Equal(actualResp, tc.expectedResp) {
				t.Errorf("got %q but want %q", actualResp, tc.expectedResp)
			}
		})
	}
}
