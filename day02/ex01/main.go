package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	redCubes   int = 12
	greenCubes int = 13
	blueCubes  int = 14
)

func findGameNumber(s string) (int, error) {
	return strconv.Atoi(strings.Split(s, " ")[1])
}

func checkCubesPerGame(sets []string) (bool, error) {
	for _, set := range sets {
		m := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		cubes := strings.Split(set, ", ")
		for _, cube := range cubes {
			numberAndColor := strings.Split(strings.Trim(cube, " "), " ")
			n, err := strconv.Atoi(numberAndColor[0])
			if err != nil {
				return false, err
			}
			color := strings.Trim(numberAndColor[1], " ")
			_, ok := m[color]
			if ok != true {
				return false, fmt.Errorf("Unexpected color: %v\n", numberAndColor[1])
			}
			m[color] += n
		}
		if !(m["red"] <= redCubes && m["blue"] <= blueCubes && m["green"] <= greenCubes) {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := 0

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Unexpected error:", err)
			return
		}
		sline := strings.Split(string(line), ": ")
		gameNumber, err := findGameNumber(strings.Trim(sline[0], " "))
		if err != nil {
			fmt.Println("Unexpected error:", err)
			return
		}
		oneGameSets := strings.Split(strings.Trim(sline[1], " "), "; ")
		isOK, err := checkCubesPerGame(oneGameSets)
		if err != nil {
			fmt.Println("Unexpected error:", err)
			return
		}
		if isOK == true {
			res += gameNumber
		}
	}
	fmt.Println(res)
}
