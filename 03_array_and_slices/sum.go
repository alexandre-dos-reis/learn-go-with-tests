// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices
package arrayandslices

func SumArray5(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlices(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllWithMake(numbersToSum ...[]int) []int {
	lenOfNumbers := len(numbersToSum)
	sums := make([]int, lenOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlices(numbers)
	}
	return sums
}

func SumAllWithAppend(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlices(numbers))
	}
	return sums
}
