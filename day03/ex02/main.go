package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func findAndParseNumber(str string) ([]int, [][]int, error) {
	numbers := make([]int, 0)
	coords := make([][]int, 0)

	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			continue
		}
		var strNum []byte
		fIndex := i
		for str[i] >= '0' && str[i] <= '9' {
			strNum = append(strNum, str[i])
			i++
		}
		lIndex := i - 1
		number, err := strconv.Atoi(string(strNum))
		if err != nil {
			return nil, nil, err
		}
		numbers = append(numbers, number)
		coords = append(coords, [][]int{{fIndex, lIndex}}...)
	}
	return numbers, coords, nil
}

func checkAllGears(m map[string][]int, array []string, line string, yCoord int) error {
	res := make([]int, 0)
	numbers, coords, err := findAndParseNumber(line)
	if err != nil {
		return err
	}
	for i := 0; i < len(numbers); i++ {
		xBeginx := coords[i][0]
		xEndx := coords[i][1]
		beginCoords := [][]int{
			{xBeginx - 1, yCoord - 1},
			{xBeginx - 1, yCoord},
			{xBeginx - 1, yCoord + 1},
		}
		endCoords := [][]int{
			{xEndx + 1, yCoord - 1},
			{xEndx + 1, yCoord},
			{xEndx + 1, yCoord + 1},
		}
		flag := false
		for j := 0; j < 3; j++ {
			xBegin := beginCoords[j][0]
			yBegin := beginCoords[j][1]
			xEnd := endCoords[j][0]
			for ; xBegin <= xEnd; xBegin++ {
				ch := array[yBegin][xBegin]
				if !(ch == '.' || (ch >= '0' && ch <= '9')) {
					if ch == '*' {
						key := strconv.Itoa(xBegin) + "-" + strconv.Itoa(yBegin)
						m[key] = append(m[key], numbers[i])
					}
					flag = true
				}
			}
		}
		if flag == true {
			res = append(res, numbers[i])
		}
	}
	return nil
}

func sum(array []string) (int, error) {
	res := 0
	m := make(map[string][]int)
	for index, line := range array {
		err := checkAllGears(m, array, line, index)
		if err != nil {
			return 0, err
		}
	}

	for _, values := range m {
		if len(values) == 2 {
			res += values[0] * values[1]
		}
	}

	return res, nil
}

func fillDotsAroundPerimeter(array []string) {
	var dots string
	for i := 0; i < len(array[0]); i++ {
		dots += "."
	}

	array = append([]string{dots}, array...)
	array = append(array, dots)

	for i := 0; i < len(array); i++ {
		array[i] = "." + array[i] + "."
	}
}

func readFrom(reader *bufio.Reader) ([]string, error) {
	array := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		array = append(array, string(line))
	}
	return array, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	array, err := readFrom(reader)
	if err != nil {
		fmt.Println("Unexpected error:", err)
		return
	}

	fillDotsAroundPerimeter(array)

	res, err := sum(array)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}
	fmt.Println(res)
}
