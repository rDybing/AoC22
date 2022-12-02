package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type dataT struct {
	p1, p2 string
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	points := fillPoints()
	weapon := selectRPS()
	var score1, score2 int
	for _, v := range d {
		score1 += points[v.p2] + points[v.p1+v.p2]
		score2 += points[weapon[v.p1+v.p2]] + points[v.p1+weapon[v.p1+v.p2]]
	}
	fmt.Printf("Part 1: Your score in RPS game were %d points\n", score1)
	fmt.Printf("Part 2: Following new strategy in part 2, score is %d\n", score2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func fillPoints() map[string]int {
	p := make(map[string]int)
	// points on selected 'weapon'
	p["A"] = 1 // p1 rock
	p["B"] = 2 // p1 paper
	p["C"] = 3 // p1 scissor
	p["X"] = 1 // p2 rock
	p["Y"] = 2 // p2 paper
	p["Z"] = 3 // p2 scissor
	// win, draw, lost player 2 truth table
	p["AX"] = 3 // r+r draw
	p["AY"] = 6 // r+p win
	p["AZ"] = 0 // r+s lost
	p["BX"] = 0 // p+r lost
	p["BY"] = 3 // p+p draw
	p["BZ"] = 6 // p+s win
	p["CX"] = 6 // s+r win
	p["CY"] = 0 // s+p lost
	p["CZ"] = 3 // s+s draw
	return p
}

func selectRPS() map[string]string {
	// X = loose, Y = Draw, Z = Win
	s := make(map[string]string)
	s["AX"] = "Z" // loss, r+s
	s["AY"] = "X" // draw, r+r
	s["AZ"] = "Y" // win, r+p
	s["BX"] = "X" // loss, p+r
	s["BY"] = "Y" // draw, p+p
	s["BZ"] = "Z" // win, p+s
	s["CX"] = "Y" // loss, s+p
	s["CY"] = "Z" // draw, s+s
	s["CZ"] = "X" // win, s+r
	return s
}

func importData() ([]dataT, error) {
	var d []dataT
	f, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		if v != "" {
			rps := strings.Split(v, " ")
			dT := dataT{
				p1: rps[0],
				p2: rps[1],
			}
			d = append(d, dT)
		}
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}
