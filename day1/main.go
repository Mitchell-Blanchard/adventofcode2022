package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	calories []int
	total    int
}

var elves = make([]elf, 1)

func calcAllCalories(elves []elf) {
	for i, elf := range elves {
		total := 0
		for _, calories := range elf.calories {
			total += calories
		}

		elves[i].total = total
	}
}

func calcTopCalories(elves []elf, top int) int {
	calorieList := []int{}
	for _, elf := range elves {
		calorieList = append(calorieList, elf.total)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calorieList)))
	fmt.Println(calorieList[:top])
	total := 0
	for _, d := range calorieList[:top] {
		total += d
	}
	return total
}

func main() {

	calories, err := os.Open("./calories.txt")
	if err != nil {
		log.Fatalf("'%s'\n", err)
	}
	scanner := bufio.NewScanner(calories)

	x := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			line, _ := strconv.Atoi(scanner.Text())
			elves[x].calories = append(elves[x].calories, line)
		} else {
			x++
			elves = append(elves, elf{})
		}
	}

	calcAllCalories(elves)

	fmt.Println(calcTopCalories(elves, 3))

}
