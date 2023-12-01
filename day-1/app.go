package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.Open("resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("\\d")

	sums := 0

	for scanner.Scan() {
		matches := r.FindAllString(scanner.Text(), -1)
		num, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		sums += num
	}

	fmt.Println("Part one:", sums)
}

func partTwo() {
	file, err := os.Open("resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rep := strings.NewReplacer(
		"zero", "ze0ro",
		"one", "on1e",
		"two", "tw2o",
		"three", "th3ree",
		"four", "fo4ur",
		"five", "fi5ve",
		"six", "si6x",
		"seven", "se7ven",
		"eight", "ei8ght",
		"nine", "ni9ne",
	)
	rex, _ := regexp.Compile("\\d")
	sums := 0

	for scanner.Scan() {
		old := ""
		replaced := rep.Replace(scanner.Text())

		for old != replaced {
			old = replaced
			replaced = rep.Replace(old)
		}

		matches := rex.FindAllString(replaced, -1)
		num, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		sums += num
	}

	fmt.Println("Part two:", sums)
}

func main() {
	partOne()
	partTwo()
}
