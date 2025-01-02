package day5

import (
	"bufio"
	"fmt"
	"myFunctions"
	"regexp"
	"sort"
	"strings"
)

//open file and read line by line
//identify type of line
//if line is a rule, create the before and after rule
//if line is an update, dispatch for processing
//procesing: add up the middle number of updates in the right order

func commaSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func searchForAnyIn(needle []int, haystack []int) (int, int) {
	for i, i_val := range needle {
		for _, j_val := range haystack {
			if i_val == j_val {
				return i, i_val
			}
		}
	}
	return -1, -1
}

func day5Part1_afterRulesViolateded(after_rules map[int][]int, update []int) (int, int, int, int) {
	for i, v := range update {
		i2, v2 := searchForAnyIn(update[i+1:], after_rules[v])
		if i2 != -1 && v2 != -1 {
			return i, v, i + 1 + i2, v2
		}
	}

	return -1, -1, -1, -1
}

func day5Part1_RulesObeyed(after_rules map[int][]int, update []int) int {
	i1, i2, i3, i4 := day5Part1_afterRulesViolateded(after_rules, update)
	if i1 == -1 && i2 == -1 && i3 == -1 && i4 == -1 {
		return (update[len(update)/2])
	}

	return 0
}

func day5Part2_Swapper(after_rules map[int][]int, update []int) int {
	i1, v1, i2, v2 := day5Part1_afterRulesViolateded(after_rules, update)

	stabilised := (i1 == -1 && v1 == -1 && i2 == -1 && v2 == -1)
	if stabilised {
		return 0
	}

	cnt := 0
	for !stabilised {
		update[i1] = v2
		update[i2] = v1

		i1, v1, i2, v2 = day5Part1_afterRulesViolateded(after_rules, update)
		stabilised = (i1 == -1 && v1 == -1 && i2 == -1 && v2 == -1)
		cnt++
		if cnt > 200 {
			panic("stuck")
		}
	}

	return (update[len(update)/2])
}

func main() {

	stringscanner, f := myFunctions.GetScannerFromFile("input_rules", bufio.ScanLines)
	defer f.Close()

	re := regexp.MustCompile(`(\d+)` + regexp.QuoteMeta("|") + `(\d+)`)
	after_rules := make(map[int][]int)
	left_side := 0
	right_side := 0

	for stringscanner.Scan() {
		if rule := re.FindStringSubmatch(stringscanner.Text()); rule != nil {
			//create the before and after rule
			left_side = myFunctions.GetIntFromString(rule[1])
			right_side = myFunctions.GetIntFromString(rule[2])

			after_rules[right_side] = append(after_rules[right_side], left_side)
		}
	}

	for k := range after_rules {
		sort.Ints(after_rules[k])
	}

	stringscanner, f = myFunctions.GetScannerFromFile("input_updates", bufio.ScanLines)
	defer f.Close()

	ch := make(chan int)
	go func() {
		defer close(ch)
		for stringscanner.Scan() {
			update := myFunctions.GetIntArrayFromWordScanner(myFunctions.GetScannerFromString(stringscanner.Text(), commaSplitFunc))
			ch <- day5Part2_Swapper(after_rules, update)
		}
	}()

	sum := 0
	for mid := range ch {
		sum += mid
	}

	fmt.Println("sum: ", sum)

}
