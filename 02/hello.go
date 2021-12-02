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
			horizontal = horizontal + value
		case "down":
			vertical = vertical + value
		case "up":
			vertical = vertical - value
		}
	}

	fmt.Println(vertical * horizontal)
}
