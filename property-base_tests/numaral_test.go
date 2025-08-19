package propertybasetests

import "testing"

func TestRomanNumeral(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
	}

	for _, tt := range cases {
		t.Run(tt.Description, func(t *testing.T) {
			got := ConvertToRoman(tt.Arabic)

			if got != tt.Want {
				t.Errorf("got %q, want %q", got, tt.Want)
			}
		})
	}

}
