package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

//const (
//	_ = iota
//	_
//	CardTwo
//	CardThree
//	CardFour
//	CardFive
//	CardSix
//	CardSeven
//	CardEight
//	CardNine
//	CardT
//	CardJ
//	CardQ
//	CardK
//	CardA
//)

var cardsKeys = map[rune]string{
	// from high to low
	'A': "L",
	'K': "E",
	'Q': "D",
	'J': "C",
	'T': "B",
	//"9": '9',
	//"8": '8',
	//"7": '7',
	//"6": '6',
	//"5": '5',
	//"4": '4',
	//"3": '3',
	//"2": '2',
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
		sort.Sort(sort.IntSlice(counts))
		if counts[0] == 1 && counts[1] == 4 {
			return FourOfAKind
		} else if counts[0] == 2 && counts[1] == 3 {
			return FullHouse
		}
	} else if len(m) == 3 {
		sort.Sort(sort.IntSlice(counts))
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

func parseCards(array []string) ([]string, []camelCard, error) {
	cards := make([]string, 0)
	bids := make([]camelCard, 0)
	for _, value := range array {
		ss := strings.Split(value, " ")
		cards = append(cards, ss[0])
		bid, err := strconv.Atoi(ss[1])
		if err != nil {
			return nil, nil, err
		}
		bids = append(bids, camelCard{Bid: bid})
	}
	return cards, bids, nil
}

//func compareCards(combinations []string) []string {
//	lenOneCard := 5
//	m := make(map[int][]string)
//	rm := make(map[string]int)
//
//	m[0] = combinations
//	maxCard := CardA
//	for i := 0; i < lenOneCard; i++ {
//		oldKeys := make([]int, 0)
//		if len(m) == len(combinations) {
//			break
//		}
//		for key, values := range m {
//			if len(values) > 1 {
//				start := key
//				for _, value := range values {
//					newKey := start + (maxCard - cardsKeys[string(value[i])])
//					m[newKey] = append(m[newKey], value)
//				}
//				oldKeys = append(oldKeys, key)
//			}
//		}
//		for _, key := range oldKeys {
//			delete(m, key)
//		}
//	}
//
//	keys := make([]int, 0)
//	for key, value := range m {
//		rm[value[0]] = key
//		keys = append(keys, key)
//	}
//
//	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
//
//	res := make([]string, 0)
//
//	for _, key := range keys {
//		res = append(res, m[key][0])
//	}
//	return res
//}

func compareCards(cards []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(cards)))
	return cards
}

func changeCards(cards []string) []string {
	for i := 0; i < len(cards); i++ {
		for j := 0; j < len(cards[i]); j++ {
			card := cards[i]
			_, ok := cardsKeys[rune(card[j])]
			if ok == true {
				strings.Replace(cards[i], string(cards[i][j]), cardsKeys[rune(card[j])], j)
			}
		}
	}
	return cards
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	array, err := readFrom(reader)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}

	cards, cardsInfo, err := parseCards(array)
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		return
	}

	cards = changeCards(cards)

	m := make(map[string]camelCard)
	for i := 0; i < len(cards); i++ {
		m[cards[i]] = cardsInfo[i]
	}

	types := make(map[int][]string)
	typesKeys := make([]int, 0)

	for key, value := range m {
		value.Type = returnType(key)
		m[key] = value
		_, ok := types[value.Type]
		if ok == false {
			typesKeys = append(typesKeys, value.Type)
		}
		types[value.Type] = append(types[value.Type], key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(typesKeys)))

	fmt.Println(types)
	fmt.Println(typesKeys)

	sortedCards := make([]string, 0)

	for _, value := range typesKeys {
		keys := types[value]
		if len(keys) == 1 {
			sortedCards = append(sortedCards, keys...)
		} else if len(keys) > 1 {
			sortedCards = append(sortedCards, compareCards(keys)...)
		}
	}
	fmt.Println(sortedCards)

	for i := 0; i < len(sortedCards); i++ {
		key := sortedCards[i]
		cardInfo := m[key]
		cardInfo.Rank = i + 1
		m[key] = cardInfo
	}

	res := 0

	for _, value := range m {
		res += value.Bid * value.Rank
	}

	fmt.Println(res)
}
