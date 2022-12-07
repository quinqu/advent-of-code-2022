package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	stacks := map[int]*Stack{
		1: {stack: []string{"W", "L", "S"}},
		2: {stack: []string{"Q", "N", "T", "J"}},
		3: {stack: []string{"J", "F", "H", "C", "S"}},
		4: {stack: []string{"B", "G", "N", "W", "M", "R", "T"}},
		5: {stack: []string{"B", "Q", "H", "D", "S", "L", "R", "T"}},
		6: {stack: []string{"L", "R", "H", "F", "V", "B", "J", "M"}},
		7: {stack: []string{"M", "J", "N", "R", "W", "D"}},
		8: {stack: []string{"J", "D", "N", "H", "F", "T", "Z", "B"}},
		9: {stack: []string{"T", "F", "B", "N", "Q", "L", "H"}},
	}

	// stacks := map[int]*Stack{
	// 	1: {stack: []string{"N", "Z"}},
	// 	2: {stack: []string{"D", "C", "M"}},
	// 	3: {stack: []string{"P"}},

	// }
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//var curr int
	for scanner.Scan() {
		for k, v := range stacks {
			fmt.Printf(" stack %v : %v\n\n",k, v)
	
		}
		line := scanner.Text()

		parts := strings.Split(line, " ")

		quantity, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		from, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(parts[5])
		if err != nil {
			log.Fatal(err)
		}

		curr := []string{}
		for i := 0; i < quantity; i++ {
			fromStack := stacks[from]
			curr = append(curr, fromStack.Pop())
			// toStack := stacks[to]
			// toStack.Push(val)
			// stacks[to] = toStack
			// stacks[from] = fromStack

		}
		toStack := stacks[to].stack
		merged := append(curr, toStack...)
		stacks[to].stack = merged

		fmt.Printf("move %v from %v to %v\n", quantity, from, to)

	}
	for k, v := range stacks {
		fmt.Printf("k: %v %v\n\n",k,  v)

	}
}

type Stack struct {
	stack []string
}

func (s *Stack) Push(val string) {
	curr := s.stack
	s.stack = []string{val}
	s.stack = append(s.stack, curr...)
}

func (s *Stack) Pop() string {
	v := s.stack[0]
	s.stack = s.stack[1:]
	return v
}
