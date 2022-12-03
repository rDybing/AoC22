package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	d := []dataT{
		{
			compA: "vJrwpWtwJgWr",
			compB: "hcsFMMfFFhFp",
		},
		{
			compA: "jqHRNqRjqzjGDLGL",
			compB: "rsFMfFZSrLrFZsSL",
		},
		{
			compA: "PmmdzqPrV",
			compB: "vPwwTWBwg",
		},
		{
			compA: "wMqvLMZHhHMvwLH",
			compB: "jbvcjnnSBnvTQFn",
		},
		{
			compA: "ttgJtRGJ",
			compB: "QctTZtZT",
		},
		{
			compA: "CrZsJsPPZsGz",
			compB: "wwsLwLmpwMDw",
		},
	}

	tt := struct {
		name string
		data []dataT
		sumQ int
		sumA int
	}{
		name: "Misplaced items score",
		data: d,
		sumA: 157,
	}
	t.Run(tt.name, func(t *testing.T) {
		for _, v := range tt.data {
			tt.sumQ += getMisplaced(v.compA, v.compB)
		}
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}

func TestPart2(t *testing.T) {
	d := []dataT{
		{
			compA: "vJrwpWtwJgWr",
			compB: "hcsFMMfFFhFp",
		},
		{
			compA: "jqHRNqRjqzjGDLGL",
			compB: "rsFMfFZSrLrFZsSL",
		},
		{
			compA: "PmmdzqPrV",
			compB: "vPwwTWBwg",
		},
		{
			compA: "wMqvLMZHhHMvwLH",
			compB: "jbvcjnnSBnvTQFn",
		},
		{
			compA: "ttgJtRGJ",
			compB: "QctTZtZT",
		},
		{
			compA: "CrZsJsPPZsGz",
			compB: "wwsLwLmpwMDw",
		},
	}

	tt := struct {
		name string
		data []dataT
		sumQ int
		sumA int
	}{
		name: "Group common score",
		data: d,
		sumA: 70,
	}
	t.Run(tt.name, func(t *testing.T) {
		var group []string
		for i, v := range tt.data {
			gTemp := v.compA + v.compB
			group = append(group, gTemp)
			fmt.Printf("%d ", i+1)
			if (i+1)%3 == 0 {
				tt.sumQ += getCommon(group)
				group = nil
			}
		}
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}
