package poker13tw

import (
	"testing"
)

func TestOrder553(t *testing.T) {
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
	result := Order553(testCards, gold)
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
