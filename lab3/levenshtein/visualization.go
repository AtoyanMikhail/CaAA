package levenshtein

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clearScreen() {
	cmd := exec.Command("clear")
	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printAlignment(s, t []rune, operations []Operation) {
	fmt.Println("\nAlignment Steps:")
	fmt.Print("  ")
	for _, op := range operations {
		fmt.Printf(" %-3s", op)
	}
	fmt.Println("\n  " + repeat("────", len(operations)))

	fmt.Print("S ")
	for i, op := range operations {
		switch op {
		case Match, Replace:
			fmt.Printf(" %-2c", s[0])
			s = s[1:]
		case Delete:
			fmt.Printf(" %-2c", s[0])
			s = s[1:]
		default:
			fmt.Print(" * ")
		}
		if i < len(operations)-1 {
			fmt.Print("│")
		}
	}
	fmt.Println("\n  " + repeat("────", len(operations)))

	fmt.Print("T ")
	for i, op := range operations {
		switch op {
		case Match, Replace:
			fmt.Printf(" %-2c", t[0])
			t = t[1:]
		case Insert:
			fmt.Printf(" %-2c", t[0])
			t = t[1:]
		default:
			fmt.Print(" * ")
		}
		if i < len(operations)-1 {
			fmt.Print("│")
		}
	}
	fmt.Println("\n")
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func printTable(dp [][]int, s, t []rune, i, j int) {
	clearScreen()

	fmt.Print("       ")
	for _, ch := range t {
		fmt.Printf(" %c ", ch)
	}
	fmt.Println()

	for row := 0; row < len(dp); row++ {
		if row == 0 {
			fmt.Print("   ")
		} else {
			fmt.Printf(" %c ", s[row-1])
		}

		for col := 0; col < len(dp[row]); col++ {
			if row == i && col == j {
				fmt.Printf(" \033[1;32m%2d\033[0m", dp[row][col])
			} else {
				fmt.Printf(" %2d", dp[row][col])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(300 * time.Millisecond) // Pause for visualization
}
