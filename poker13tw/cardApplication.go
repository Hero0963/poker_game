package poker13tw

func GetPokerHandScore(cards []int, gold int) int {
	return getPokerHandScore(cards, gold)
}

func GetRecommendPokerHand(cards []int, gold int) [][][]int {
	return Order553(cards, gold)

}

func GetNewSplitDeck(rm *RandomManager) [][]int {
	newDeck := GetNewDeck(rm)

	gold := newDeck[0]
	h1 := newDeck[1 : 16+1]
	h2 := newDeck[17 : 32+1]
	h3 := newDeck[33 : 48+1]
	h4 := newDeck[49 : 64+1]

	return [][]int{{gold}, h1, h2, h3, h4}

}
