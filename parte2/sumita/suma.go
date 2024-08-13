package main

func Suma(numeros []int) int {
	sum := 0
	for _, numeros := range numeros {
		sum += numeros
	}
	return sum
}
