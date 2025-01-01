package day4

import (
	"bufio"
	"fmt"
	"myFunctions"
)

func checkforXsequenceOnly(twoDimRuneArray [][]rune, row, col int, sequence1 [3]rune, sequence2 [3]rune) int {
	result := 0
	foundrdiagonal := false
	ldiagcol := 0
	rdiagonalsequenceToSearchFor := make([]rune, 3)
	ldiagonalsequenceToSearchFor := make([]rune, 3)

	if twoDimRuneArray[row][col] == sequence1[0] {
		rdiagonalsequenceToSearchFor = sequence1[:] //rdiagonalsequenceToSearchFor = sequence1
	} else if twoDimRuneArray[row][col] == sequence2[0] {
		rdiagonalsequenceToSearchFor = sequence2[:]
	}

	//check the rdiagonal
	//careful of out of bounds
	if (col+2 < len(twoDimRuneArray[row+2])) && (col+1 < len(twoDimRuneArray[row+1])) {
		if twoDimRuneArray[row+1][col+1] == rdiagonalsequenceToSearchFor[1] {
			if twoDimRuneArray[row+2][col+2] == rdiagonalsequenceToSearchFor[2] {
				//set coordinate for ldiagonal search
				foundrdiagonal = true
			}
		}
	}

	//check the ldiagonal
	//careful of out of bounds
	if foundrdiagonal {
		ldiagcol = col + 2
		if twoDimRuneArray[row][ldiagcol] == sequence1[0] {
			ldiagonalsequenceToSearchFor = sequence1[:]
		} else if twoDimRuneArray[row][ldiagcol] == sequence2[0] {
			ldiagonalsequenceToSearchFor = sequence2[:]
		}
	}
	if (len(ldiagonalsequenceToSearchFor) != 0) && (ldiagcol-2 >= 0) && (ldiagcol-2 < len(twoDimRuneArray[row+2])) && (ldiagcol-1 < len(twoDimRuneArray[row+1])) {
		if twoDimRuneArray[row+1][ldiagcol-1] == ldiagonalsequenceToSearchFor[1] {
			if twoDimRuneArray[row+2][ldiagcol-2] == ldiagonalsequenceToSearchFor[2] {
				result++
			}
		}
	}

	return result
}

func checkforsequence(twoDimRuneArray [][]rune, row, col int, sequence [4]rune) int {
	result := 0
	if twoDimRuneArray[row][col] == sequence[0] {
		//check the ldiagonal
		//careful of out of bounds
		if (col-3 >= 0) && (col-3 < len(twoDimRuneArray[row+3])) && (col-2 < len(twoDimRuneArray[row+2])) && (col-1 < len(twoDimRuneArray[row+1])) {
			if twoDimRuneArray[row+1][col-1] == sequence[1] {
				if twoDimRuneArray[row+2][col-2] == sequence[2] {
					if twoDimRuneArray[row+3][col-3] == sequence[3] {
						result++
					}
				}
			}
		}
		//check the horizontal -forward
		//careful of out of bounds
		if col+3 < len(twoDimRuneArray[row]) {
			if twoDimRuneArray[row][col+1] == sequence[1] {
				if twoDimRuneArray[row][col+2] == sequence[2] {
					if twoDimRuneArray[row][col+3] == sequence[3] {
						result++
					}
				}
			}
		}
		//check the rdiagonal
		//careful of out of bounds
		if (col+3 < len(twoDimRuneArray[row+3])) && (col+2 < len(twoDimRuneArray[row+2])) && (col+1 < len(twoDimRuneArray[row+1])) {
			if twoDimRuneArray[row+1][col+1] == sequence[1] {
				if twoDimRuneArray[row+2][col+2] == sequence[2] {
					if twoDimRuneArray[row+3][col+3] == sequence[3] {
						result++
					}
				}
			}
		}
		//check the vertical
		//careful of out of bounds
		if (col < len(twoDimRuneArray[row+3])) && (col < len(twoDimRuneArray[row+2])) && (col < len(twoDimRuneArray[row+1])) {
			if twoDimRuneArray[row+1][col] == sequence[1] {
				if twoDimRuneArray[row+2][col] == sequence[2] {
					if twoDimRuneArray[row+3][col] == sequence[3] {
						result++
					}
				}
			}
		}
	}
	return result
}

func day4Part1(twoDimRuneArray [][]rune) int {
	result := 0
	row := 0

	for col := range twoDimRuneArray[row] {
		result += checkforsequence(twoDimRuneArray, row, col, [4]rune{'X', 'M', 'A', 'S'})
		result += checkforsequence(twoDimRuneArray, row, col, [4]rune{'S', 'A', 'M', 'X'})
	}

	return result
}
func day4Part2(twoDimRuneArray [][]rune) int {
	result := 0
	row := 0

	for col := range twoDimRuneArray[row] {
		result += checkforXsequenceOnly(twoDimRuneArray, row, col, [3]rune{'M', 'A', 'S'}, [3]rune{'S', 'A', 'M'})
	}

	return result
}
func main() {
	scanner, f := myFunctions.GetScannerFromFile("input", bufio.ScanLines)
	defer f.Close()

	twoDimRuneArray := make([][]rune, len("MAS"))

	for i := 0; i < len("MAS")-1; i++ {
		scanner.Scan()
		twoDimRuneArray[i] = []rune(scanner.Text())
	}

	ch := make(chan int)
	go func() {
		defer close(ch)
		for scanner.Scan() {
			twoDimRuneArray[len("MAS")-1] = []rune(scanner.Text())
			ch <- day4Part2(twoDimRuneArray)
			for i := 0; i < len("MAS")-1; i++ {
				twoDimRuneArray[i] = twoDimRuneArray[i+1]
			}
		}
		//bubble the last 3 lines for processing
		for j := 2; j > 0; j-- {
			twoDimRuneArray[j] = []rune("")
			ch <- day4Part2(twoDimRuneArray)
			for i := 0; i < j; i++ {
				twoDimRuneArray[i] = twoDimRuneArray[i+1]
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}()

	sum := 0
	for cnt := range ch {
		sum += cnt
	}

	fmt.Println("sum: ", sum)

}
