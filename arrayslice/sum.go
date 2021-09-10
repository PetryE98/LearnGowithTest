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

	for _, number := range numbersToSum {
		sums = append(sums, Sumslice(number))
	}
	return sums

}

func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {

		tail := numbers[1:]
		sums = append(sums, Sumslice(tail))
	}
	return sums

}
