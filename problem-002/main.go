package main

import (
	"fmt"
)

func main() {
	currFibo, sum := 0, 0
	nextFibo := fibo()

	for true {
		currFibo = nextFibo()
		if currFibo > 4000000 {
			break
		}
		if currFibo%2 == 0 {
			sum += currFibo
		}
	}
	
	fmt.Printf("Sum of even fibo numbers below 4000000 is: %d", sum)
}

func fibo() func() int {
	n1, n2, curr := 1, 0, 0

	return func() int {
		curr = n1 + n2
		n2 = n1
		n1 = curr
		return curr
	}
}
