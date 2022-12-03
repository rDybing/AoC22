package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type dataT struct {
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()

	for _, v := range d {
	}
	//fmt.Printf("Part 1: Your score in RPS game were %d points\n", score1)
	//fmt.Printf("Part 2: Following new strategy in part 2, score is %d\n", score2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func importData() ([]dataT, error) {
	var d []dataT
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		dT := dataT{}
		d = append(d, dT)
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}
