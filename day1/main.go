package main

import (
	"aoc2024/internal/read"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func makeLists() ([]int, []int) {
	var leftList []int
	var rightList []int

	origList := read.ReadStrArrayByLine("./input.txt")

	for _, line := range origList {
		words := read.StrToWordArray(line)
		left, _ := strconv.Atoi(words[0])
		right, _ := strconv.Atoi(words[len(words)-1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}

func part1() {
	var sum int
	leftList, rightList := makeLists()
	//fmt.Printf("%v\n", leftList)
	//fmt.Printf("%v\n", rightList)

	//making assumption the lists are the same length
	for i := 0; i < len(leftList); i++ {
		temp := leftList[i] - rightList[i]
		if temp < 0 {
			temp = -temp
		}
		sum += temp
	}

	fmt.Printf("Sum: %v\n", sum)
}

func part2() {
	var sum int
	rightMap := make(map[int]int)
	leftList, rightList := makeLists()

	for i := 0; i < len(rightList); i++ {
		rightMap[rightList[i]] += 1
	}
	for _, item := range leftList {
		sum += item * rightMap[item]
	}

	fmt.Printf("Sum: %v\n", sum)
}

func main() {
	p1b := time.Now()
	part1()
	mid := time.Now()
	part2()
	p2a := time.Now()
	part1Time := mid.Sub(p1b)
	part2Time := p2a.Sub(mid)
	fmt.Printf("Part 1 Time: %dμs\nPart 2 Time: %dμs\n", part1Time.Microseconds(), part2Time.Microseconds())
}
