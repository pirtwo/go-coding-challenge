package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	primes := sieveOfEratosthenes(2000000)
	for _, v := range primes {
		sum+=v
	}
	
	fmt.Printf("Sum of all prime numbers below 2000000 is: %v", sum)
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