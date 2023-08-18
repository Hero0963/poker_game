package poker13tw

import (
	"testing"
	"time"
)

func TestGetPokerHandScore(t *testing.T) {

	testCases := []struct {
		cards    []int
		gold     int
		expected int
	}{
		{
			cards: []int{
				ConvertToCardValue("黑桃", "7"),
				ConvertToCardValue("黑桃", "7"),
				ConvertToCardValue("方塊", "7"),
				ConvertToCardValue("愛心", "7"),
				ConvertToCardValue("梅花", "7"),
			},
			gold:     4,
			expected: int(FiveKind)*CombatBase + 7,
		},
		{
			cards: []int{
				ConvertToCardValue("黑桃", "7"),
				ConvertToCardValue("黑桃", "8"),
				ConvertToCardValue("黑桃", "9"),
				ConvertToCardValue("黑桃", "T"),
				ConvertToCardValue("黑桃", "J"),
			},
			gold:     2,
			expected: int(StraightFlush)*CombatBase + 11,
		},

		{
			cards: []int{
				ConvertToCardValue("黑桃", "7"),
				ConvertToCardValue("方塊", "7"),
				ConvertToCardValue("愛心", "7"),
				ConvertToCardValue("梅花", "5"),
				ConvertToCardValue("方塊", "5"),
			},
			gold:     6,
			expected: int(FullHouse)*CombatBase + 705,
		},
	}

	for _, tc := range testCases {
		startTime := time.Now()
		result := GetPokerHandScore(tc.cards, tc.gold)
		if result != tc.expected {
			t.Errorf("Expected result to be %v, but got %v", tc.expected, result)
		}
		elapsed := time.Since(startTime)
		t.Logf("Test took %s\n", elapsed)
	}

}

func TestGetRecommendPokerHand(t *testing.T) {
	testCards := []int{
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("愛心", "6"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("梅花", "8"),
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "Q"),
		ConvertToCardValue("方塊", "Q"),
		ConvertToCardValue("愛心", "A"),
		ConvertToCardValue("梅花", "A"),
	}

	gold := 4
	result := GetRecommendPokerHand(testCards, gold)
	t.Logf("Result: %v\n", result)
	for _, res := range result {
		head, middle, tail := res[0], res[1], res[2]
		t.Logf("Result: %v, %v, %v\n", head, middle, tail)
		for _, card := range head {
			suit := GetCardSuit(card)
			rank := GetCardRank(card)
			t.Logf("%v, %v\n", suit, rank)
		}

		for _, card := range middle {
			suit := GetCardSuit(card)
			rank := GetCardRank(card)
			t.Logf("%v, %v\n", suit, rank)
		}

		for _, card := range tail {
			suit := GetCardSuit(card)
			rank := GetCardRank(card)
			t.Logf("%v, %v\n", suit, rank)
		}

	}
}

func TestGetNewSplitDeck(t *testing.T) {
	seed := time.Now().UnixNano()
	randomMgr := RandomManager{}
	randomMgr.SetRandSeed(seed)

	for i := 0; i < 20; i++ {
		result := GetNewSplitDeck(&randomMgr)
		t.Logf("Result #%d (i = %d): %v\n", i+1, i, result)
	}
}
