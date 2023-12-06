package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type fromSrcToDstConverter struct {
	DestRangeStart   int64
	SourceRangeStart int64
	RangeLength      int64
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

func parseSeedsNumbers(array []string) ([]int64, error) {
	numbers := make([]int64, 0)
	for _, n := range strings.Split(strings.Split(array[0], ": ")[1], " ") {
		number, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
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
				d, err := strconv.ParseInt(numbers[0], 10, 64)
				if err != nil {
					return nil, err
				}
				s, err := strconv.ParseInt(numbers[1], 10, 64)
				if err != nil {
					return nil, err
				}
				r, err := strconv.ParseInt(numbers[2], 10, 64)
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

func findMinimumFromRange(inputData map[string][]fromSrcToDstConverter, seeds []int64) int64 {
	minimum := int64(-1)
	seedsNumbers := make(map[int64][]int64, 0)
	for _, mapKey := range mapKeys {
		srcDstRanges := inputData[mapKey]
		for _, seed := range seeds {
			_, ok := seedsNumbers[seed]
			if !ok {
				seedsNumbers[seed] = make([]int64, 0)
			}
			seedsKey := seed
			seedsValues := seedsNumbers[seedsKey]
			var sourceNumber int64
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

			if mapKey == "humidity-to-location" {
				location := seedsValues[len(seedsValues)-1]
				if minimum == -1 {
					minimum = location
				} else if minimum > location {
					minimum = location
				}
			}
		}
	}
	return minimum
}

func getRangeOfNumbers(start, r int64) []int64 {
	numbers := make([]int64, 0)
	for j := int64(start); j < start+r; j++ {
		numbers = append(numbers, j)
	}
	return numbers
}

func divideToLittlePieces(array []int64) []int64 {
	arr := make([]int64, 0)
	for i := 0; i < len(array); i += 2 {
		start := array[i]
		r := array[i+1]
		newRange := r / 4096
		rem := r - newRange*4096
		for j := 0; j < 4096; j += 1 {
			arr = append(arr, start, newRange)
			start += newRange
		}
		arr = append(arr, start, rem)
	}
	return arr
}

func dataProcess(reader *bufio.Reader) (int64, error) {
	array, err := readFrom(reader)
	if err != nil {
		return -1, err
	}

	seedsNumbersAndRanges, err := parseSeedsNumbers(array)
	if err != nil {
		return -1, err
	}
	fmt.Println("parse seeds numbers DONE")

	//keys := removeNotUniqueSubsequences(seedsNumbersAndRanges)

	//fmt.Println("keys DONE, len =", len(keys))

	inputData, err := parseSrcDstNumbers(array[2:])
	if err != nil {
		return -1, err
	}

	fmt.Println("parse src dst numbers DONE")

	seedsNumbersAndRanges = divideToLittlePieces(seedsNumbersAndRanges)

	fmt.Println("divide to little pieces DONE")
	//fmt.Println(seedsNumbersAndRanges)

	var wg sync.WaitGroup
	mu := sync.Mutex{}

	min := make([]int64, 0)

	for i := 0; i < len(seedsNumbersAndRanges); i += 2 {
		wg.Add(1)
		start := seedsNumbersAndRanges[i]
		r := seedsNumbersAndRanges[i+1]
		go func() {
			res := getRangeOfNumbers(start, r)
			//mu.Lock()
			//seeds := make(map[int64][]int64, 0)
			//for j := 0; j < len(res); j++ {
			//	seeds[res[j]] = make([]int64, 0)
			//}
			mu.Lock()
			min = append(min, findMinimumFromRange(inputData, res))
			mu.Unlock()
			//rangeSeedsNumbers = append(rangeSeedsNumbers, res)
			//mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	//fmt.Println(min)

	minimum := int64(-1)
	for _, m := range min {
		if minimum == -1 {
			minimum = m
		} else if m < minimum {
			minimum = m
		}
	}
	return minimum, nil
}

//func findMinimumForRange(seedsNumbers map[int64][]int64, ch chan int64) {
//	var minimum int64 = -1
//	for _, values := range seedsNumbers {
//		if minimum == -1 {
//			minimum = values[len(values)-1]
//		} else if minimum > values[len(values)-1] {
//			minimum = values[len(values)-1]
//		}
//	}
//	ch <- minimum
//}

func main() {
	reader := bufio.NewReader(os.Stdin)

	minimum, err := dataProcess(reader)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}

	fmt.Println(minimum)

	//var minimum int64 = -1
	//for _, values := range seedsNumbers {
	//	if minimum == -1 {
	//		minimum = values[len(values)-1]
	//	} else if minimum > values[len(values)-1] {
	//		minimum = values[len(values)-1]
	//	}
	//}
	//fmt.Println(minimum)
}
