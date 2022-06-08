package integers

const repeatCount = 5

func Add(num1, num2 int) int {
	return num1 + num2
}

func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
