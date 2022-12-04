package main

import (
	"testing"
)

func TestPart2(t *testing.T) {
	d := []dataT{
		{
			e1s: 2,
			e1e: 4,
			e2s: 6,
			e2e: 8,
		},
		{
			e1s: 2,
			e1e: 3,
			e2s: 4,
			e2e: 5,
		},
		{
			e1s: 5,
			e1e: 7,
			e2s: 7,
			e2e: 9,
		},
		{
			e1s: 2,
			e1e: 8,
			e2s: 3,
			e2e: 7,
		},
		{
			e1s: 6,
			e1e: 6,
			e2s: 4,
			e2e: 6},
		{
			e1s: 2,
			e1e: 6,
			e2s: 4,
			e2e: 8,
		},
	}

	tt := struct {
		name string
		data []dataT
		sumQ int
		sumA int
	}{
		name: "Any overlap",
		data: d,
		sumA: 4,
	}
	t.Run(tt.name, func(t *testing.T) {
		for _, v := range tt.data {
			if overlap(v) {
				tt.sumQ++
			}
		}
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}
