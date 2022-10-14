package main

/* fix the deep_copy function
 *
 */

import (
	"fmt"
	"reflect"
	"sort"
)

var (
	SortByNumberSequence = "SortByNumberSequence"
)

func main() {
	input := []int{534, 127, 86, 419, 94, 7, 351, 25, 912, 464}
	{

		slice := make([]int, len(input))
		copy(slice, input)
		var sequence []int = sortSequence(slice, SortByNumberSequence)
		var expectedSequence = []int{7, 25, 86, 94, 127, 351, 419, 464, 534, 912}
		if ok := reflect.DeepEqual(sequence, expectedSequence); !ok {
			fmt.Printf("assert sequence, expected %v, got %v\n", expectedSequence, sequence)
		} else {
			fmt.Printf("%v\n", sequence) // [7 25 86 94 127 351 419 464 534 912].
		}
	}
	{
		slice := make([]int, 5)
		copy(slice, input)
		// FIXME: 🔨 fix this issue!!
		var sequence []int = sortSequence(slice, SortByNumberSequence)
		var expectedSequence = []int{86, 94, 127, 419, 534}
		if ok := reflect.DeepEqual(sequence, expectedSequence); !ok {
			fmt.Printf("assert sequence, expected %v, got %v\n", expectedSequence, sequence)
		} else {
			fmt.Printf("%v\n", sequence) // [86 94 127 419 534]
		}
	}
}

func sortSequence(slice []int, mode string) []int {
	switch mode {
	case SortByNumberSequence:
		sort.Ints(slice)
		return slice
	}

	panic(fmt.Sprintf("unsupported mode '%s'", mode))
}
