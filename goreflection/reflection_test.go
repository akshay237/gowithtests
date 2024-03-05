package goreflection

import (
	"reflect"
	"strings"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong no of function calls, got %d want %d", len(got), 1)
	}

	if expected != got[0] {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}

func TestWalk1(t *testing.T) {

	testcases := []struct {
		name         string
		input        interface{}
		expectedResp []string
	}{
		{
			name: "struct with one string field",
			input: struct {
				Name string
			}{"chris"},
			expectedResp: []string{"chris"},
		},
		{
			name: "struct with two string fields",
			input: struct {
				Name string
				City string
			}{"Chris", "London"},
			expectedResp: []string{"Chris", "London"},
		},
		{
			name: "struct with different type of fields",
			input: struct {
				Name string
				Age  int
			}{"Chris", 27},
			expectedResp: []string{"Chris"},
		},
		{
			name: "struct with nested struct",
			input: struct {
				Name string
				Info struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			expectedResp: []string{"Chris", "London"},
		},
		{
			name: "struct passed as pointer",
			input: &Person{
				"Chris",
				Profile{
					33,
					"London",
				},
			},
			expectedResp: []string{"Chris", "London"},
		},
		{
			name: "struct having slices of struct",
			input: []Profile{
				{
					35,
					"New Jersy",
				},
				{
					34,
					"Washington DC",
				},
			},
			expectedResp: []string{"New Jersy", "Washington DC"},
		},
		{
			name: "struct having arrays of struct",
			input: [2]Profile{
				{
					35,
					"New Jersy",
				},
				{
					34,
					"Washington DC",
				},
			},
			expectedResp: []string{"New Jersy", "Washington DC"},
		},
		{
			name: "when input is maps",
			input: map[string]string{
				"cow": "moo",
				"bar": "thar",
			},
		},
	}

	for _, tc := range testcases {
		if strings.Contains(tc.name, "maps") {
			t.Run(tc.name, func(t *testing.T) {
				var got []string
				walk(tc.input, func(input string) {
					got = append(got, input)
				})

				for _, val := range tc.input.(map[string]string) {
					assertionContains(t, got, val)
				}
			})
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			walk(tc.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tc.expectedResp) {
				t.Errorf("got %v, want %v", got, tc.expectedResp)
			}
		})
	}

	t.Run("with Channels", func(t *testing.T) {
		aChan := make(chan Profile)

		go func() {
			aChan <- Profile{30, "New Delhi"}
			aChan <- Profile{31, "Bangalore"}
			close(aChan)
		}()

		var got []string
		want := []string{"New Delhi", "Bangalore"}
		walk(aChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("with Functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{31, "New Delhi"}, Profile{32, "Mumbai"}
		}

		var got []string
		want := []string{"New Delhi", "Mumbai"}
		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertionContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if needle == x {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
