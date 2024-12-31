package day3

import (
	"bufio"
	"fmt"
	"myFunctions"
	"regexp"
	"strings"
)

func mulSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "mul("); i >= 0 {
		return i + 4, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func doSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "do()"); i >= 0 {
		return i + 4, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
func dontSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "don't()"); i >= 0 {
		return i + 7, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
func day3Part1(s string) int {
	r, _ := regexp.Compile("^([0-9]+),([0-9]+)" + regexp.QuoteMeta(")"))
	result := 0
	matches := r.FindStringSubmatch(s)
	if matches != nil {
		result = myFunctions.GetIntFromString(matches[1]) * myFunctions.GetIntFromString(matches[2])
	}
	return result
}

func day3Part2(s string) int {
	result := 0
	scanner1 := bufio.NewScanner(strings.NewReader(s))
	scanner1.Split(dontSplitFunc)
	scanner1.Scan()
	s1 := scanner1.Text()
	if len(s1) > 0 {
		scanner2 := bufio.NewScanner(strings.NewReader(s1))
		scanner2.Split(mulSplitFunc)
		for scanner2.Scan() {
			result += day3Part1(scanner2.Text())
		}
	}
	return result
}
func main() {

	scanner, f := myFunctions.GetScannerFromFile("input", doSplitFunc)
	defer f.Close()

	ch := make(chan int)
	go func() {
		defer close(ch)
		for scanner.Scan() {
			ch <- day3Part2(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}()

	sum := 0
	for mul := range ch {
		sum += mul
	}

	fmt.Println(sum)

}
