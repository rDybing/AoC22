package main

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	d := dataT{
		bStr: []string{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			"nppdvjthqldpwncqszvftbrmjlhg",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		},
	}
	tt := struct {
		name  string
		data  dataT
		posQ1 int
		posQ2 int
		posA1 []int
		posA2 []int
	}{
		name:  "Test",
		data:  d,
		posA1: []int{7, 5, 6, 10, 11},
		posA2: []int{19, 23, 23, 29, 26},
	}
	t.Run(tt.name, func(t *testing.T) {
		for i, v := range tt.data.bStr {
			fmt.Printf("Test %d\n", i)
			tt.posQ1 = getPosition(v, 4)
			tt.posQ2 = getPosition(v, 14)
			if tt.posA1[i] != tt.posQ1 {
				t.Fatalf("\nP1 %s %d: Expected: %d - got: %d\n", tt.name, i, tt.posA1[i], tt.posQ1)
			}
			if tt.posA2[i] != tt.posQ2 {
				t.Fatalf("\nP2 %s %d: Expected: %d - got: %d\n", tt.name, i, tt.posA2[i], tt.posQ2)
			}
		}
	})
}
