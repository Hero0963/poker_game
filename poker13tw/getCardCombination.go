package poker13tw

import (
	"sort"
)

func GetFiveOfAKindFromPokerHand(cards []int) [][]int {
	if !HasFiveOfAKind(cards) {
		return nil
	}
	result := [][]int{}
	recordRKValue := make(map[int]int)
	recordSplitCards := make(map[int][]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRKValue[rkValue]++
		recordSplitCards[rkValue] = append(recordSplitCards[rkValue], card)
	}

	for rkValue, count := range recordRKValue {
		if count >= 5 {
			result = append(result, recordSplitCards[rkValue])
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)

	return sortedResult
}

func GetStraightFlushFromPokerHand(cards []int) [][]int {
	if !HasStraightFlush(cards) {
		return nil
	}

	result := [][]int{}

	getStraightFlush := func(poker []int) [][]int {
		res := [][]int{}
		if len(poker) < 5 {
			return res
		}

		suitRecover := poker[0] / 13
		points := []int{}
		for _, p := range poker {
			rank := GetCardRank(p)
			rkValue := ConvertRankToPoint(rank)
			points = append(points, rkValue)
		}

		sort.Ints(points)

		record := []int{}
		preVal := -1
		for _, p := range points {
			if preVal == -1 {
				record = []int{p}
				preVal = p
				continue
			}

			if p != preVal+1 {
				record = []int{p}
			} else {
				record = append(record, p)
				preVal = p
			}

			if len(record) == 5 {
				res = append(res, append([]int{}, record...))
				record = record[1:]
			}
		}

		specialStraight := []int{2, 3, 4, 5, 14}
		checkSpecialStraight := true
		sp := make(map[int]bool)
		for _, ss := range specialStraight {
			sp[ss] = true
		}
		for _, p := range points {
			if _, exists := sp[p]; !exists {
				checkSpecialStraight = false
				break
			}
		}

		if checkSpecialStraight {
			res = append(res, append([]int{}, specialStraight...))
		}

		ans := [][]int{}
		for _, rkValues := range res {
			straightFlush := []int{}
			for _, v := range rkValues {
				r := recoverOriginalPointByRKValueAndSuit(v, suitRecover)
				straightFlush = append(straightFlush, r)
			}
			ans = append(ans, straightFlush)
		}

		return ans
	}

	splitBySuit := make(map[string][]int)
	for _, card := range cards {
		suit := GetCardSuit(card)
		splitBySuit[suit] = append(splitBySuit[suit], card)
	}

	for _, singleSuitCards := range splitBySuit {
		res := getStraightFlush(singleSuitCards)
		if len(res) > 0 {
			result = append(result, res...)
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetFiveOfGoldTigerFromPokerHand(cards []int, gold int) [][]int {
	if !HasFiveOfGoldTiger(cards, gold) {
		return nil
	}

	result := [][]int{}
	res := []int{}

	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		if rkValue == gold {
			res = append(res, card)
		}
	}

	result = append(result, res)
	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetFourOfAKindFromPokerHand(cards []int) [][]int {
	if !HasFourOfAKind(cards) {
		return nil
	}

	result := [][]int{}
	recordRKValue := make(map[int][]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRKValue[rkValue] = append(recordRKValue[rkValue], card)
	}

	for _, srv := range recordRKValue {
		if len(srv) >= 4 {
			allFourKind := Combinations(srv, 4)
			for _, kind := range allFourKind {
				result = append(result, kind)
			}
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetGoldFullHouseFromPokerHand(cards []int, gold int) [][]int {
	if !HasGoldFullHouse(cards, gold) {
		return nil
	}

	result := [][]int{}
	goldenCard := []int{}
	recordRKValues := make(map[int][]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		if rkValue == gold {
			goldenCard = append(goldenCard, card)
		} else {
			recordRKValues[rkValue] = append(recordRKValues[rkValue], card)
		}
	}

	goldenAllPossible := Combinations(goldenCard, 3)
	for _, golden := range goldenAllPossible {
		for _, values := range recordRKValues {
			if len(values) >= 2 {
				pairs := Combinations(values, 2)
				for _, p := range pairs {
					res := append(append([]int{}, golden...), p...)
					result = append(result, res)
				}
			}
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetFullHouseFromPokerHand(cards []int) [][]int {
	if !HasFullHouse(cards) {
		return nil
	}

	result := [][]int{}
	recordRKValues := make(map[int][]int)
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRKValues[rkValue] = append(recordRKValues[rkValue], card)
	}

	var allKinds [][]int
	var allPairs [][]int
	for _, values := range recordRKValues {
		if len(values) >= 3 {
			kinds := Combinations(values, 3)
			allKinds = append(allKinds, kinds...)
		}
		if len(values) >= 2 {
			pairs := Combinations(values, 2)
			allPairs = append(allPairs, pairs...)
		}
	}

	for _, kind := range allKinds {
		rkValueKind := ConvertRankToPoint(GetCardRank(kind[0]))
		for _, pair := range allPairs {
			rkValuePair := ConvertRankToPoint(GetCardRank(pair[0]))
			if rkValueKind != rkValuePair {
				res := append(append([]int{}, kind...), pair...)
				result = append(result, res)
			}
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetPairWithFlushFromPokerHand(cards []int) [][]int {
	if !HasPairWithFlush(cards) {
		return nil
	}

	result := [][]int{}
	recordSuit := make(map[string][]int)
	for _, card := range cards {
		suit := GetCardSuit(card)
		recordSuit[suit] = append(recordSuit[suit], card)
	}

	for _, values := range recordSuit {
		flushes := []int{}
		pairCard := -1
		for _, card := range values {
			if !ContainsInt(flushes, card) {
				flushes = append(flushes, card)
			} else {
				pairCard = card
				flushes = RemoveInt(flushes, card)
			}
		}

		if pairCard != -1 {
			flush := Combinations(flushes, 3)
			for _, fh := range flush {
				res := []int{pairCard, pairCard}
				res = append(res, fh...)
				result = append(result, res)
			}
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetFlushFromPokerHand(cards []int) [][]int {
	if !HasFlush(cards) {
		return nil
	}

	result := [][]int{}
	recordSuit := make(map[string][]int)
	for _, card := range cards {
		suit := GetCardSuit(card)
		recordSuit[suit] = append(recordSuit[suit], card)
	}

	for _, values := range recordSuit {
		if len(values) >= 5 {
			suitAllPossibles := Combinations(values, 5)
			for _, suitAP := range suitAllPossibles {
				res := append([]int{}, suitAP...)
				result = append(result, res)
			}
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetStraightFromPokerHand(cards []int) [][]int {
	if !HasStraight(cards) {
		return nil
	}
	result := [][]int{}

	points := make([]map[int]bool, 14+1)
	for i := range points {
		points[i] = make(map[int]bool)
	}

	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		points[rkValue][card] = true
	}

	cnt := 0
	prev := -1
	straights := [][]int{}
	for i := 1; i < len(points); i++ {
		if len(points[i]) == 0 {
			cnt = 0
			prev = -1
		} else {
			if prev == -1 {
				cnt = 1
				prev = i
			} else {
				if prev != i-1 {
					cnt = 1
					prev = i
				} else {
					cnt++
					prev = i
				}
			}

			if cnt >= 5 {
				straights = append(straights, []int{i - 4, i - 3, i - 2, i - 1, i})
			}
		}
	}

	checkSpecialStraight := true
	specialStraight := []int{2, 3, 4, 5, 14}
	for _, ss := range specialStraight {
		if len(points[ss]) == 0 {
			checkSpecialStraight = false
			break
		}
	}

	if checkSpecialStraight {
		straights = append(straights, specialStraight)
	}

	for _, ss := range straights {
		cs := make([][]int, len(ss))
		for i, s := range ss {
			cs[i] = MapKeys(points[s])
		}

		allPossibles := CartesianProduct(cs...)
		for _, ap := range allPossibles {
			result = append(result, ap)
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetGoldThreeOfAKindFromPokerHand(cards []int, gold int) [][]int {
	if !HasGoldThreeOfAKind(cards, gold) {
		return nil
	}
	result := [][]int{}
	golden := []int{}
	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		if rkValue == gold {
			golden = append(golden, card)
		}
	}

	allPossibles := Combinations(golden, 3)
	for _, ap := range allPossibles {
		result = append(result, ap)
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetThreeOfAKindFromPokerHand(cards []int) [][]int {
	if !HasThreeOfAKind(cards) {
		return nil
	}
	result := [][]int{}
	recordRKValues := make(map[int][]int)

	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRKValues[rkValue] = append(recordRKValues[rkValue], card)
	}

	for _, poker := range recordRKValues {
		if len(poker) >= 3 {
			allPossibles := Combinations(poker, 3)
			for _, ap := range allPossibles {
				result = append(result, ap)
			}
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetTwoPairsFromPokerHand(cards []int) [][]int {
	if !HasTwoPairs(cards) {
		return nil
	}
	result := [][]int{}
	recordRKValues := make(map[int][]int)

	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRKValues[rkValue] = append(recordRKValues[rkValue], card)
	}

	pairs := [][]int{}
	for _, poker := range recordRKValues {
		if len(poker) >= 2 {
			pairs = append(pairs, poker)
		}
	}

	twoPairs := CombinationsPairs(pairs, 2)
	for _, twoPair := range twoPairs {
		tp0 := Combinations(twoPair[0], 2)
		tp1 := Combinations(twoPair[1], 2)
		for _, tp0Comb := range tp0 {
			for _, tp1Comb := range tp1 {
				res := MergeSlices(tp0Comb, tp1Comb)
				result = append(result, res)
			}
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetOnePairFromPokerHand(cards []int) [][]int {
	if !HasOnePair(cards) {
		return nil
	}

	result := [][]int{}
	recordRKValues := make(map[int][]int)

	for _, card := range cards {
		rank := GetCardRank(card)
		rkValue := ConvertRankToPoint(rank)
		recordRKValues[rkValue] = append(recordRKValues[rkValue], card)
	}

	pairs := [][]int{}
	for _, poker := range recordRKValues {
		if len(poker) >= 2 {
			pairs = append(pairs, poker)
		}
	}

	for _, pair := range pairs {
		allPossible := Combinations(pair, 2)
		for _, ap := range allPossible {
			result = append(result, ap)
		}
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}

func GetHighCardFromPokerHand(cards []int) [][]int {
	if !HasHighCard(cards) {
		return nil
	}

	result := [][]int{}
	for _, card := range cards {
		result = append(result, []int{card})
	}

	sortedResult := GetSort2DIntArrayReverse(result)
	return sortedResult
}
