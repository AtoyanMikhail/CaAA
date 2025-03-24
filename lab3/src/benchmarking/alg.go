package benchmarking

import (
	"fmt"
)

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func levenshteinDistance(s, t []rune) int {
	m := len(s)
	n := len(t)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Filling in the firsts row and columns
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Filling in the distance table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1], // Replacement
					dp[i][j-1],   // Insertion
					dp[i-1][j],   // Deletion
				)
			}
		}
	}

	return dp[m][n]
}

func main() {
	var s1, s2 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	distance := levenshteinDistance([]rune(s1), []rune(s2))
	fmt.Println(distance)
}
