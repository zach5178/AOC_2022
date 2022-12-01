package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const numberOfElvesToTrack = 3

func main() {
	var fileName string = os.Args[1]

	readFile, err := os.Open(fileName)
	var elfs []int

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var totalCal int = 0
	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		if line != "" {
			intVal, _ := strconv.Atoi(line)
			totalCal += intVal
		} else {
			elfs = append(elfs, totalCal)
			totalCal = 0
		}
	}
	elfs = append(elfs, totalCal)
	readFile.Close()
	sort.Slice(elfs, func(i, j int) bool { return i > j })
	printSlice(elfs)
	var topTotal int = 0
	for i := 0; i < numberOfElvesToTrack; i++ {
		topTotal += elfs[i]
	}
	fmt.Println(topTotal)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
