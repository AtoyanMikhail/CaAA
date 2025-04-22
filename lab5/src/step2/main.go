package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const alphabetSize = 5

// Node представляет узел автомата Ахо-Корасик
type Node struct {
	children [alphabetSize]*Node // Дочерние узлы
	failure  *Node               // Failure-ссылка
	outputs  []int               // Позиции фиксированных символов в шаблоне
}

// FixedChar содержит информацию о фиксированном символе в шаблоне
type FixedChar struct {
	Position int  // Позиция в шаблоне
	Char     byte // Символ
}

// charToIndex конвертирует символ в индекс алфавита
func charToIndex(c byte) (int, error) {
	switch c {
	case 'A':
		return 0, nil
	case 'C':
		return 1, nil
	case 'G':
		return 2, nil
	case 'T':
		return 3, nil
	case 'N':
		return 4, nil
	default:
		return 0, fmt.Errorf("invalid character: %c", c)
	}
}

// buildTrie строит префиксное дерево для фиксированных символов
func buildTrie(fixed []FixedChar) (*Node, error) {
	root := &Node{}
	for _, fc := range fixed {
		current := root
		idx, err := charToIndex(fc.Char)
		if err != nil {
			return nil, err
		}

		if current.children[idx] == nil {
			current.children[idx] = &Node{}
		}
		current = current.children[idx]
		current.outputs = append(current.outputs, fc.Position)
		fmt.Printf("Добавлен фиксированный символ %c на позиции %d в узел %p\n",
			fc.Char, fc.Position, current)
	}
	return root, nil
}

// buildFailureLinks строит failure-ссылки для автомата
func buildFailureLinks(root *Node) {
	queue := make([]*Node, 0)

	// Инициализация первого уровня
	for i := 0; i < alphabetSize; i++ {
		if child := root.children[i]; child != nil {
			child.failure = root
			queue = append(queue, child)
			fmt.Printf("Инициализация failure для узла %p -> корень\n", child)
		}
	}

	// Построение остальных ссылок
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for i, child := range u.children {
			if child == nil {
				continue
			}

			fail := u.failure
			for fail != nil && fail.children[i] == nil {
				fail = fail.failure
			}

			if fail == nil {
				child.failure = root
			} else {
				child.failure = fail.children[i]
			}

			child.outputs = append(child.outputs, child.failure.outputs...)
			queue = append(queue, child)

			fmt.Printf("Установка failure для узла %p -> %p (outputs: %v)\n",
				child, child.failure, child.outputs)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024*10)

	scanner.Scan()
	text := scanner.Text()

	scanner.Scan()
	pattern := scanner.Text()

	scanner.Scan()
	wildcardChar := scanner.Text()[0]

	var fixed []FixedChar
	for i := 0; i < len(pattern); i++ {
		if pattern[i] != wildcardChar {
			fixed = append(fixed, FixedChar{i, pattern[i]})
		}
	}

	if len(fixed) == 0 {
		fmt.Println("Ошибка: шаблон содержит только джокеры")
		return
	}

	fmt.Println("\n=== Этап 1: Построение префиксного дерева ===")
	root, err := buildTrie(fixed)
	if err != nil {
		fmt.Println("Ошибка построения дерева:", err)
		return
	}

	fmt.Println("\n=== Этап 2: Построение failure-ссылок ===")
	buildFailureLinks(root)

	fmt.Println("\n=== Этап 3: Поиск кандидатов ===")
	counters := make(map[int]int)
	current := root
	patternLen := len(pattern)

	for textPos, c := range text {
		idx, err := charToIndex(byte(c))
		if err != nil {
			continue
		}

		// Переход по failure-ссылкам
		for current != nil && current.children[idx] == nil {
			current = current.failure
		}

		if current == nil {
			current = root
		} else {
			current = current.children[idx]
		}

		fmt.Printf("Обработан символ %c на позиции %d, текущий узел: %p\n",
			c, textPos+1, current)

		// Обновление счетчиков
		for _, p := range current.outputs {
			i := textPos - p
			if i >= 0 && i+patternLen <= len(text) {
				counters[i]++
				fmt.Printf("Найден кандидат на позиции %d (счетчик: %d)\n", i+1, counters[i])
			}
		}
	}

	fmt.Println("\n=== Этап 4: Проверка кандидатов ===")
	var results []int
	required := len(fixed)
	for i, cnt := range counters {
		if cnt == required {
			valid := true
			for _, f := range fixed {
				if i+f.Position >= len(text) || text[i+f.Position] != f.Char {
					valid = false
					break
				}
			}
			if valid {
				results = append(results, i+1)
				fmt.Printf("Подтверждено вхождение на позиции %d\n", i+1)
			}
		}
	}

	fmt.Println("\n=== Результаты ===")
	sort.Ints(results)
	for _, pos := range results {
		fmt.Println(pos)
	}
}