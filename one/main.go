package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var max [3]int

	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var curr int
	for scanner.Scan() {
		line := scanner.Text()
		// string to int
		var v int
		if line != "" {
			v, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
		}

		if v == 0 {
			if curr > max[2] {
				tempLargest := max[2]
				tempMiddle := max[1]
				max[0] = tempMiddle
				max[1] = tempLargest
				max[2] = curr
			} else if curr > max[1] {
				tempMiddle := max[1]
				max[0] = tempMiddle
				max[1] = curr
			} else if curr > max[0] {
				max[0] = curr
			}

			curr = 0
		} else {
			curr += v
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)

}

func sumCalories(calorieList []int) int {
	var sum int
	for _, val := range calorieList {
		sum += val
	}
	return sum
}
