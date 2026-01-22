package reverse

func Reverse(input string) string {
	var output string
	for _, char := range input {
		output = string(char) + output
	}
	return output
}
