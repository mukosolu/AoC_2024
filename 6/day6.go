/*
create 2 maps:
1 associating columns with objects to rows
Another associating rows with objects to columns
sort the stored array values

note the coordinates of the starting point.
for each "turn": on first turn look "up" in that column, where is the next object? i,e in the map of column to row
search map[column] for the next row with an object in that direction.
row identified? (count nb of rows between starting point and identified row)
time for next turn.
change direction to be 90deg right turn
search map[row] for the next column with an object in that direction - if we need to stay in the row
do the same with map[column]] if we need to stay in that column
object identified? (count nb of rows/cols between starting point and identified object)
time for next turn: are we done? i.e. is there no nearest row or column returned?
*/

package main

import (
	"bufio"
	"fmt"
	"myFunctions"
	"regexp"
	"sort"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type position struct {
	row int
	col int
}

type guardPosition struct {
	pos position
	dir direction
}

func findClosestHigher(arr []int, value int) int {

	for _, v := range arr {
		if v > value {
			return v
		}
	}
	return -1
}

func findClosestLower(arr []int, value int) int {

	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		if v < value {
			return v
		}
	}
	return -1
}

func goToNextObject(guardPosition *guardPosition, objRowToColMap map[int][]int, objColToRowMap map[int][]int, stepCounter *map[position]int, maxRow int, maxCol int) guardPosition {
	var newvalue int

	switch guardPosition.dir {
	case up:
		newvalue = findClosestLower(objColToRowMap[guardPosition.pos.col], guardPosition.pos.row)
		if newvalue != -1 {
			for i := newvalue + 1; i < guardPosition.pos.row; i++ {
				(*stepCounter)[position{i, guardPosition.pos.col}]++
			}
			guardPosition.pos.row = newvalue + 1
			guardPosition.dir = right
			return *guardPosition
		} else {
			for i := 0; i < guardPosition.pos.row; i++ {
				(*stepCounter)[position{i, guardPosition.pos.col}]++
			}
		}

	case down:
		newvalue = findClosestHigher(objColToRowMap[guardPosition.pos.col], guardPosition.pos.row)
		if newvalue != -1 {
			for i := guardPosition.pos.row + 1; i < newvalue; i++ {
				(*stepCounter)[position{i, guardPosition.pos.col}]++
			}
			guardPosition.pos.row = newvalue - 1
			guardPosition.dir = left
			return *guardPosition
		} else {
			for i := guardPosition.pos.row + 1; i < maxRow; i++ {
				(*stepCounter)[position{i, guardPosition.pos.col}]++
			}
		}

	case left:
		newvalue = findClosestLower(objRowToColMap[guardPosition.pos.row], guardPosition.pos.col)
		if newvalue != -1 {
			for i := newvalue + 1; i < guardPosition.pos.col; i++ {
				(*stepCounter)[position{guardPosition.pos.row, i}]++
			}
			guardPosition.pos.col = newvalue + 1
			guardPosition.dir = up
			return *guardPosition
		} else {
			for i := 0; i < guardPosition.pos.col; i++ {
				(*stepCounter)[position{guardPosition.pos.row, i}]++
			}
		}

	case right:
		newvalue = findClosestHigher(objRowToColMap[guardPosition.pos.row], guardPosition.pos.col)
		if newvalue != -1 {
			for i := guardPosition.pos.col + 1; i < newvalue; i++ {
				(*stepCounter)[position{guardPosition.pos.row, i}]++
			}
			guardPosition.pos.col = newvalue - 1
			guardPosition.dir = down
			return *guardPosition
		} else {
			for i := guardPosition.pos.col + 1; i < maxCol; i++ {
				(*stepCounter)[position{guardPosition.pos.row, i}]++
			}
		}

	}

	guardPosition.pos.row = -1
	guardPosition.pos.col = -1
	return *guardPosition
}

func main() {
	scanner, f := myFunctions.GetScannerFromFile("input", bufio.ScanLines)
	defer f.Close()

	objRowToColMap := make(map[int][]int)
	objColToRowMap := make(map[int][]int)
	stepCounter := make(map[position]int)

	row := 0
	re_hash := regexp.MustCompile(regexp.QuoteMeta("#"))
	re_starter := regexp.MustCompile(regexp.QuoteMeta("^"))
	guardCurrentPosition := guardPosition{position{0, 0}, up}

	var s string
	for scanner.Scan() {
		//read line from file
		s = scanner.Text()

		//identufy objects
		objectMatches := re_hash.FindAllStringIndex(s, -1)
		for _, v := range objectMatches {
			objRowToColMap[row] = append(objRowToColMap[row], v[0])
			objColToRowMap[v[0]] = append(objColToRowMap[v[0]], row)
		}

		//sort the stored array values
		sort.Ints(objRowToColMap[row])
		sort.Ints(objColToRowMap[row])

		//look for starter
		starterMatch := re_starter.FindStringIndex(s)
		if starterMatch != nil {
			guardCurrentPosition = guardPosition{position{row, starterMatch[0]}, up}
		}

		row++
	}
	maxRow := row
	maxCol := len(s)

	/*fmt.Println("objRowToColMap: ")
	fmt.Println(objRowToColMap)
	fmt.Println("objColToRowMap: ")
	fmt.Println(objColToRowMap)
	fmt.Println("------")*/

	stop := (guardCurrentPosition.pos.row == -1)
	if !stop {
		stepCounter[guardCurrentPosition.pos]++
	}

	for !stop {
		//fmt.Println("guard moved from: ", guardCurrentPosition)
		guardCurrentPosition = goToNextObject(&guardCurrentPosition, objRowToColMap, objColToRowMap, &stepCounter, maxRow, maxCol)
		//fmt.Println("guard moved to: ", guardCurrentPosition)
		//fmt.Println(stepCounter)
		stop = (guardCurrentPosition.pos.row == -1)
	}

	fmt.Println("nb of steps: ", len(stepCounter))
}
