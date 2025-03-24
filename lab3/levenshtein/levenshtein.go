package levenshtein

type Operation int

const (
	Replace Operation = iota
	Insert
	Delete
	Match
)

func (op Operation) String() string {
	return [...]string{"R", "I", "D", "M"}[op]
}

func min(a, b, c int) (int, Operation) {
	minVal := a
	op := Replace
	if b < minVal {
		minVal = b
		op = Insert
	}
	if c < minVal {
		minVal = c
		op = Delete
	}
	return minVal, op
}

func LevenshteinDistance(s, t []rune) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	ops := make([][]Operation, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
		ops[i] = make([]Operation, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
		ops[i][0] = Delete
		printTable(dp, s, t, i, 0)
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
		ops[0][j] = Insert
		printTable(dp, s, t, 0, j)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
				ops[i][j] = Match
			} else {
				val, op := min(
					dp[i-1][j-1]+1, // Replacement
					dp[i][j-1]+1,   // Insertion
					dp[i-1][j]+1,   // Deletion
				)
				dp[i][j] = val
				ops[i][j] = op
				printTable(dp, s, t, i, j)
			}
		}
	}

	i, j := m, n
	var operations []Operation
	for i > 0 || j > 0 {
		op := ops[i][j]
		operations = append(operations, op)
		switch op {
		case Match, Replace:
			i--
			j--
		case Insert:
			j--
		case Delete:
			i--
		}
	}

	for i, j := 0, len(operations)-1; i < j; i, j = i+1, j-1 {
		operations[i], operations[j] = operations[j], operations[i]
	}

	printTable(dp, s, t, len(dp), len(dp[0]))
	printAlignment(s, t, operations)

	return dp[m][n]
}
