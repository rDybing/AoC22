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
	stacks [9]stackT
	moves  []moveT
}

type stackT struct {
	containers []rune
}

type moveT struct {
	quantity int
	from     int
	to       int
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	stack2 := d.stacks
	for _, v := range d.moves {
		var subStack stackT
		// part 1
		for q := 0; q < v.quantity; q++ {
			cargo := d.stacks[v.from-1].Pop()
			cargo2 := stack2[v.from-1].Pop()
			d.stacks[v.to-1].Push(cargo)
			subStack.Push(cargo2)
		}
		// part 2
		subStack.Reverse()
		for _, cv := range subStack.containers {
			stack2[v.to-1].Push(cv)
		}
	}
	fmt.Print("Part 1: Top of stacks contain ")
	printTopStack(d.stacks[:])
	fmt.Println()
	fmt.Print("Part 2: Top of stacks now contain ")
	printTopStack(stack2[:])
	fmt.Println()
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func printTopStack(s []stackT) {
	for _, v := range s {
		index := len(v.containers) - 1
		cargo := v.Peek(index)
		fmt.Print(string(cargo))
	}
}

func (s *stackT) Push(cargo rune) {
	s.containers = append(s.containers, cargo)
}

func (s *stackT) Pop() rune {
	top := len(s.containers) - 1
	cargo := s.containers[top]
	s.containers = s.containers[0:top]
	return cargo
}

func (s stackT) Peek(i int) rune {
	return s.containers[i]
}

func (s *stackT) Reverse() {
	size := len(s.containers)
	middle := size / 2
	for i := 0; i < middle; i++ {
		j := size - i - 1
		s.containers[i], s.containers[j] = s.containers[j], s.containers[i]
	}
}

func importData() (dataT, error) {
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return dataT{}, err
	}
	arr := strings.Split(string(f), "\n")
	var modeChange bool
	var stacks [9]stackT
	var moves []moveT
	for _, v := range arr {
		if v != "" {
			// load container stacks
			if !modeChange {
				if string(v[1]) != "1" {
					for si, sv := range v {
						if sv != ' ' && sv != '[' && sv != ']' {
							index := getStackIndex(si)
							stacks[index].containers = append(stacks[index].containers, sv)
						}
					}
				}
			}
			// load crane moves
			if modeChange {
				move := strings.Split(v, " ")
				moveTemp := moveT{
					quantity: getInt(move[1]),
					from:     getInt(move[3]),
					to:       getInt(move[5]),
				}
				moves = append(moves, moveTemp)
			}
		} else {
			modeChange = true
		}
	}
	for _, v := range stacks {
		v.Reverse()
	}
	d := dataT{
		stacks: stacks,
		moves:  moves,
	}
	fmt.Printf("loaded %d stacks and %d moves\n", len(d.stacks), len(d.moves))
	return d, nil
}

func getStackIndex(i int) int {
	i--
	if i != 0 {
		i = i / 4
	}
	return i
}

func getInt(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return out
}
