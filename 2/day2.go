package day2

import (
	"bufio"
	"fmt"
	"myFunctions"
	"sort"
)

type level struct {
	arr                    []int
	numberOfTransgressions int
}

func part2Addition(arr []int) (alternativeTrueLevels []level) {
	var slice1 []int
	var slice2 []int
	var returnSlice []level

	if sort.IntsAreSorted(arr) || sort.IsSorted(sort.Reverse(sort.IntSlice(arr))) {
		returnSlice = append(returnSlice, level{arr, 0})
		return returnSlice
	} else {

		for i := 0; i < len(arr); i++ {
			slice1 = arr[:i]

			if i == len(arr)-1 {
				slice2 = []int{}
			} else {
				slice2 = arr[i+1:]
			}

			slice3 := []int{}
			slice3 = append(slice3, slice1...)
			slice3 = append(slice3, slice2...)

			if sort.IntsAreSorted(slice3) || sort.IsSorted(sort.Reverse(sort.IntSlice(slice3))) {
				returnSlice = append(returnSlice, level{slice3, 1})
			}
		}
	}
	return returnSlice
}
func amIaTrueLevel(wordscanner *bufio.Scanner, numberofpermittedtransgressions int) bool {
	var arr []int

	for wordscanner.Scan() {
		arr = append(arr, myFunctions.GetIntFromWordScanner(wordscanner))
	}

	potentialLevels := part2Addition(arr)

	for _, leveli := range potentialLevels {

		for x := 0; x < len(leveli.arr)-1; x++ {
			v := myFunctions.Abs(leveli.arr[x] - leveli.arr[x+1])

			if !((1 <= v) && (v <= 3)) {
				leveli.numberOfTransgressions++
				if v > 0 && (x > 0 && x < len(leveli.arr)-2) {
					leveli.numberOfTransgressions++
				}
			}

		}
		if leveli.numberOfTransgressions <= numberofpermittedtransgressions {
			//we've found a true level, release memories and return
			arr = nil
			return true
		}
	}

	//release memories and return
	arr = nil
	return false
}

func main() {
	linescanner, f := myFunctions.GetScannerFromFile("input", bufio.ScanLines)
	defer f.Close()

	ch := make(chan bool)
	go func() {
		defer close(ch)
		for linescanner.Scan() {
			//Part1
			//ch <- amIaTrueLevel(myFunctions.GetWordScannerFromLine(linescanner.Text()), 0)
			//Part2
			ch <- amIaTrueLevel(myFunctions.GetScannerFromString(linescanner.Text(), bufio.ScanWords), 1)

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
