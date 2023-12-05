package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func firstNumberOccurrence(s []byte) (int, error) {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			return strconv.Atoi(string(ch))
		}
	}
	return -1, fmt.Errorf("No numbers in the string\n")
}

func lastNumberOccurrence(s []byte) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return strconv.Atoi(string(s[i]))
		}
	}
	return -1, fmt.Errorf("No numbers in the string\n")
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
		num, err := firstNumberOccurrence(line)
		if err != nil {
			fmt.Println("Unexpected error:", err)
			return
		}
		res += num * 10
		num, err = lastNumberOccurrence(line)
		if err != nil {
			fmt.Println("Unexpected error:", err)
			return
		}
		res += num
	}

	fmt.Println(res)
}
