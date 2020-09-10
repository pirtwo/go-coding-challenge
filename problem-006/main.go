package main

import (
	"fmt"
	"math"
)

func main() {
	num := 100
	sumSqr := sumOfSquares(num)
	sqrSum := squareOfSum(num)
	fmt.Printf("sqrSum - sumSqr = %v - %v = %v", sqrSum, sumSqr, sqrSum-sumSqr)
}

func sumOfSquares(ceil int) int {
	var sum float64
	for i := 1; i <= ceil; i++ {
		sum += math.Pow(float64(i), 2)
	}

	return int(sum)
}

func squareOfSum(ceil int) int {
	var sum float64
	for i := 1; i <= ceil; i++ {
		sum += float64(i)
	}

	return int(math.Pow(sum, 2))
}
