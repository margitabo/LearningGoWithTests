package arrays

func Sum(numbers []int) int {

	var sum int
	//for i := 0; i < len(numbers); i++ {
	//	sum += numbers[i]
	//}

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersArrayToSum ...[]int) []int {
	//lengthOfNumbers := len(numbersArrayToSum)
	//sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, numberArray := range numbersArrayToSum {
		sums = append(sums, Sum(numberArray))
	}
	return sums
}
