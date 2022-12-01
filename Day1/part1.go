package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var fileName string = os.Args[1]

	readFile, err := os.Open(fileName)
	var elfs [][]int

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var elf []int
	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		if line != "" {
			intVal, _ := strconv.Atoi(line)
			elf = append(elf, intVal)
		} else {
			elfs = append(elfs, elf)
			elf = nil
		}
	}
	elfs = append(elfs, elf)
	readFile.Close()
	var largest int = 0
	for i := range elfs {
		elf := elfs[i]
		var totalCal int = 0
		for j := range elf {
			totalCal += elf[j]
		}
		if totalCal > largest {
			largest = totalCal
		}
		totalCal = 0
	}

	fmt.Println(largest)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
