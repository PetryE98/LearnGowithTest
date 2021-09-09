package arrayslice

func Sum(numbers [5]int) int {

	sum := 0

	for _, number := range numbers {

		sum += number
	}
	return sum
}

func Sumslice(numberss []int) int {

	sume := 0

	for _, integer := range numberss {

		sume += integer
	}
	return sume
}

func SumAll(numbersToSum ...[]int) []int {

	lengthOFNumbers := len(numbersToSum)
	sums := make([]int, lengthOFNumbers)

	for i, number := range numbersToSum {
		sums[i] = Sumslice(number)
	}
	return sums

}
