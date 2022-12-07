package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// a = rock
// b = paper
// c = scissors

// x = rock
// y = paper
// z = scissors

// rock - 1
// paper - 2
// scisscors - 3

// plus
//
// draw  +3
// won +6

func main() {

	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var myScore int

	for scanner.Scan() {

		line := scanner.Text()
		arr := strings.Split(line, " ")
		opponentMove := arr[0]
		myMove := arr[1]

		switch {
		case opponentMove == "A" && myMove == "X":
			myScore += 4
		case opponentMove == "B" && myMove == "X":
			myScore += 1
		case opponentMove == "C" && myMove == "X":
			myScore += 7

		case opponentMove == "A" && myMove == "Y":
			myScore += 8
		case opponentMove == "B" && myMove == "Y":
			myScore += 5
		case opponentMove == "C" && myMove == "Y":
			myScore += 2

		case opponentMove == "A" && myMove == "Z":
			myScore += 3
		case opponentMove == "B" && myMove == "Z":
			myScore += 9
		case opponentMove == "C" && myMove == "Z":
			myScore += 6

		}

	}
	fmt.Println(myScore)

}
