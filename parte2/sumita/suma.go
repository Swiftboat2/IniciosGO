package main

func Suma(numeros []int) int {
	sum := 0
	for _, numeros := range numeros {
		sum += numeros
	}
	return sum
}
func SumaTodo(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Suma(numbers))
	}
	return sums
}
func SumaTodoFinal(numbersToSum ...[]int) []int {
	var sumas []int
	for _, numbers := range numbersToSum {
	if len(numbers) == 0{
	sumas = append(sumas, 0)
	}else{
		tail := numbers[1:]
		sumas = append(sumas, Suma(tail))
		}
	}
	return sumas
}
