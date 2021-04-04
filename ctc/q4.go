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
	}
}