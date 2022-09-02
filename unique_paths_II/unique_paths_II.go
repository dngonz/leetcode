package unique_paths_II

func uniquePaths(obstacle [][]int) int {
	if len(obstacle) == 0 {
		return 0
	}

	m, n := len(obstacle)-1, len(obstacle[0])-1
	if obstacle[m][n] == 1 {
		return 0
	}

	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if obstacle[i][j] == 1 {
				continue
			}
			if i == 0 && j == 0 {
				matrix[i][j] = 1 - obstacle[0][0]
			} else if i == 0 {
				matrix[i][j] = matrix[i][j-1]
			} else if j == 0 {
				matrix[i][j] = matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
			}
		}
	}

	return matrix[m][n]
}
