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

type timeAndDistance struct {
	Time     int
	Distance int
}

func parse(s []string) ([]timeAndDistance, error) {
	timeDistance := make([]timeAndDistance, 0)
	timeRanges := strings.Split(strings.Trim(strings.Split(s[0], ":")[1], " "), " ")
	distanceRanges := strings.Split(strings.Trim(strings.Split(s[1], ":")[1], " "), " ")

	for _, val := range timeRanges {
		if val != "" {
			number, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			timeDistance = append(timeDistance, timeAndDistance{Time: number})
		}
	}
	counter := 0
	for _, val := range distanceRanges {
		if val != "" {
			number, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			timeDistance[counter].Distance = number
			counter++
		}
	}
	return timeDistance, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	array, err := readFrom(reader)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}

	timeDistance, err := parse(array)
	res := 1
	for _, td := range timeDistance {
		m := make(map[int]int, 0)
		for i := 0; i <= td.Time; i++ {
			m[i*(td.Time-i)] += 1
		}
		curr := 0
		for key, value := range m {
			if key > td.Distance {
				curr += value
			}
		}
		res *= curr
	}
	fmt.Println(res)
}
