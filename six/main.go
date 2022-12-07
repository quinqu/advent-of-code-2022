package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// the start of a packet is indicated by a sequence of four characters that are all different.

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//var curr int
	var strArr []string
	var line string
	for scanner.Scan() {

		line = scanner.Text()
	}

	chars := strings.Split(line, "")
	for _, v := range chars {
		strArr = append(strArr, v)
	}
	start := 0
	end := 13
	for end < len(strArr) {
		m := map[string]bool{}
		fmt.Printf("current start: %v current end: %v \n\n", start, end)
		currStart := start
		currEnd := end
		for currStart <= currEnd {
			fmt.Printf(" -------- INNER start: %v  end: %v \n", currStart, currEnd)
			m[strArr[currStart]] = true
			currStart++

		}
		if len(m) == 14 {
			log.Fatal(end+1)
		}
		start += 1
		end += 1
	}
}
