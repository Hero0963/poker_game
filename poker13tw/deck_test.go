package poker13tw

import (
	"testing"
)

func TestGetNewDeck(t *testing.T) {
	randomMgr := &RandomManager{}
	randomMgr.SetRandSeed(123456) // 使用特定種子進行測試

	for i := 0; i < 10; i++ {
		newDeck := GetNewDeck(randomMgr)

		t.Logf("Deck #%d:", i+1)
		for j, card := range newDeck {
			t.Logf("Card %d: %d", j+1, card)
		}

		t.Log("------------------------------")
	}
}
