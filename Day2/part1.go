package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fileName string = os.Args[1]
	scoring := make(map[string]int)
	scoring["X"] = 1
	scoring["Y"] = 2
	scoring["Z"] = 3
	scoring["Win"] = 6
	scoring["Draw"] = 3
	scoring["Lose"] = 0

	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var myScore int = 0
	for fileScanner.Scan() {
		var round []string = strings.Split(fileScanner.Text(), " ")
		myScore += scoring[round[1]]
		switch round[0] {
		case "A":
			if round[1] == "X" {
				myScore += scoring["Draw"]
			} else if round[1] == "Y" {
				myScore += scoring["Win"]
			} else {
				myScore += scoring["Lose"]
			}
		case "B":
			if round[1] == "X" {
				myScore += scoring["Lose"]
			} else if round[1] == "Y" {
				myScore += scoring["Draw"]
			} else {
				myScore += scoring["Win"]
			}
		case "C":
			if round[1] == "X" {
				myScore += scoring["Win"]
			} else if round[1] == "Y" {
				myScore += scoring["Lose"]
			} else {
				myScore += scoring["Draw"]
			}
		}

	}

	readFile.Close()
	fmt.Println(myScore)
}
