// using bitvector 
func isUniqueChars (str string) bool {
	if len(str) > 128 {
		return false
	}
	var checker unit = 0; // use this int as a bit vector
	for _, char := range str {
		val := uint(char) -uint('a')
	}
	
}