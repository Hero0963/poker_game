package poker13tw

type PokerService struct{}

var _ PokerAPI = (*PokerService)(nil)

func NewPokerService() PokerAPI {
	return &PokerService{}
}

func (s *PokerService) GetPokerHandScore(cards []int, goldCard int) int {
	gold := ConvertCardToPoint(goldCard)
	return getPokerHandScore(cards, gold)
}

func (s *PokerService) GetRecommendPokerHand(cards []int, goldCard int) [][][]int {
	gold := ConvertCardToPoint(goldCard)
	return Order553(cards, gold)
}

func (s *PokerService) GetNewSplitDeck(rm *RandomManager) [][]int {
	newDeck := GetNewDeck(rm)

	gold := newDeck[0]
	h1 := newDeck[1 : 16+1]
	h2 := newDeck[17 : 32+1]
	h3 := newDeck[33 : 48+1]
	h4 := newDeck[49 : 64+1]

	return [][]int{{gold}, h1, h2, h3, h4}
}
