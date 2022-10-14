package main

///* fix the sort_imp function
// *  原题
// */
//
//import (
//	"fmt"
//	"sort_imp"
//	"reflect"
//)
//
//var (
//	SortByNumberSequence = "SortByNumberSequence"
//)
//
//func main() {
//	input := []int{ 534, 127, 86, 419, 94, 7, 351, 25, 912, 464 }
//	{
//		slice := input
//		var sequence []int = sortSequence(slice, SortByNumberSequence)
//		var expectedSequence = []int{ 7, 25, 86, 94, 127, 351, 419, 464, 534, 912 }
//		if ok := reflect.DeepEqual(sequence, expectedSequence); !ok {
//			fmt.Printf("assert sequence, expected %v, got %v\n", expectedSequence, sequence)
//		} else {
//			fmt.Printf("%v\n", sequence) // [7 25 86 94 127 351 419 464 534 912]
//		}
//	}
//	{
//		slice := input[:5]
//		// FIXME: 🔨 fix this issue!!
//		var sequence []int = sortSequence(slice, SortByNumberSequence)
//		var expectedSequence = []int{ 86, 94, 127, 419, 534 }
//		if ok := reflect.DeepEqual(sequence, expectedSequence); !ok {
//			fmt.Printf("assert sequence, expected %v, got %v\n", expectedSequence, sequence)
//		} else {
//			fmt.Printf("%v\n", sequence) // [86 94 127 419 534]
//		}
//	}
//}
//
//
//func sortSequence(slice []int, mode string) []int {
//	switch (mode) {
//	case SortByNumberSequence:
//		sort_imp.Ints(slice)
//		return slice
//	}
//
//	panic(fmt.Sprintf("unsupported mode '%s'", mode))
//}
