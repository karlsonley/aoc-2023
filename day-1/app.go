package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
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

	fmt.Println(sums)
}
