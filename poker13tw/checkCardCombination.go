package poker13tw

import (
	"sort"
)

func HasFiveOfAKind(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	for _, count := range recordRkValue {
		if count >= 5 {
			return true
		}
	}

	return false
}

func HasStraightFlush(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	checkStraight := func(poker []int) bool {
		points := make([]int, len(poker))
		copy(points, poker)
		sort.Ints(points)

		cnt := 1
		preVal := -1
		for _, p := range points {
			if preVal == -1 {
				preVal = p
				continue
			}

			if p != preVal+1 {
				cnt = 1
				preVal = p
			} else {
				cnt++
				preVal = p
			}

			if cnt >= 5 {
				return true
			}
		}

		specialStraight := []int{2, 3, 4, 5, 14}
		sp := make(map[int]bool)
		for _, val := range points {
			sp[val] = true
		}

		for _, ss := range specialStraight {
			if !sp[ss] {
				return false
			}
		}

		return true
	}

	splitBySuit := make(map[string][]int)
	for _, card := range cards {
		suit, rank := GetCardSuit(card), GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		splitBySuit[suit] = append(splitBySuit[suit], rkValue)
	}

	for _, values := range splitBySuit {
		if checkStraight(values) {
			return true
		}
	}

	return false
}

func HasFiveOfGoldTiger(cards []int, gold int) bool {
	if len(cards) < 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	if recordRkValue[gold] >= 4 {
		return true
	}

	return false
}

func HasFourOfAKind(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	for _, count := range recordRkValue {
		if count >= 4 {
			return true
		}
	}

	return false
}

func HasGoldFullHouse(cards []int, gold int) bool {
	if len(cards) < 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	checkGold := false
	checkPair := false
	for rkValue, count := range recordRkValue {
		if rkValue == gold {
			if count == 3 {
				checkGold = true
			}
		} else {
			if count >= 2 {
				checkPair = true
			}
		}
	}

	return checkGold && checkPair
}

func HasFullHouse(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	countKind := 0
	countPair := 0
	for _, count := range recordRkValue {
		if count >= 3 {
			countKind++
		} else {
			if count == 2 {
				countPair++
			}
		}
	}

	return countKind >= 1 && (countKind+countPair >= 2)
}

func HasPairWithFlush(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	checkFlush := func(poker []int) bool {
		recordSuit := make(map[string]int)
		for _, p := range poker {
			suit := GetCardSuit(p)
			recordSuit[suit]++
		}

		for _, count := range recordSuit {
			if count >= 3 {
				return true
			}
		}

		return false
	}

	recordPairs := make(map[int][]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordPairs[rkValue] = append(recordPairs[rkValue], card)
	}

	for _, pair := range recordPairs {
		if len(pair) < 2 {
			continue
		}

		combinations := Combinations(pair, 2)
		for _, c := range combinations {
			values := make([]int, 0)
			for _, card := range cards {
				if !InSlice(card, c) {
					values = append(values, card)
				}
			}

			if checkFlush(values) {
				return true
			}
		}
	}

	return false
}

func HasFlush(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	recordSuit := make(map[string]int)
	for _, card := range cards {
		suit := GetCardSuit(card)
		recordSuit[suit]++
	}

	for _, count := range recordSuit {
		if count >= 5 {
			return true
		}
	}

	return false
}

func HasStraight(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	s := make(map[int]bool)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		s[rkValue] = true
	}

	if len(s) < 5 {
		return false
	}

	values := make([]int, 0, len(s))
	for val := range s {
		values = append(values, val)
	}
	sort.Ints(values)

	cnt := 1
	preVal := -1
	for _, x := range values {
		if preVal == -1 {
			preVal = x
			continue
		}

		if x != preVal+1 {
			cnt = 1
			preVal = x
		} else {
			cnt++
			preVal = x
		}

		if cnt >= 5 {
			return true
		}
	}

	specialStraight := []int{2, 3, 4, 5, 14}
	for _, ss := range specialStraight {
		if !s[ss] {
			return false
		}
	}

	return true
}

func HasGoldThreeOfAKind(cards []int, gold int) bool {
	if len(cards) < 3 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	if recordRkValue[gold] >= 3 {
		return true
	}

	return false
}

func HasThreeOfAKind(cards []int) bool {
	if len(cards) < 3 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	for _, count := range recordRkValue {
		if count >= 3 {
			return true
		}
	}

	return false
}

func HasTwoPairs(cards []int) bool {
	if len(cards) < 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	pairCount := 0
	for _, count := range recordRkValue {
		if count >= 2 {
			pairCount++
		}
	}

	return pairCount >= 2
}

func HasOnePair(cards []int) bool {
	if len(cards) < 2 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	pairCount := 0
	for _, count := range recordRkValue {
		if count >= 2 {
			pairCount++
		}
	}

	return pairCount >= 1
}

func HasHighCard(cards []int) bool {
	if len(cards) < 1 {
		return false
	}
	return true
}
