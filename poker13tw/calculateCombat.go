package poker13tw

import (
	"sort"
)

func UpdateOffset(oriOffset int, x int) int {
	offsetBase := 100
	return oriOffset*offsetBase + x
}

func GetScoreFiveOfAKind(cards []int) int {
	baseScore := CombatBase * int(FiveKind)
	offset := 0
	card := cards[0]
	rank := GetCardRank(card)
	rkValue := ConvertRankToPoint(rank)

	offset = UpdateOffset(offset, rkValue)
	return baseScore + offset
}

func GetScoreStraightFlush(cards []int) int {
	baseScore := CombatBase * int(StraightFlush)
	offset := 0
	recordRkValues := make([]int, 0, 5)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValues = append(recordRkValues, rkValue)
	}

	sort.Ints(recordRkValues)
	maxRkValue := recordRkValues[len(recordRkValues)-1]
	offset = UpdateOffset(offset, maxRkValue)
	return baseScore + offset
}

func GetScoreFiveGoldTiger(cards []int, gold int) int {
	baseScore := CombatBase * int(FiveOfGoldTiger)
	offset := 0
	offset = UpdateOffset(offset, gold)
	singleValue := 0
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		if rkValue != gold {
			singleValue = rkValue
			break
		}
	}

	offset = UpdateOffset(offset, singleValue)
	return baseScore + offset
}

func GetScoreFourOfAKind(cards []int) int {
	baseScore := CombatBase * int(FourOfAKind)
	offset := 0
	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}
	kindValue, singleValue := 0, 0
	for rkValue, count := range recordRkValue {
		if count >= 4 {
			kindValue = rkValue
		} else {
			singleValue = rkValue
		}
	}

	if singleValue == 0 {
		singleValue = kindValue
	}
	offset = UpdateOffset(offset, kindValue)
	offset = UpdateOffset(offset, singleValue)
	return baseScore + offset
}

func GetScoreGoldFullHouse(cards []int, gold int) int {
	baseScore := CombatBase * int(GoldFullHouse)
	offset := 0
	offset = UpdateOffset(offset, gold)

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	for rkValue, count := range recordRkValue {
		if rkValue != gold && count >= 2 {
			offset = UpdateOffset(offset, rkValue)
		}
	}

	return baseScore + offset
}

func GetScoreFullHouse(cards []int) int {
	baseScore := CombatBase * int(FullHouse)
	offset := 0

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	updateOrder := [2]int{}
	for rkValue, count := range recordRkValue {
		if count == 3 {
			updateOrder[0] = rkValue
		} else {
			updateOrder[1] = rkValue
		}
	}

	for _, v := range updateOrder {
		offset = UpdateOffset(offset, v)
	}

	return baseScore + offset
}

func GetScorePairWithFlush(cards []int) int {
	baseScore := CombatBase * int(PairWithFlush)
	offset := 0

	recordRkValues := make(map[int]int)
	pairValue := -1
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValues[rkValue]++
		if recordRkValues[rkValue] > 1 {

			if rkValue > pairValue {
				pairValue = rkValue
			}

		}
	}

	offset = UpdateOffset(offset, pairValue)
	values := []int{}
	for rkValue := range recordRkValues {
		if rkValue != pairValue {
			values = append(values, rkValue)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	for _, v := range values {
		offset = UpdateOffset(offset, v)
	}

	return baseScore + offset
}

func GetScoreFlush(cards []int) int {
	baseScore := CombatBase * int(Flush)
	offset := 0
	rkValues := []int{}
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		rkValues = append(rkValues, rkValue)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(rkValues)))
	for _, x := range rkValues {
		offset = UpdateOffset(offset, x)
	}

	return baseScore + offset
}

func GetScoreStraight(cards []int) int {
	baseScore := CombatBase * int(Straight)
	offset := 0
	rkValues := []int{}
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		rkValues = append(rkValues, rkValue)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(rkValues)))
	offset = UpdateOffset(offset, rkValues[0])

	return baseScore + offset
}

