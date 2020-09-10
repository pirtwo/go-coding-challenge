package main

import (
	"fmt"
)

func main() {
	sum := sumOfMultiples(1000, 3, 5)
	fmt.Println(sum)
}

func sumOfMultiples(ceil int, nums ...int) int {
	sum := 0
	for i := 1; i < ceil; i++ {
		for _, num := range nums {
			if i%num == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
