package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	transitions map[rune]*Node
	suffixLink  *Node
	outputLink  *Node
	patterns    []int
}

func buildTrie(patterns []string) (root *Node) {
	fmt.Println("\n=== Построение бора ===")
	root = &Node{
		transitions: make(map[rune]*Node),
	}
	for i, p := range patterns {
		fmt.Printf("\nДобавление шаблона %d: '%s'\n", i+1, p)
		current := root
		for pos, c := range p {
			if _, ok := current.transitions[c]; !ok {
				fmt.Printf("  Создание нового узла для символа '%c' (позиция %d)\n", c, pos+1)
				current.transitions[c] = &Node{
					transitions: make(map[rune]*Node),
					suffixLink:  nil,
					outputLink:  nil,
					patterns:    nil,
				}
			}
			current = current.transitions[c]
		}
		current.patterns = append(current.patterns, i+1)
		fmt.Printf("  Шаблон '%s' добавлен в узел (ID шаблона: %d)\n", p, i+1)
	}
	return root
}

func buildAutomaton(root *Node) {
	fmt.Println("\n=== Построение автомата ===")
	queue := []*Node{}
	
	fmt.Println("\nИнициализация суффиксных ссылок для первого уровня:")
	for char, child := range root.transitions {
		child.suffixLink = root
		queue = append(queue, child)
		fmt.Printf("  Установка суффиксной ссылки для '%c' -> корень\n", char)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("\nОбработка узла (паттерны: %v)\n", current.patterns)
		for char, child := range current.transitions {
			fmt.Printf("  Обработка символа '%c'\n", char)
			
			temp := current.suffixLink
			for temp != nil && temp.transitions[char] == nil {
				temp = temp.suffixLink
			}
			
			if temp == nil {
				child.suffixLink = root
				fmt.Printf("    Суффиксная ссылка -> корень (не найдено соответствие)\n")
			} else {
				child.suffixLink = temp.transitions[char]
				fmt.Printf("    Суффиксная ссылка -> узел для '%c'\n", char)
			}

			if child.suffixLink.patterns != nil {
				child.outputLink = child.suffixLink
				fmt.Printf("    Выходная ссылка -> узел с паттернами %v\n", child.suffixLink.patterns)
			} else {
				child.outputLink = child.suffixLink.outputLink
				if child.outputLink != nil {
					fmt.Printf("    Выходная ссылка унаследована -> узел с паттернами %v\n", child.outputLink.patterns)
				} else {
					fmt.Println("    Выходная ссылка отсутствует")
				}
			}

			queue = append(queue, child)
		}
	}
}

func findMatches(text string, root *Node, patternLengths map[int]int) (results [][2]int) {
	fmt.Println("\n=== Поиск совпадений ===")
	current := root
	for i, c := range text {
		fmt.Printf("\nСимвол '%c' (позиция %d):\n", c, i+1)
		
		for current != nil && current.transitions[c] == nil {
			if current == root {
				fmt.Println("  Нет перехода из корня - остаемся в корне")
			} else {
				fmt.Printf("  Нет перехода - переход по суффиксной ссылке к узлу (паттерны: %v)\n", current.suffixLink.patterns)
			}
			current = current.suffixLink
		}
		
		if current == nil {
			current = root
			fmt.Println("  Переход в корень")
			continue
		}
		
		current = current.transitions[c]
		fmt.Printf("  Переход в узел для '%c' (паттерны: %v)\n", c, current.patterns)

		temp := current
		for temp != nil {
			if len(temp.patterns) > 0 {
				for _, p := range temp.patterns {
					pos := i - patternLengths[p] + 2
					fmt.Printf("    Найдено совпадение: шаблон %d на позиции %d\n", p, pos)
					results = append(results, [2]int{pos, p})
				}
			}
			if temp.outputLink != nil {
				fmt.Printf("    Переход по выходной ссылке к узлу (паттерны: %v)\n", temp.outputLink.patterns)
			}
			temp = temp.outputLink
		}
	}
	return results
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	scanner.Scan()
	text := scanner.Text()

	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

	patterns := make([]string, n)
	patternLengths := make(map[int]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		patterns[i] = scanner.Text()
		patternLengths[i+1] = len(patterns[i])
	}

	root := buildTrie(patterns)
	buildAutomaton(root)
	results := findMatches(text, root, patternLengths)

	sort.Slice(results, func(i, j int) bool {
		if results[i][0] == results[j][0] {
			return results[i][1] < results[j][1]
		}
		return results[i][0] < results[j][0]
	})

	fmt.Println("\n=== Результаты ===")
	for _, res := range results {
		fmt.Printf("%d %d\n", res[0], res[1])
	}
}