package main

import (
	"fmt"
	"myFunctions"
	"sort"
)

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
