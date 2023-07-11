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

	points := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}

	for matchScanner.Scan() {
		totalScore += points[matchScanner.Text()]
	}
	fmt.Println(totalScore)
}
