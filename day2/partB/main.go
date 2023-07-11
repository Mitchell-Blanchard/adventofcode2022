package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matches, _ := os.Open("spr.txt")
	matchScanner := bufio.NewScanner(matches)

	totalScore := 0

	points := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}

	for matchScanner.Scan() {
		totalScore += points[matchScanner.Text()]
	}
	fmt.Println(totalScore)
}

// A rock			1
// B paper		2
// C scissor	3

// X lose
// Y draw
// Z win
