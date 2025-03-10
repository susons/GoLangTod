package main

import "fmt"

func main() {
	fmt.Println("***************- Hands on 0 Recursion -*****************")
	factor := factorial(4)
	fmt.Println("factor: ", factor)

	fmt.Println("***************- Hands on 0 Recursion -*****************")
}

func factorial(n int) int {

	fmt.Println("This is a N: ", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
