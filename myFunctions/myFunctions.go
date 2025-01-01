package myFunctions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetScannerFromFile(fileName string, splitfunc bufio.SplitFunc) (*bufio.Scanner, *os.File) {
	f, err := os.Open(fileName)
	Check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(splitfunc)
	return scanner, f
}

func GetScannerFromString(scannedLine string, splitfunc bufio.SplitFunc) *bufio.Scanner {
	scanner := bufio.NewScanner(strings.NewReader(scannedLine))
	scanner.Split(splitfunc)
	return scanner
}

func GetIntFromWordScanner(wordscanner *bufio.Scanner) int {
	return GetIntFromString(wordscanner.Text())
}

func GetIntArrayFromWordScanner(wordscanner *bufio.Scanner) []int {
	var arr []int
	for wordscanner.Scan() {
		arr = append(arr, GetIntFromString(wordscanner.Text()))
	}
	return arr
}

func GetIntFromString(scannedLine string) int {
	i, e := strconv.Atoi(scannedLine)
	Check(e)
	return i
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
//	defer timer("sum")()
func FuncTimer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
