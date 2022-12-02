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
		switch round[0] {
		case "A":
			if round[1] == "X" {
				myScore += scoring["Lose"] + scoring["Z"]
			} else if round[1] == "Y" {
				myScore += scoring["Draw"] + scoring["X"]
			} else {
				myScore += scoring["Win"] + scoring["Y"]
			}
		case "B":
			if round[1] == "X" {
				myScore += scoring["Lose"] + scoring["X"]
			} else if round[1] == "Y" {
				myScore += scoring["Draw"] + scoring["Y"]
			} else {
				myScore += scoring["Win"] + scoring["Z"]
			}
		case "C":
			if round[1] == "X" {
				myScore += scoring["Lose"] + scoring["Y"]
			} else if round[1] == "Y" {
				myScore += scoring["Draw"] + scoring["Z"]
			} else {
				myScore += scoring["Win"] + scoring["X"]
			}
		}

	}

	readFile.Close()
	fmt.Println(myScore)
}
