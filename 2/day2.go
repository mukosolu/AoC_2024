package main

import (
	"bufio"
	"fmt"
	"myFunctions"
	"sort"
)

func amIaTrueLevel(wordscanner *bufio.Scanner) bool {
	var arr []int

	for wordscanner.Scan() {
		arr = append(arr, myFunctions.GetIntFromWordScanner(wordscanner))
	}
	result := sort.IntsAreSorted(arr) || sort.IsSorted(sort.Reverse(sort.IntSlice(arr)))

	for x := 0; x < len(arr)-1; x++ {
		v := myFunctions.Abs(arr[x] - arr[x+1])
		result = result && ((1 <= v) && (v <= 3))
	}
	arr = nil
	return result
}

func main() {
	linescanner, f := myFunctions.GetLineScannerFromFile("input")
	defer f.Close()

	ch := make(chan bool)
	go func() {
		defer close(ch)
		for linescanner.Scan() {
			ch <- amIaTrueLevel(myFunctions.GetWordScannerFromLine(linescanner.Text()))
		}
	}()

	var sum int
	for levelCheck := range ch {
		if levelCheck {
			sum += 1
		}
	}

	fmt.Println(sum)
}
