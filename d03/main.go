package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type dataT struct {
	compA, compB string
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	var score, gScore int
	var group []string
	for i, v := range d {
		score += getMisplaced(v.compA, v.compB)
		gTemp := v.compA + v.compB
		group = append(group, gTemp)
		if (i+1)%3 == 0 {
			gScore += getCommon(group)
			group = nil
		}
	}
	fmt.Println(len(d))
	fmt.Printf("Part 1: Your score in misplaced items were %d \n", score)
	fmt.Printf("Part 2: In groups of three, value of common items are %d\n", gScore)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func getCommon(g []string) int {
	var score int
	chars := []byte(g[0])
	for _, cv := range chars {
		char := string(cv)
		if strings.Contains(g[1], char) && strings.Contains(g[2], char) {
			score = calcScore(cv)
			break
		}
	}
	return score
}

func getMisplaced(a, b string) int {
	var score int
	chars := []byte(a)
	for _, cv := range chars {
		char := string(cv)
		if strings.Contains(b, char) {
			score = calcScore(cv)
			break
		}
	}
	return score
}

func calcScore(c byte) int {
	// upper case
	var score int
	if int(c) < 91 {
		score = int(c) - 64 + 26
	}
	// lower case
	if int(c) > 91 {
		score = int(c) - 96
	}
	return score
}

func importData() ([]dataT, error) {
	var d []dataT
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		// split v into two in middle
		size := len(v)
		middle := size / 2
		cA := v[0:middle]
		cB := v[middle:size]
		if len(cA) != len(cB) {
			return d, fmt.Errorf("sizes do not match: %s - %s", cA, cB)
		}
		dT := dataT{
			compA: cA,
			compB: cB,
		}
		d = append(d, dT)
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}
