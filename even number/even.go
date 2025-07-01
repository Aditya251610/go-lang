package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		if isEven(i) {
			sum += i
		}
	}
	fmt.Println("Sum of even numbers from 0 to 100:", sum)
}
