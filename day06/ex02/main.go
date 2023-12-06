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

func getNumber(parts []string) (int64, error) {
	number := ""
	for _, val := range parts {
		if val != "" {
			number += val
		}
	}

	res, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return -1, err
	}
	return res, nil
}

func parse(s []string) (int64, int64, error) {
	timeRanges := strings.Split(strings.Trim(strings.Split(s[0], ":")[1], " "), " ")
	distanceRanges := strings.Split(strings.Trim(strings.Split(s[1], ":")[1], " "), " ")

	raceTime, err := getNumber(timeRanges)
	if err != nil {
		return -1, -1, err
	}

	raceDistance, err := getNumber(distanceRanges)
	if err != nil {
		return -1, -1, err
	}

	return raceTime, raceDistance, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	array, err := readFrom(reader)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}

	raceTime, raceDistance, err := parse(array)
	m := make(map[int64]int64, 0)
	for i := int64(0); i <= raceTime; i++ {
		m[i*(raceTime-i)] += 1
	}
	curr := int64(0)
	for key, value := range m {
		if key > raceDistance {
			curr += value
		}
	}
	fmt.Println(curr)
}
