package iteration

func Repeat(char string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += char
	}

	return result
}
