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
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {

		line := scanner.Text()
		arr := strings.Split(line, ",")

		fmt.Println(arr)
		arr1 := strings.Split(arr[0], "-")
		fmt.Println(arr1)

		arr2 := strings.Split(arr[1], "-")
		fmt.Println(arr2)

		r1Start, err := strconv.Atoi(arr1[0])
		if err != nil {
			log.Fatal(err)
		}
		r1End, err := strconv.Atoi(arr1[1])
		if err != nil {
			log.Fatal(err)
		}
		r2Start, err := strconv.Atoi(arr2[0])
		if err != nil {
			log.Fatal(err)
		}
		r2End, err := strconv.Atoi(arr2[1])
		if err != nil {
			log.Fatal(err)
		}
		if overlaps([2]int{r1Start, r1End}, [2]int{r2Start, r2End}) {
			count += 1
		}
	}
	fmt.Println(count)
}

func overlaps(range1, range2 [2]int) bool {

	if range1[0] > range2[0] {
		tmp := range1
		range1 = range2
		range2 = tmp
	}
	if range1[1] >= range2[0] {
		return true
	}

	// Part One: 
	// if (range1[0] <= range2[0] && range1[1] >= range2[1]) || (range2[0] <= range1[0] && range2[1] >= range1[1]){ // overlaps
	// 	return true
	// }
	return false
}
