package day1

import (
	"fmt"
	"myFunctions"
	"runtime"
	"sort"
)

func Part1(arr1 []int, arr2 []int) {
	defer myFunctions.FuncTimer("part1")()

	sort.Ints(arr1)
	sort.Ints(arr2)

	//create channel for each distance calculation

	ch := make(chan int)
	go func() {
		defer close(ch)
		for i, v1 := range arr1 {
			ch <- myFunctions.Abs(v1 - arr2[i])
		}
	}()

	//wait for all channels to complete
	//sum all distances
	//print sum
	var sum int
	for diff := range ch {
		sum += diff
	}
	fmt.Println(sum)
}

func Part2(arr1 []int, arr2 []int) {
	defer myFunctions.FuncTimer("part2")()

	// count the number of appearances of each unique integer in the array
	counts_arr1 := make(map[int]int)
	for _, v := range arr1 {
		counts_arr1[v]++
	}

	counts_arr2 := make(map[int]int)
	for _, v := range arr2 {
		counts_arr2[v]++
	}

	// for each entry in counts, look for it counts2 then calculate smimlarity as:
	// similarity = counts2[counts.key] * counts.key * counts.value

	//create channel for each similarity calculation

	ch := make(chan int)
	go func() {
		defer close(ch)
		for k1, v1 := range counts_arr1 {
			ch <- k1 * v1 * counts_arr2[k1]
		}
	}()

	//wait for all channels to complete
	//sum all distances
	//print sum
	var sum int
	for sim := range ch {
		sum += sim
	}
	fmt.Println(sum)
}

func main() {
	wordscanner, f := myFunctions.GetWordScannerFromFile("input")
	defer f.Close()

	var arr1 []int
	var arr2 []int

	for wordscanner.Scan() {

		arr1 = append(arr1, myFunctions.GetIntFromWordScanner(wordscanner))
		if wordscanner.Scan() {
			arr2 = append(arr2, myFunctions.GetIntFromWordScanner(wordscanner))
		}
	}
	f.Close()

	Part1(arr1, arr2)
	Part2(arr1, arr2)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.Sys, mem.HeapAlloc, mem.HeapSys, mem.HeapInuse, mem.HeapReleased)

}
