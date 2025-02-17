package main

import (
	"fmt"

	"molexar.org/gauss/internal"
)

func main() {
	matrix, err := internal.ReadMatrix("input.txt")
	if err != nil {
		fmt.Println("ошибка при чтении матрицы")
	}

	internal.GaussianElimination(matrix)
	x := internal.BackSubstitution(matrix)

	fmt.Println(x)
}
