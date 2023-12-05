package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var m = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var lenStrings = []int{3, 4, 5}

func firstAndLastNumberOccurrence(s string) (int, int) {
	first := -1
	last := -1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if first == -1 {
				first, _ = strconv.Atoi(string(s[i]))
				last = first
			} else {
				last, _ = strconv.Atoi(string(s[i]))
			}
		}
		for _, l := range lenStrings {
			if len(s[i:]) < l {
				continue
			}
			val, ok := m[s[i:i+l]]
			if ok == true {
				if first == -1 {
					first = val
					last = first
				} else {
					last = val
				}
			}
		}
	}
	return first, last
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
		f, l := firstAndLastNumberOccurrence(string(line))
		res += f * 10
		res += l
	}

	fmt.Println(res)
}
