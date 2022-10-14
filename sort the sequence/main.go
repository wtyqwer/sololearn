package main

/* deep_copy the sequence
 *  solve 策略模式
 */

import (
	"fmt"
	"sort"
	"strconv"
)

var (
	SortByNumberSequence           = "SortByNumberSequence"
	SortByNumberDescendingSequence = "SortByNumberDescendingSequence"
	SortByAsciiSequence            = "SortByAsciiSequence"
	SortByCheckSum                 = "SortByCheckSum"
	sortTool                       *SortTool
	sortMap                        = map[string]Algorithm{
		"SortByNumberSequence":           &SortByNumber{},
		"SortByNumberDescendingSequence": &SortByNumberDescending{},
		"SortByAsciiSequence":            &SortByAscii{},
		"SortByCheckSum":                 &SortByCheckSums{},
	}
)

func init() {
	unknownSort := &UnknownSort{}
	sortTool = initSortTool(unknownSort)
}

func main() {

	{
		slice := []int{534, 127, 86, 419, 94, 7, 351, 25, 912, 464}
		var sequence []int = sortSequence(slice, SortByNumberSequence)
		fmt.Printf("%v\n", sequence)
	}
	{
		slice := []int{534, 127, 86, 419, 94, 7, 351, 25, 912, 464}
		var sequence []int = sortSequence(slice, SortByNumberDescendingSequence)
		fmt.Printf("%v\n", sequence)
	}
	{
		slice := []int{534, 127, 86, 419, 94, 7, 351, 25, 912, 464}
		var sequence []int = sortSequence(slice, SortByAsciiSequence)
		fmt.Printf("%v\n", sequence)
	}
	{
		slice := []int{534, 127, 86, 419, 94, 7, 351, 25, 912, 464}
		var sequence []int = sortSequence(slice, SortByCheckSum)
		fmt.Printf("%v\n", sequence)
	}
	{
		slice := []int{534, 127, 86, 419, 94, 7, 351, 25, 912, 464}
		var sequence []int = sortSequence(slice, "test")
		fmt.Printf("%v\n", sequence)
	}
}

type Algorithm interface {
	algo(slice []int) []int
}

type UnknownSort struct {
}

func (s *UnknownSort) algo(slice []int) []int {

	panic(fmt.Sprintf("unsupported mode "))
}

type SortByNumber struct {
}

func (s *SortByNumber) algo(slice []int) []int {
	sort.Ints(slice)
	return slice
}

type SortByNumberDescending struct {
}

func (s *SortByNumberDescending) algo(slice []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return slice
}

type SortByAscii struct {
}

func (s *SortByAscii) algo(slice []int) []int {

	sort.Slice(slice,
		func(i, j int) bool {
			var a string = strconv.Itoa(slice[i])
			var b string = strconv.Itoa(slice[j])
			return a < b
		})
	return slice
}

type SortByCheckSums struct {
}

func (s *SortByCheckSums) algo(slice []int) []int {
	sort.Slice(slice,
		func(i, j int) bool {
			var a int = checksum(slice[i])
			var b int = checksum(slice[j])

			if a == b {
				return slice[i] < slice[j]
			}
			return a < b
		})
	return slice
}

type SortTool struct {
	algorithm Algorithm
}

func initSortTool(a Algorithm) *SortTool {

	return &SortTool{
		algorithm: a,
	}
}

func (c *SortTool) setAlgorithm(a Algorithm) {
	c.algorithm = a
}

// TODO: 🔨 Refactor this function!!
func sortSequence(slice []int, mode string) []int {
	if algo, ok := sortMap[mode]; ok {
		sortTool.setAlgorithm(algo)
		sortTool.algorithm.algo(slice)
		return slice
	}

	panic(fmt.Sprintf("unsupported mode '%s'", mode))
}

func checksum(number int) int {
	var digits string = strconv.Itoa(number)

	var value int = 0
	for _, v := range digits {
		d := int(v - '0')
		value += d
	}
	return value
}

//// TODO: 🔨 Refactor this function!!
//func sortSequence(slice []int, mode string) []int {
//	switch mode {
//	case SortByNumberSequence:
//
//		sort.Ints(slice)
//		return slice
//	case SortByNumberDescendingSequence:
//		sort.Sort(sort.Reverse(sort.IntSlice(slice)))
//		return slice
//	case SortByAsciiSequence:
//		sort.Slice(slice,
//			func(i, j int) bool {
//				var a string = strconv.Itoa(slice[i])
//				var b string = strconv.Itoa(slice[j])
//				return a < b
//			})
//		return slice
//	case SortByCheckSum:
//		sort.Slice(slice,
//			func(i, j int) bool {
//				var a int = checksum(slice[i])
//				var b int = checksum(slice[j])
//
//				if a == b {
//					return slice[i] < slice[j]
//				}
//				return a < b
//			})
//		return slice
//	}
//
//	panic(fmt.Sprintf("unsupported mode '%s'", mode))
//}
