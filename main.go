package main

import "fmt"

// BubbleSoft - takes in a slice of ints and sorts using the bubble sorting algorithm
func BubbleSort(input []int) []int {
	n := len(input)
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < n-1; i++ {

			if input[i] > input[i+1] {
				fmt.Println("Swapping values")

				input[i], input[i+1] = input[i+1], input[i]

				swapped = true
			}
		}
	}

	return input
}

func main() {
	fmt.Println("Bubble Sorting Algorithm in Go")

	unsortedInput := []int{5, 3, 4, 1, 2}
	sorted := BubbleSort(unsortedInput)

	fmt.Println(sorted)
}
