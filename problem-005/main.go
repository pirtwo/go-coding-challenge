package main

import (
	"fmt"
)

func main() {
	curr, from, to := 0, 2, 20
	curr = from
	for true {
		if isEvenlyDivisible(curr, from, to) {
			break
		}
		curr++
	}

	fmt.Printf("Number %v is evenly divisable to all numbers in range %v to %v.", curr, from, to)
}

func isEvenlyDivisible(n int, from int, to int) bool {
	isDivisible := true
	for i := to; i >= from; i-- {
		if n%i != 0 {
			isDivisible = false
			break
		}
	}

	return isDivisible
}
