package poker13tw

import "sort"

// 單一指定墩贏某家且滿足牌型，輸的玩家需支付以下分數

// 頭墩
// 衝三 + 1
// 彩金衝三 +2

// 中墩
// 彩金三條 +1
// 彩金葫蘆 +2
// 鐵支 +2
// 同花順 +2
// 彩金五虎將 +5
// 五枚 +10

// 尾墩
// 彩金三條 +1
// 彩金葫蘆 +1
// 鐵支 +1
// 同花順 +1
// 彩金五虎將 +5
// 五枚 +5

func GetRecommendPoker(hc []int, tc []int, mc []int, gold int) [][]int {
	poker := [][]int{append([]int(nil), hc...), append([]int(nil), tc...), append([]int(nil), mc...)}
	sort.Slice(poker, func(i, j int) bool {
		combatI := GetPokerHandScore(poker[i], gold)
		combatJ := GetPokerHandScore(poker[j], gold)
		return combatI < combatJ
	})
	return poker
}

func TopUpCards(curPokerHand []int, unusedCard *[]int, number int) []int {
	for len(curPokerHand) < number {
		card := (*unusedCard)[len(*unusedCard)-1]
		*unusedCard = (*unusedCard)[:len(*unusedCard)-1]
		curPokerHand = append(curPokerHand, card)
	}
	return curPokerHand
}

func GetSortedUnusedCards(cards []int) []int {
	sortedCards := append([]int(nil), cards...)
	sort.Slice(sortedCards, func(i, j int) bool {
		return ConvertRankToPoint(GetCardRank(sortedCards[i])) < ConvertRankToPoint(GetCardRank(sortedCards[j]))
	})
	return sortedCards
}

func UpdateUnusedCards(usedCards []int, originalCards []int) []int {
	record := make(map[int]int)
	for _, card := range originalCards {
		record[card]++
	}
	for _, uc := range usedCards {
		record[uc]--
	}
	var updatedUnusedCards []int
	for card, count := range record {
		for i := 0; i < count; i++ {
			updatedUnusedCards = append(updatedUnusedCards, card)
		}
	}
	return updatedUnusedCards
}

func GetHeadCoreCards(cards []int, gold int) [][]int {
	goldThreeOfAKind := GetGoldThreeOfAKindFromPokerHand(cards, gold)
	if len(goldThreeOfAKind) > 0 {
		return goldThreeOfAKind
	}

	threeOfAKind := GetThreeOfAKindFromPokerHand(cards)
	if len(threeOfAKind) > 0 {
		return threeOfAKind
	}

	onePair := GetOnePairFromPokerHand(cards)
	if len(onePair) > 0 {
		return onePair
	}

	highCard := GetHighCardFromPokerHand(cards)
	if len(highCard) > 0 {
		return highCard
	}

	return nil
}

func GetMiddleTailCoreCards(cards []int, gold int) [][]int {
	fiveOfAKind := GetFiveOfAKindFromPokerHand(cards)
	if len(fiveOfAKind) > 0 {
		return fiveOfAKind
	}

	fiveOfGoldTiger := GetFiveOfGoldTigerFromPokerHand(cards, gold)
	if len(fiveOfGoldTiger) > 0 {
		return fiveOfGoldTiger
	}

	straightFlush := GetStraightFlushFromPokerHand(cards)
	fourOfAKind := GetFourOfAKindFromPokerHand(cards)
	if len(straightFlush) > 0 || len(fourOfAKind) > 0 {
		return append(straightFlush, fourOfAKind...)
	}

	goldFullHouse := GetGoldFullHouseFromPokerHand(cards, gold)
	if len(goldFullHouse) > 0 {
		return goldFullHouse
	}

	goldThreeOfAKind := GetGoldThreeOfAKindFromPokerHand(cards, gold)
	if len(goldThreeOfAKind) > 0 {
		return goldThreeOfAKind
	}

	fullHouse := GetFullHouseFromPokerHand(cards)
	if len(fullHouse) > 0 {
		return fullHouse
	}

	pairWithFlush := GetPairWithFlushFromPokerHand(cards)
	if len(pairWithFlush) > 0 {
		return pairWithFlush
	}

	flush := GetFlushFromPokerHand(cards)
	if len(flush) > 0 {
		return flush
	}

	straight := GetStraightFromPokerHand(cards)
	if len(straight) > 0 {
		return straight
	}

	threeOfAKind := GetThreeOfAKindFromPokerHand(cards)
	if len(threeOfAKind) > 0 {
		return threeOfAKind
	}

	twoPairs := GetTwoPairsFromPokerHand(cards)
	if len(twoPairs) > 0 {
		return twoPairs
	}

	onePair := GetOnePairFromPokerHand(cards)
	if len(onePair) > 0 {
		return onePair
	}

	highCard := GetHighCardFromPokerHand(cards)
	if len(highCard) > 0 {
		return highCard
	}

	return nil
}

func Order553(cards []int, gold int) [][][]int {
	result := make([][][]int, 0)
	unusedCards := append([]int(nil), cards...)
	middleCore := GetMiddleTailCoreCards(unusedCards, gold)
	for _, mc := range middleCore {
		unusedCardsM := UpdateUnusedCards(mc, unusedCards)
		tailCore := GetMiddleTailCoreCards(unusedCardsM, gold)
		for _, tc := range tailCore {
			unusedCardsMT := UpdateUnusedCards(tc, unusedCardsM)
			headCore := GetHeadCoreCards(unusedCardsMT, gold)
			for _, hc := range headCore {
				unusedCardsMTH := UpdateUnusedCards(hc, unusedCardsMT)
				sortedUnusedCards := GetSortedUnusedCards(unusedCardsMTH)
				completedHc := TopUpCards(hc, &sortedUnusedCards, 3)
				completedMc := TopUpCards(mc, &sortedUnusedCards, 5)
				completedTc := TopUpCards(tc, &sortedUnusedCards, 5)
				recommendedCard := GetRecommendPoker(completedMc, completedTc, completedHc, gold)

				if len(recommendedCard[0]) == 3 {
					result = append(result, recommendedCard)
				}
			}
		}
	}
	if len(result) > 0 {
		result = result[:1]
	}
	return result
}
