package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testCasesString := scanner.Text()
	testCases, _ := strconv.Atoi(testCasesString)
	for testCase := 0; testCase < testCases; testCase++ {
		scanner.Scan()
		sizeStr := scanner.Text()
		size, _ := strconv.Atoi(sizeStr)
		matrix := make([][]int, size)
		for row := 0; row < size; row++ {
			scanner.Scan()
			rowStr := strings.Split(scanner.Text(), " ")
			matrix[row] = make([]int, size)
			for col := 0; col < size; col++ {
				matrix[row][col], _ = strconv.Atoi(rowStr[col])
			}
		}

		trace := getTrace(matrix)
		rowDuplicates := getRowDuplicates(matrix)
		colDuplicates := getColDuplicates(matrix)
		fmt.Printf("Case #%v: %v %v %v\n", testCase+1, trace, rowDuplicates, colDuplicates)
	}
}

func getTrace(matrix [][]int) int {
	sum := 0
	for index := 0; index < len(matrix[0]); index++ {
		sum += matrix[index][index]
	}

	return sum
}

func getRowDuplicates(matrix [][]int) int {
	rDups := 0
	for row := 0; row < len(matrix[0]); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			subject := matrix[row][col]
			if isDuplicate(subject, col, matrix[row]) {
				rDups++
				break
			}
		}
	}

	return rDups
}

func getColDuplicates(matrix [][]int) int {
	nMatrix := make([][]int, len(matrix[0]))
	for row := 0; row < len(matrix[0]); row++ {
		nMatrix[row] = make([]int, len(matrix[0]))
		for col := 0; col < len(matrix[0]); col++ {
			nMatrix[row][col] = matrix[col][row]
		}
	}

	return getRowDuplicates(nMatrix)
}

func isDuplicate(subject int, ownPos int, toCheck []int) bool {
	for pos, val := range toCheck {
		if pos == ownPos {
			continue
		}
		if subject == val {
			return true
		}
	}

	return false
}
