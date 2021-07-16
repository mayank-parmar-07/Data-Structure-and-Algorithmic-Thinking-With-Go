package algos

import (
	"bufio"
	"os"
	"strconv"
)

func Input_two_dimensional_arr() ([][]int, int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	matrix := make([][]int, M)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			scanner.Scan()
			matrix[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}
	return matrix, M, N
}

func Check_connected_components(matrix [][]int, visited *[][]bool, row, col, M, N int) {
	if row < M && row >= 0 && col < N && col >= 0 {
		if !(*visited)[row][col] && matrix[row][col] == 1 {
			(*visited)[row][col] = true
			Check_connected_components(matrix, visited, row+1, col+1, M, N)
			Check_connected_components(matrix, visited, row, col+1, M, N)
			Check_connected_components(matrix, visited, row-1, col+1, M, N)
			Check_connected_components(matrix, visited, row-1, col+1, M, N)
			Check_connected_components(matrix, visited, row-1, col-1, M, N)
			Check_connected_components(matrix, visited, row, col-1, M, N)
			Check_connected_components(matrix, visited, row+1, col-1, M, N)
			Check_connected_components(matrix, visited, row+1, col, M, N)
		} else {
			return
		}
	} else {
		return
	}
}
