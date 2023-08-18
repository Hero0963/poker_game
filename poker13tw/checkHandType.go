package poker13tw

import (
	"sort"
)



func IsLegalCards(cards []int) bool {
	if len(cards) != 5 {
		return false
	}

	// 待補

	return true
}

func IsFiveOfAKind(cards []int) bool {
	if len(cards) != 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	for _, count := range recordRkValue {
		if count == 5 {
			return true
		}
	}
	return false
}

func IsStraightFlush(cards []int) bool {
	if len(cards) != 5 {
		return false
	}

	recordSuit := make(map[string]int)
	recordRkValue := make([]int, 0, 5)
	for _, card := range cards {
		suit, rank := GetCardSuit(card), GetCardRank(card)
		recordSuit[suit]++
		rkValue := ConvertRankToPoint(rank)
		recordRkValue = append(recordRkValue, rkValue)
	}

	if len(recordSuit) > 1 {
		return false
	}

	sort.Ints(recordRkValue)
	cnt := 1
	preVal := -1
	for _, x := range recordRkValue {
		if preVal == -1 {
			preVal = x
			continue
		}

		if x != preVal+1 {
			break
		} else {
			cnt++
			preVal = x
		}
	}
	if cnt == 5 {
		return true
	}

	if cnt == 4 && recordRkValue[0] == 2 && recordRkValue[4] == 14 {
		return true
	}

	return false
}

func IsFiveOfGoldTiger(cards []int, gold int) bool {
	if len(cards) != 5 {
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

func IsFourOfAKind(cards []int) bool {
	if len(cards) != 5 {
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

func IsGoldFullHouse(cards []int, gold int) bool {
	if len(cards) != 5 {
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

func IsFullHouse(cards []int) bool {
	if len(cards) != 5 {
		return false
	}

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	checkThree := false
	checkPair := false
	for _, count := range recordRkValue {
		if count == 3 {
			checkThree = true
		}
		if count == 2 {
			checkPair = true
		}
	}

	return checkThree && checkPair
}

func IsPairWithFlush(cards []int) bool {
	if len(cards) != 5 {
		return false
	}

	recordSuit := make(map[string]int)
	for _, card := range cards {
		suit := GetCardSuit(card)
		recordSuit[suit]++
	}

	if len(recordSuit) != 1 {
		return false
	}

	// 使用 map 來檢查是否有重複元素
	seen := make(map[int]bool)
	for _, card := range cards {
		if seen[card] {
			return true // 有重複元素，返回 true
		}
		seen[card] = true
	}

	return true
}

func IsFlush(cards []int) bool {
	if len(cards) != 5 {
		return false
	}

	recordSuit := make(map[string]int)
	for _, card := range cards {
		suit := GetCardSuit(card)
		recordSuit[suit]++
	}

	if len(recordSuit) == 1 {
		return true
	}
	return false
}

func IsStraight(cards []int) bool {
	if len(cards) != 5 {
		return false
	}

	recordRkValue := make([]int, 0, 5)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue = append(recordRkValue, rkValue)
	}

	sort.Ints(recordRkValue)
	cnt := 1
	preVal := -1
	for _, x := range recordRkValue {
		if preVal == -1 {
			preVal = x
			continue
		}

		if x != preVal+1 {
			break
		} else {
			cnt++
			preVal = x
		}
	}
	if cnt == 5 {
		return true
	}

	if cnt == 4 && recordRkValue[0] == 2 && recordRkValue[4] == 14 {
		return true
	}
	return false
}

func IsGoldThreeOfAKind(cards []int, gold int) bool {
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

func IsThreeOfAKind(cards []int) bool {
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

func IsTwoPairs(cards []int) bool {
	if len(cards) != 5 {
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

	if pairCount >= 2 {
		return true
	}
	return false
}

func IsOnePair(cards []int) bool {
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

	if pairCount >= 1 {
		return true
	}
	return false
}

func IsHighCard(cards []int) bool {
	if len(cards) < 1 {
		return false
	}
	return true
}
