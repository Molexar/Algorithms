package main

import (
	"fmt"

	"molexar.org/binsearch/internal"
)

func main() {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	targets := []int{7, 9, 10}

	fmt.Println(internal.Search(inputs, targets, 3))
}
