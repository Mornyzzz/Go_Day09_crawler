package main

import (
	"reflect"
	"testing"
)

func TestSleepSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{5, 2, 7, 1, 4, 3}, []int{1, 2, 3, 4, 5, 7}},
		{[]int{}, []int{}},
		{[]int{9, 5, 2, 6, 8}, []int{2, 5, 6, 8, 9}},
	}

	for _, test := range tests {
		result := make([]int, 0)
		c := sleepSort(test.input)
		for val := range c {
			result = append(result, val)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestSleepSortEdgeCases(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}},
		{[]int{100, 10, 1000, 1}, []int{1, 10, 100, 1000}},
		{[]int{2, 1, 3, 2, 1}, []int{1, 1, 2, 2, 3}},
	}

	for _, test := range tests {
		result := make([]int, 0)
		c := sleepSort(test.input)
		for val := range c {
			result = append(result, val)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
