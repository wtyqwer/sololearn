package main

/* deep_copy the sequence
 *
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
)

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

type SortByNumber []int

func (s SortByNumber) Len() int      { return len(s) }
func (s SortByNumber) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByNumber) Less(i, j int) bool {
	return s[i] < s[j]
}

type SortByAscii []int

func (s SortByAscii) Len() int      { return len(s) }
func (s SortByAscii) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByAscii) Less(i, j int) bool {
	var a string = strconv.Itoa(s[i])
	var b string = strconv.Itoa(s[j])
	return a < b
}

type SortByCheckSums []int

func (s SortByCheckSums) Len() int      { return len(s) }
func (s SortByCheckSums) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByCheckSums) Less(i, j int) bool {
	var a int = checksum(s[i])
	var b int = checksum(s[j])

	if a == b {
		return s[i] < s[j]
	}
	return a < b
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
func sortSequence(slice []int, mode string) []int {
	switch mode {
	case SortByNumberSequence:
		sort.Sort(SortByNumber(slice))
		return slice
	case SortByNumberDescendingSequence:
		sort.Sort(sort.Reverse(SortByNumber(slice)))
		return slice
	case SortByAsciiSequence:
		sort.Sort(SortByAscii(slice))
		return slice
	case SortByCheckSum:
		sort.Sort(SortByCheckSums(slice))
		return slice
	}

	panic(fmt.Sprintf("unsupported mode '%s'", mode))
}
