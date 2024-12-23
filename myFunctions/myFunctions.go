package myFunctions

import (
	"bufio"
	"os"
	"strconv"
)

func GetWordScannerFromFile(fileName string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(fileName)
	Check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	return scanner, f
}

func GetIntFromWordScanner(wordscanner *bufio.Scanner) int {
	i, e := strconv.Atoi(wordscanner.Text())
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
