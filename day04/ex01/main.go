package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

func convertToNumbers(str []string) ([]int, error) {
	numbers := make([]int, 0)
	for _, s := range str {
		if s == "" {
			continue
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	pointsCount := 0

	array, err := readFrom(reader)
	if err != nil {
		fmt.Println("Unexpected error:", err)
		return
	}

	for _, arr := range array {
		ss := strings.Split(arr, "|")
		winNumbers, err := convertToNumbers(strings.Split(strings.Split(strings.Trim(ss[0], " "), ":")[1], " "))
		if err != nil {
			fmt.Println("Unexpected error:", err)
			return
		}
		numbers, err := convertToNumbers(strings.Split(strings.Trim(ss[1], " "), " "))
		m := make(map[int]bool, 0)
		for _, number := range winNumbers {
			m[number] = false
		}
		for _, number := range numbers {
			_, ok := m[number]
			if ok == true {
				m[number] = true
			}
		}
		res := 0
		counter := 0
		for _, value := range m {
			if value == true {
				if counter == 0 {
					res += 1
				} else {
					res *= 2
				}
				counter++
			}
		}
		pointsCount += res
	}
	fmt.Println(pointsCount)
}
