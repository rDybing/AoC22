package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	calories int
}

type elfT struct {
	calSum int
}

type topT struct {
	elfKey int
	calSum int
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	elf := make(map[int]elfT)
	var top []topT
	var calSum, elfKey, elfLead int
	for _, v := range d {
		calSum += v.calories
		if v.calories == 0 {
			elfKey++
			elfTemp := elfT{
				calSum: calSum,
			}
			elf[elfKey] = elfTemp
			calSum = 0
			if elfKey == 1 {
				elfLead = 1
			} else {
				if elf[elfLead].calSum < elf[elfKey].calSum {
					elfLead = elfKey
				}
			}
			topTemp := topT{
				elfKey: elfKey,
				calSum: elf[elfKey].calSum,
			}
			top = append(top, topTemp)
		}
	}
	sort.Slice(top, func(i, j int) bool {
		return top[i].calSum > top[j].calSum
	})
	fmt.Printf("Part 1: Elf %03d have most calories with %d\n", elfLead, elf[elfLead].calSum)
	fmt.Printf("Part 2: Top 3 elfs:\n")
	var topSum int
	for i := 0; i < 3; i++ {
		fmt.Printf("\t%03d - %d\n", top[i].elfKey, top[i].calSum)
		topSum += top[i].calSum
	}
	fmt.Printf("They carry in total %d calories\n", topSum)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func importData() ([]dataT, error) {
	var d []dataT
	f, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		var cal int
		if v != "" {
			cal, _ = strconv.Atoi(v)
		}
		dT := dataT{
			calories: cal,
		}
		d = append(d, dT)
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}
