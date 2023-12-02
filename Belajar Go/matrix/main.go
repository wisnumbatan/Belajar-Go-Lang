package main

import (
	"fmt"
)

func simplexMethod(c []float64, A [][]float64, b []float64) ([]float64, float64) {
	// Initialize variables
	m := len(A)
	n := len(A[0])

	// Create an empty matrix
	matrix := make([][]float64, m+1)
	for i := range matrix {
		matrix[i] = make([]float64, m+n+1)
	}

	// Initialize basis variables
	basis := make([]int, m)
	for i := range basis {
		basis[i] = n + i
	}

	// Fill the matrix
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i+1][j] = A[i][j]
		}
	}
	for i := 0; i < m; i++ {
		matrix[i+1][m+n] = b[i]
	}

	iterations := 0
	for {
		fmt.Printf("### Iteration %d ###\n", iterations)
		// Display the simplex table at each iteration
		for _, row := range matrix {
			fmt.Println(row)
		}

		// Find the column with the largest negative coefficient in the first row
		pivotColumn := 0
		for j := 0; j < n; j++ {
			if matrix[0][j] < matrix[0][pivotColumn] {
				pivotColumn = j
			}
		}

		// If all coefficients in the first row are non-negative, we have an optimal solution
		if matrix[0][pivotColumn] >= 0 {
			// Optimal solution found
			break
		}

		// Compute the ratio for each row to find the pivot row
		ratios := make([]float64, m)
		for i := 0; i < m; i++ {
			if matrix[i+1][pivotColumn] > 0 {
				ratios[i] = matrix[i+1][m+n] / matrix[i+1][pivotColumn]
			} else {
				ratios[i] = -1 // Set negative ratio for non-positive values
			}
		}

		// Find the pivot row
		pivotRow := 0
		for i := 1; i < m; i++ {
			if ratios[i] > 0 && (ratios[i] < ratios[pivotRow] || ratios[pivotRow] < 0) {
				pivotRow = i
			}
		}

		// Make the pivot element 1
		pivotElement := matrix[pivotRow+1][pivotColumn]
		for j := 0; j <= m+n; j++ {
			matrix[pivotRow+1][j] /= pivotElement
		}

		// Make other elements in the pivot column zero
		for i := 0; i <= m; i++ {
			if i != pivotRow+1 {
				ratio := matrix[i][pivotColumn]
				for j := 0; j <= m+n; j++ {
					matrix[i][j] -= ratio * matrix[pivotRow+1][j]
				}
			}
		}

		// Update the basis
		basis[pivotRow] = pivotColumn

		iterations++
	}

	// Extract the solution from the table
	solution := make([]float64, n)
	for i := 0; i < m; i++ {
		if basis[i] < n {
			solution[basis[i]] = matrix[i+1][m+n]
		}
	}

	optimalValue := -matrix[0][m+n]

	return solution, optimalValue
}

func main() {
	// Example linear programming problem
	c := []float64{-9, -12}
	A := [][]float64{
		{12, 1},
		{3, 1},
	}
	b := []float64{15, 18}

	solution, optimalValue := simplexMethod(c, A, b)
	fmt.Println("Optimal solution:")
	for i, val := range solution {
		fmt.Printf("X%d = %.2f\n", i+1, val)
	}
	fmt.Printf("Maximum value of Z = %.2f\n", optimalValue)
}
