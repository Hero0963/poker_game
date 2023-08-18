package poker13tw

type PokerAPI interface {

	// 計算每一墩的分數
	GetPokerHandScore(cards []int, goldCard int) int

	// 返回一組推薦的組法，格式為 頭墩/中墩/尾墩
	GetRecommendPokerHand(cards []int, goldCard int) [][][]int

	// 返回一副新發的牌，格式為 彩金牌/玩家A手牌/玩家B手牌/玩家C手牌/玩家D手牌
	GetNewSplitDeck(rm *RandomManager) [][]int
}
