package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type fromSrcToDstConverter struct {
	Dest   int
	Source int
	Range  int
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
	//m := make(map[string][]fromSrcToDstConverter)

	array, err := readFrom(reader)
	if err != nil {
		fmt.Println("Unexpected error:", err)
		return
	}
	fmt.Println(array)
}
