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
}