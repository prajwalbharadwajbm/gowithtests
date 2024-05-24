package arraysandslices

import "fmt"

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum = sum + number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lenOfNumbers := len(numbersToSum)
	sum := make([]int, lenOfNumbers)

	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers)
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
}
