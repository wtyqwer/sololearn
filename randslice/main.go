package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(genRandomSlice(9, 7))
}

func genRandomSlice(length, ignore int) []int {
	if length <= ignore {
		return nil
	}
	s := make([]int, length)
	es := make([]int, length-1)
	for i := 0; i < length; i++ {
		s[i] = i
	}
	s = append(s[:ignore], s[ignore+1:]...)
	rand.Seed(time.Now().Unix())
	curIndex := 0

	for curIndex < len(es) {
		//随机从s拿一个
		randIndex := rand.Intn(len(es) - curIndex)
		es[curIndex] = s[randIndex]
		curIndex++
		//把最后一个放到被拿出来的位置上，以防止重复
		s[randIndex] = s[len(es)-curIndex]
	}
	return es
}

func generateNums(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}

	nums := make([]int, 0)
	for len(nums) < count {
		num := rand.Intn(end-start) + start

		exist := false
		for _, v := range nums {
			if v == num {
				//if v==num{
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