func GetScoreGoldThreeOfAKind(cards []int, gold int) int {
	baseScore := CombatBase * int(GoldThreeOfAKind)
	offset := 0
	offset = UpdateOffset(offset, gold)

	rkValues := []int{}
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		if rkValue != gold {
			rkValues = append(rkValues, rkValue)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(rkValues)))
	for _, x := range rkValues {
		offset = UpdateOffset(offset, x)
	}
	return baseScore + offset
}

func GetScoreThreeOfAKind(cards []int) int {
	baseScore := CombatBase * int(ThreeOfAKind)
	offset := 0

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	kindValue := 0
	rkValues := []int{}
	for rkValue, count := range recordRkValue {
		if count >= 3 {
			kindValue = rkValue
		} else {
			rkValues = append(rkValues, rkValue)
		}
	}

	for len(rkValues) < 2 {
		rkValues = append(rkValues, kindValue)
	}

	offset = UpdateOffset(offset, kindValue)
	sort.Sort(sort.Reverse(sort.IntSlice(rkValues)))
	for _, x := range rkValues {
		offset = UpdateOffset(offset, x)
	}
	return baseScore + offset
}

func GetScoreTwoPairs(cards []int) int {
	baseScore := CombatBase * int(TwoPairs)
	offset := 0

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	pairValues := []int{}
	singleValue := 0
	for rkValue, count := range recordRkValue {
		if count >= 2 {
			pairValues = append(pairValues, rkValue)
		} else {
			singleValue = rkValue
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(pairValues)))
	for _, x := range pairValues {
		offset = UpdateOffset(offset, x)
	}

	offset = UpdateOffset(offset, singleValue)
	return baseScore + offset
}

func GetScoreOnePair(cards []int) int {
	baseScore := CombatBase * int(OnePair)
	offset := 0

	recordRkValue := make(map[int]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRkValue[rkValue]++
	}

	pairValue := 0
	rkValues := []int{}
	for rkValue, count := range recordRkValue {
		if count >= 2 {
			pairValue = rkValue
		} else {
			rkValues = append(rkValues, rkValue)
		}
	}

	offset = UpdateOffset(offset, pairValue)
	sort.Sort(sort.Reverse(sort.IntSlice(rkValues)))
	for _, x := range rkValues {
		offset = UpdateOffset(offset, x)
	}
	return baseScore + offset
}

func GetScoreHighCard(cards []int) int {
	baseScore := CombatBase * int(HighCard)
	offset := 0

	rkValues := []int{}
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		rkValues = append(rkValues, rkValue)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(rkValues)))
	for _, x := range rkValues {
		offset = UpdateOffset(offset, x)
	}

	return baseScore + offset
}

func getPokerHandScore(cards []int, gold int) int {
	if IsFiveOfAKind(cards) {
		score := GetScoreFiveOfAKind(cards)
		return score
	}

	if IsStraightFlush(cards) {
		score := GetScoreStraightFlush(cards)
		return score
	}

	if IsFiveOfGoldTiger(cards, gold) {
		score := GetScoreFiveGoldTiger(cards, gold)
		return score
	}

	if IsFourOfAKind(cards) {
		score := GetScoreFourOfAKind(cards)
		return score
	}

	if IsGoldFullHouse(cards, gold) {
		score := GetScoreGoldFullHouse(cards, gold)
		return score
	}

	if IsFullHouse(cards) {
		score := GetScoreFullHouse(cards)
		return score
	}

	if IsPairWithFlush(cards) {
		score := GetScorePairWithFlush(cards)
		return score
	}

	if IsFlush(cards) {
		score := GetScoreFlush(cards)
		return score
	}

	if IsStraight(cards) {
		score := GetScoreStraight(cards)
		return score
	}

	if IsGoldThreeOfAKind(cards, gold) {
		score := GetScoreGoldThreeOfAKind(cards, gold)
		return score
	}

	if IsThreeOfAKind(cards) {
		score := GetScoreThreeOfAKind(cards)
		return score
	}

	if IsTwoPairs(cards) {
		score := GetScoreTwoPairs(cards)
		return score
	}

	score := GetScoreHighCard(cards)
	return score
}
