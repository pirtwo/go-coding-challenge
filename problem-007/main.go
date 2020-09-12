package main

import (
	"fmt"
	"math"
)

func main() {
	primes := sieveOfEratosthenes(500000)
	fmt.Printf("10001st Prime number is: %v", primes[10001])
}

func sieveOfEratosthenes(n int) []int {
	var numbers []Number

	for i := 2; i <= n; i++ {
		numbers = append(numbers, Number{i, true})
	}

	for i := 0; i < int(math.Sqrt(float64(len(numbers)))); i++ {
		if numbers[i].isPrime {
			curr := numbers[i].value
			for j := curr; j*curr <= n; j++ {
				numbers[(j*curr)-2].isPrime = false
			}
		}
	}

	return filterNotPrime(numbers)
}

func filterNotPrime(list []Number) []int {
	var primes []int
	for _, v := range list {
		if v.isPrime {
			primes = append(primes, v.value)
		}
	}

	return primes
}

type Number struct {
	value   int
	isPrime bool
}
