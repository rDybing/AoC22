package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	e1s, e1e, e2s, e2e int
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	var cSum, olSum int
	for _, v := range d {
		if cover(v) {
			cSum++
		}
		if overlap(v) {
			olSum++
		}
	}
	fmt.Printf("Part 1: Cover of assignments happen with %d elfs\n", cSum)
	fmt.Printf("Part 2: Overlap of assignments happen with %d elfs\n", olSum)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func overlap(d dataT) bool {
	for r1 := d.e1s; r1 <= d.e1e; r1++ {
		for r2 := d.e2s; r2 <= d.e2e; r2++ {
			if r1 == r2 {
				return true
			}
		}
	}
	return false
}

func cover(d dataT) bool {
	var cover bool
	// elf 1 range v elf 2 range
	if d.e1s >= d.e2s && d.e1e <= d.e2e {
		cover = true
	}
	// elf 2 range v elf 1 range
	if d.e2s >= d.e1s && d.e2e <= d.e1e {
		cover = true
	}
	return cover
}

func importData() ([]dataT, error) {
	var d []dataT
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		if v != "" {
			elfs := strings.Split(v, ",")
			e1 := strings.Split(elfs[0], "-")
			e2 := strings.Split(elfs[1], "-")
			dT := dataT{
				e1s: getInt(e1[0]),
				e1e: getInt(e1[1]),
				e2s: getInt(e2[0]),
				e2e: getInt(e2[1]),
			}
			d = append(d, dT)
		}
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}

func getInt(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return out
}
