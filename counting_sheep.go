//https://www.codewars.com/kata/54edbc7200b811e956000556/go

package main

func CountSheeps(numbers []bool) int {
	var sum int = 0
	for _, elem := range numbers {
		if elem {
			sum++
		}
	}
	return sum
}

func main(){}
