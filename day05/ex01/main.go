package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type fromSrcToDstConverter struct {
	DestRangeStart   int
	SourceRangeStart int
	RangeLength      int
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

var mapKeys = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func parseSeedsNumbers(array []string) (map[int][]int, error) {
	numbers := make(map[int][]int, 0)
	for _, n := range strings.Split(strings.Split(array[0], ": ")[1], " ") {
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		numbers[number] = make([]int, 0)
	}
	return numbers, nil
}

func parseSrcDstNumbers(array []string) (map[string][]fromSrcToDstConverter, error) {
	inputData := map[string][]fromSrcToDstConverter{
		"seed-to-soil":            make([]fromSrcToDstConverter, 0),
		"soil-to-fertilizer":      make([]fromSrcToDstConverter, 0),
		"fertilizer-to-water":     make([]fromSrcToDstConverter, 0),
		"water-to-light":          make([]fromSrcToDstConverter, 0),
		"light-to-temperature":    make([]fromSrcToDstConverter, 0),
		"temperature-to-humidity": make([]fromSrcToDstConverter, 0),
		"humidity-to-location":    make([]fromSrcToDstConverter, 0),
	}
	for i := 0; i < len(array); i++ {
		if array[i] == "" {
			continue
		}
		key := strings.Split(array[i], " ")[0]
		value, ok := inputData[key]
		if ok == true {
			i++
			counter := 0
			for i < len(array) && array[i] != "" {
				numbers := strings.Split(array[i], " ")
				d, err := strconv.Atoi(numbers[0])
				if err != nil {
					return nil, err
				}
				s, err := strconv.Atoi(numbers[1])
				if err != nil {
					return nil, err
				}
				r, err := strconv.Atoi(numbers[2])
				if err != nil {
					return nil, err
				}
				value = append(value, fromSrcToDstConverter{DestRangeStart: d, SourceRangeStart: s, RangeLength: r})
				inputData[key] = value
				counter++
				i++
			}
		}
	}
	return inputData, nil
}

func findSrcDstVariations(inputData map[string][]fromSrcToDstConverter, seedsNumbers map[int][]int) {
	for _, mapKey := range mapKeys {
		srcDstRanges := inputData[mapKey]
		for seedsKey, seedsValues := range seedsNumbers {
			var sourceNumber int
			if len(seedsValues) == 0 {
				sourceNumber = seedsKey
			} else {
				sourceNumber = seedsValues[len(seedsValues)-1]
			}
			bestResult := sourceNumber
			for _, srcDstRange := range srcDstRanges {
				needToPlus := sourceNumber - srcDstRange.SourceRangeStart
				if needToPlus >= 0 && needToPlus <= srcDstRange.RangeLength {
					bestResult = srcDstRange.DestRangeStart + needToPlus
					break
				}
			}
			seedsValues = append(seedsValues, bestResult)
			seedsNumbers[seedsKey] = seedsValues
		}
	}
}

func dataProcess(reader *bufio.Reader) (map[int][]int, error) {
	array, err := readFrom(reader)
	if err != nil {
		return nil, err
	}
	seedsNumbers, err := parseSeedsNumbers(array)
	if err != nil {
		return nil, err
	}

	inputData, err := parseSrcDstNumbers(array[2:])
	if err != nil {
		return nil, err
	}

	findSrcDstVariations(inputData, seedsNumbers)
	return seedsNumbers, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	seedsNumbers, err := dataProcess(reader)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}
	fmt.Println(seedsNumbers)

	minimum := -1
	for _, values := range seedsNumbers {
		if minimum == -1 {
			minimum = values[len(values)-1]
		} else if minimum > values[len(values)-1] {
			minimum = values[len(values)-1]
		}
	}
	fmt.Println(minimum)
}
