package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numbers := []int{2, 3, 17, 25, 29, 30}
	for _, number := range numbers {
		if isPrime(number) {
			fmt.Printf("%d é primo\n", number)
		} else {
			fmt.Printf("%d não é primo\n", number)
		}
	}
}
