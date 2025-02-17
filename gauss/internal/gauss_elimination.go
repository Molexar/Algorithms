package internal

import "sync"

func GaussianElimination(matrix [][]float64) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	n := len(matrix)

	for k := 0; k < n-1; k++ {
		wg.Add(n - k - 1)
		for i := k + 1; i < n; i++ {
			go func(i, k int) {
				defer wg.Done()
				factor := matrix[i][k] / matrix[k][k]

				mu.Lock()
				for j := k; j < n+1; j++ {
					matrix[i][j] -= factor * matrix[k][j]
				}
				mu.Unlock()
			}(i, k)
		}
		wg.Wait()
	}
}

func BackSubstitution(matrix [][]float64) []float64 {
	n := len(matrix)
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		x[i] = matrix[i][n] / matrix[i][i]
		for j := i - 1; j >= 0; j-- {
			matrix[j][n] -= matrix[j][i] * x[i]
		}
	}
	return x
}
