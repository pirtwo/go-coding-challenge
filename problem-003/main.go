package main

import (
	"fmt"
)

func main() {
	primeFactors := primeFactors(600851475143)
	fmt.Printf("Prime factors of 600851475143 are: %d", primeFactors)
}

func findDivisible(num int) int {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return i
		}
	}
	return 1
}

func primeFactors(num int) []int {
	curr, factorTree, primeFactors := 0, []int{}, []int{}
	factorTree = append(factorTree, num)
	for len(factorTree) > 0 {
		curr, factorTree = factorTree[len(factorTree)-1], factorTree[:len(factorTree)-1]
		f1 := findDivisible(curr)
		f2 := curr / f1
		if f1 != 1 {
			factorTree = append(factorTree, f1, f2)
		} else {
			primeFactors = append(primeFactors, curr)
		}
	}

	return primeFactors
}
