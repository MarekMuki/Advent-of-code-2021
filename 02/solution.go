package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var vertical, horizontal int = 0, 0

	for scanner.Scan() {

		line := strings.Fields(scanner.Text())
		command := line[0]
		value, _ := strconv.Atoi(line[1])

		fmt.Println(command, value)

		switch command {
		case "forward":
			horizontal += value
		case "down":
			vertical += value
		case "up":
			vertical -= value
		}
	}

	fmt.Println(vertical * horizontal)
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var horizontal, depth, aim int = 0, 0, 0

	for scanner.Scan() {

		line := strings.Fields(scanner.Text())
		command := line[0]
		value, _ := strconv.Atoi(line[1])

		fmt.Println(command, value)

		switch command {
		case "forward":
			horizontal += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	fmt.Println(depth * horizontal)
}

func main() {

	part2()
}
