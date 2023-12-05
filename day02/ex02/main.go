package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func findGameNumber(s string) (int, error) {
	return strconv.Atoi(strings.Split(s, " ")[1])
}

func maxNecessaryCubesPerGame(sets []string) (int, int, int, error) {
	max := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		for _, cube := range cubes {
			numberAndColor := strings.Split(strings.Trim(cube, " "), " ")
			n, err := strconv.Atoi(numberAndColor[0])
			if err != nil {
				return -1, -1, -1, err
			}
			color := strings.Trim(numberAndColor[1], " ")
			if max[color] < n {
				max[color] = n
			}
		}
	}
	return max["red"], max["green"], max["blue"], nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	power := 0

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
		oneGameSets := strings.Split(strings.Trim(sline[1], " "), "; ")
		red, green, blue, err := maxNecessaryCubesPerGame(oneGameSets)
		if err != nil {
			fmt.Println("Unexpected error:", err)
			return
		}
		power += red * green * blue
	}
	fmt.Println(power)
}
