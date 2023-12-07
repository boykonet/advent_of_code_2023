package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
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

const (
	_ = iota
	_
	CardTwo
	CardThree
	CardFour
	CardFive
	CardSix
	CardSeven
	CardEight
	CardNine
	CardT
	CardJ
	CardQ
	CardK
	CardA
)

var cardsKeys = map[string]int{
	// from high to low
	"A": CardA,
	"K": CardK,
	"Q": CardQ,
	"J": CardJ,
	"T": CardT,
	"9": CardNine,
	"8": CardEight,
	"7": CardSeven,
	"6": CardSix,
	"5": CardFive,
	"4": CardFour,
	"3": CardThree,
	"2": CardTwo,
}

const (
	FiveOfAKing  = iota // AAAAA
	FourOfAKind         // AA8AA
	FullHouse           // 23332
	ThreeOfAKind        // TTT98
	TwoPair             // 23432
	OnePair             // A23A4
	HighCard            // 23456
)

func compareSets(sets []string) {

}

type camelCard struct {
	Type int
	Bid  int
	Rank int
}

func returnType(card string) int {
	m := make(map[string]int)
	keys := make([]string, 0)
	counts := make([]int, 0)

	// add card into the map
	for _, cardValue := range card {
		_, ok := m[string(cardValue)]
		if ok == false {
			keys = append(keys, string(cardValue))
		}
		m[string(cardValue)] += 1
	}

	for _, key := range keys {
		counts = append(counts, m[key])
	}

	if len(m) == 1 {
		return FiveOfAKing
	} else if len(m) == 2 {
		slices.Sort(counts)
		if counts[0] == 1 && counts[1] == 4 {
			return FourOfAKind
		} else if counts[0] == 2 && counts[1] == 3 {
			return FullHouse
		}
	} else if len(m) == 3 {
		slices.Sort(counts)
		if counts[0] == 1 && counts[1] == 1 && counts[2] == 3 {
			return ThreeOfAKind
		} else if counts[0] == 1 && counts[1] == 2 && counts[2] == 2 {
			return TwoPair
		}
	} else if len(m) == 4 {
		return OnePair
	} else if len(m) == 5 {
		return HighCard
	}
	return -1
}

func parseCards(array []string) (map[string]camelCard, error) {
	cards := make(map[string]camelCard)
	for _, value := range array {
		ss := strings.Split(value, " ")
		bid, err := strconv.Atoi(ss[1])
		if err != nil {
			return nil, err
		}
		cards[ss[0]] = camelCard{Bid: bid}
	}
	return cards, nil
}

func compareTwo(first string, second string) []string {
	lenOneCard := 5
	res := make([]string, 2)
	for i := 0; i < lenOneCard; i++ {
		if first[i] == second[i] {
			continue
		} else if cardsKeys[string(first[i])] > cardsKeys[string(second[i])] {
			res = []string{first, second}
			break
		} else {
			res = []string{second, first}
			break
		}
	}
	return res
}

//func compareCards(cards []string) []int {
//	rankes := make([]int, len(cards))
//	lenOneCard := len(cards[0])
//	m := make(map[string]int)
//
//	maxCard := CardA
//	for i := 0; i < lenOneCard; i++ {
//		if len(rankes) == len(cards) {
//			break
//		}
//		for _, card := range cards {
//			m[card] += maxCard - cardsKeys[string(card[i])]
//		}
//		numbers := make([]int, 0)
//		for _, value := range m {
//			numbers = append(numbers, value)
//		}
//		slices.Sort(numbers)
//		checkNumber := -1
//		fflag := true
//		for _, val := range numbers {
//			if checkNumber == -1 {
//				checkNumber = val
//			} else {
//				if checkNumber < val {
//					checkNumber = val
//				} else {
//					fflag = false
//					break
//				}
//			}
//		}
//		if fflag == true {
//			break
//		}
//	}
//
//	for i := 0; i < len(cards); i++ {
//		m[cards[i]]
//	}
//}

func main() {
	reader := bufio.NewReader(os.Stdin)

	array, err := readFrom(reader)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}

	cards, err := parseCards(array)

	types := make(map[int][]string)
	typesKeys := make([]int, 0)

	for key, value := range cards {
		value.Type = returnType(key)
		cards[key] = value
		_, ok := types[value.Type]
		if ok == false {
			typesKeys = append(typesKeys, value.Type)
		}
		types[value.Type] = append(types[value.Type], key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(typesKeys)))

	//fmt.Println(cards)
	fmt.Println(types)
	//fmt.Println(typesKeys)

	counter := 1
	for _, value := range typesKeys {
		keys := types[value]
		if len(keys) == 1 {
			v := cards[keys[0]]
			v.Rank = counter
			cards[keys[0]] = v
			counter++
		} else {
			types[value] = compareTwo(keys[0], keys[1])
			key := types[value][1]
			v := cards[key]
			v.Rank = counter
			cards[key] = v
			counter++
			key = types[value][0]
			v = cards[key]
			v.Rank = counter
			cards[key] = v
			counter++
		}
	}
	//fmt.Println(cards)

	res := 0

	for _, value := range cards {
		res += value.Bid * value.Rank
	}

	fmt.Println(res)
}
