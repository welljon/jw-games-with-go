// todo make program print out amount of tries.
// todo check is player is lying

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan()

	for {
		// Binary search strategy
		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("A. too high?")
		fmt.Println("B. too low?")
		fmt.Println("C. correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Please enter either a, b, or c")
		}
	}
}
