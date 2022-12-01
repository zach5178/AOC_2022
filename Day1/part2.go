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
	var top []int

	for i := range elfs {
		elf := elfs[i]
		// printSlice(elf)
		var totalCal int = 0
		for j := range elf {
			totalCal += elf[j]
		}

		if len(top) == numberOfElvesToTrack {
			// fmt.Println(top[2])
			// fmt.Println(totalCal)
			if top[2] < totalCal {
				top[2] = totalCal
			}
		} else {
			top = append(top, totalCal)
		}
		sort.Slice(top, func(i, j int) bool { return top[i] > top[j] })
		totalCal = 0
	}
	var topTotal int = 0
	for i := range top {
		topTotal += top[i]
	}
	fmt.Println(topTotal)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
