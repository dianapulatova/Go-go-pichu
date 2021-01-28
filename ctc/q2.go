import (
	"testing"
)

func testReverse(t *testing.T) {
	cases := [struct] {
		input string
		expect string
	}
	{
		{"abcd", "dcba"},
		{"abcde", "edcba"}
	}
}

for _, thisCase := range cases {
	result := reverse(thisCase.input)
	if result != thisCase.expect {
		t.Errorf("reverse: input %q, expect %q, but got %q\n", thisCase.input, thisCase.expect, result)
	}
}

// function to reverse a C style string, by returning a new reversed string
// func reverse(str string) string {
	length := len(str)
	revArr := []byte(str)
	for i := length / 2; i >= 0; i-- {
		revArr[i], revArr[length-i-1] = revArr[length-i-1], revArr[i]
	}
	return string(revArr)
}