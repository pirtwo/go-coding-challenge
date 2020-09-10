package main

import (
	"fmt"
	"strconv"
)

func main() {
	numA, numB, largest := 0, 0, 0
	for i := 999; i > 99; i-- {		
		for j := 999; j > 99; j-- {			
			if isPalindrome(i*j) && i*j > largest {
				numA = i
				numB = j
				largest = i * j
			}
		}
	}

	fmt.Printf("Largest Palindrome made from the product of two 3-digit numbers: %d * %d = %d", numA, numB, largest)
}

func isPalindrome(n int) bool {
	num := strconv.Itoa(n)
	rev := reverse(num)
	return num == rev
}

func reverse(str string) string {
	rev := ""
	for _, v := range str {
		rev = string(v) + rev
	}

	return rev
}
