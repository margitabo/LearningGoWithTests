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

func SumAll(arraysToSum ...[]int) []int {
	//lengthOfNumbers := len(numbersArrayToSum)
	//sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, arr := range arraysToSum {
		sums = append(sums, Sum(arr))
	}
	return sums
}

func SumAllTails(arraysToSum ...[]int) []int {
	var sums []int
	for _, arr := range arraysToSum {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			tailArr := arr[1:]
			sums = append(sums, Sum(tailArr))
		}
	}
	return sums
}
