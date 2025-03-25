package levenshtein

type Config struct {
	SpecialInsertionCharacter   rune
	SpecialInsertionCost        int
	SpecialReplacementCharacter rune
	SpecialReplacementCost      int
}

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

func calculateCosts(dp [][]int, i, j int, sChar, tChar rune, cfg Config) (int, Operation) {
	replaceCost := dp[i-1][j-1]
	if sChar == cfg.SpecialReplacementCharacter {
		replaceCost += cfg.SpecialReplacementCost
	} else if sChar != tChar {
		replaceCost += 1
	}

	insertCost := dp[i][j-1]
	if tChar == cfg.SpecialInsertionCharacter {
		insertCost += cfg.SpecialInsertionCost
	} else {
		insertCost += 1
	}

	deleteCost := dp[i-1][j] + 1

	return min(replaceCost, insertCost, deleteCost)
}

func LevenshteinDistance(s, t []rune, cfg Config) int {
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
				val, op := calculateCosts(dp, i, j, s[i-1], t[j-1], cfg)

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
