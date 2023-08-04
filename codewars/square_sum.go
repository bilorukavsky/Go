//https://www.codewars.com/kata/515e271a311df0350d00000f

package codewars

func SquareSum(numbers []int) int {
	var summ = 0
	for i := 0; i < len(numbers); i++ {
		summ += numbers[i] * numbers[i]
	}
	return summ
}
