import (
	"testing"
)
func TestPremutation(t *testing.T) {
	cases := []struct {
		input1 string
		input2 string
		expect bool
	}{
		{"apple", "papel", true},
		{"carrot", "tarroc", true},
		{"hello", "holll", false}
	}
	for -, thiscase := range cases {
		result := premutation(thisCase.input1, thisCase.input2)
		if result != thisCase.expect {
			t.Errorf("reverse: input (%q, %q), expected %v, but got %v\n",
				thisCase.input1, thisCase.input2, thisCase.expect, result)

		}
	}
}