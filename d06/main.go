package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type dataT struct {
	bStr []string
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	var pos1, pos2 int
	for _, v := range d.bStr {
		pos1 = getPosition(v, 4)
		pos2 = getPosition(v, 14)
	}
	fmt.Printf("Part 1: Non-repeat of char happens at position %d\n", pos1)
	fmt.Printf("Part 2: Message start at position %d\n", pos2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func getPosition(s string, l int) int {
	var st string
	var pos int
	for i, v := range s {
		if i < l {
			st = insertRune(st, v, l)
			fmt.Println(st)
		} else {
			// check new against existing
			var found bool
			if strings.ContainsRune(st, v) {
				found = true
			} else {
				max := l - 1
				// check existing
				for ri := 0; ri < max; ri++ {
					cr := st[ri]
					ss := string(st[ri+1 : max])
					if strings.ContainsRune(ss, rune(cr)) {
						found = true
						break
					}
					fmt.Printf("%03d: %s in %s + %v\n", i, string(s[ri]), ss, found)
				}
			}
			if !found {
				pos = i + 1
				fmt.Printf("%s + %s\n", st, string(v))
				break

			}
			st = insertRune(st, v, l)
			fmt.Println(st)
		}
	}
	return pos
}

func insertRune(s string, r rune, l int) string {
	if len(s) < l-1 {
		s += string(r)
	} else {
		s = s[1 : l-1]
		s += string(r)
	}
	return s
}

func importData() (dataT, error) {
	var d dataT
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	d.bStr = append(d.bStr, arr[0])
	fmt.Printf("loaded %d data-points\n", len(d.bStr))
	return d, nil
}
