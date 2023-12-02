package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func partOne() {
	file, err := os.Open("resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	maxGame := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	gameNumRex := regexp.MustCompile("Game (\\d+)")
	colourCountsRex := regexp.MustCompile("(\\d+) ([a-z]+)")

	sum := 0

	for scanner.Scan() {
		gameNum, _ := strconv.Atoi(gameNumRex.FindStringSubmatch(scanner.Text())[1])
		gamePossible := true

		colourCountsMatches := colourCountsRex.FindAllStringSubmatch(scanner.Text(), -1)

		for i := range colourCountsMatches {
			count, _ := strconv.Atoi(colourCountsMatches[i][1])
			colour := colourCountsMatches[i][2]

			if maxGame[colour] < count {
				gamePossible = false
				break
			}
		}

		if gamePossible {
			sum += gameNum
		}
	}

	fmt.Println("Game 1:", sum)
}

func partTwo() {
	file, err := os.Open("resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	colourCountsRex := regexp.MustCompile("(\\d+) ([a-z]+)")

	sum := 0

	for scanner.Scan() {
		minGame := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		colourCountsMatches := colourCountsRex.FindAllStringSubmatch(scanner.Text(), -1)

		for i := range colourCountsMatches {
			count, _ := strconv.Atoi(colourCountsMatches[i][1])
			colour := colourCountsMatches[i][2]

			if minGame[colour] < count {
				minGame[colour] = count
			}
		}

		product := 1

		for _, value := range minGame {
			product *= value
		}

		sum += product
	}

	fmt.Println("Game 2:", sum)
}

func main() {
	partOne()
	partTwo()
}
