import (
	"testing"
)
func countSpaces(str string) int {
	charsArr := []rune(str)
	countSpaces := 0
	for i := 0; i < len(str); i++ {
		if charsArr[i] == rune(' '){
			countSpaces++
		}
	}
	return countSpaces	
}
func TestReplaceSpace(t *testing.T){
	cases := []struct{
		input string
		expect string
	}{
		{"abc d e f ", `abc%20d%20e%20f%20`},
		{"abc d e f", `abc%20d%20e%20f`},
		{"  ", `%20%20`},
	}
	for _, thisCase := range cases {
		
	}
}