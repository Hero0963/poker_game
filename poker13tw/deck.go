package poker13tw

import (
	"math/rand"
)

type RandomManager struct {
	randSeed int64
	randGen  *rand.Rand
}

func (rm *RandomManager) SetRandSeed(seed int64) {
	rm.randSeed = seed
	source := rand.NewSource(seed)
	rm.randGen = rand.New(source)
}

func (rm *RandomManager) GetRandSeed() int64 {
	return rm.randSeed
}

func GetNewDeck(rm *RandomManager) []int {
	deck := make([]int, 0)
	for i := 1; i <= 52; i++ {
		deck = append(deck, i)
	}

	for i := 1; i <= 13; i++ {
		deck = append(deck, i)
	}

	for i := 0; i < 1200; i++ {
		rm.randGen.Shuffle(len(deck), func(i, j int) {
			deck[i], deck[j] = deck[j], deck[i]
		})
	}

	newDeck := make([]int, len(deck))
	copy(newDeck, deck)

	return newDeck

}
