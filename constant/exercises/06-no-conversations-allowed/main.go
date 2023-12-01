package main

import "fmt"

func main() {
    // 1. Declare the number of rows and columns
    const rows, cols = 3, 3

    // 2. Create a 2D slice (matrix)
    matrix := make([][]int, rows)
    for i := range matrix {
        matrix[i] = make([]int, cols)
    }

    // 3. Fill the matrix with numbers
    num := 1
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            matrix[i][j] = num
            num++
        }
    }

    // 4. Print the matrix
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
}