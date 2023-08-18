package poker13tw

import (
	"testing"
	"time"
)

func TestUpdateOffset(t *testing.T) {
	tests := []struct {
		oriOffset int
		x         int
		expected  int
	}{
		{1, 1, 100 + 1},
	}

	for _, test := range tests {
		startTime := time.Now()
		result := UpdateOffset(test.oriOffset, test.x)
		if result != test.expected {
			t.Errorf("Test failed: %+v, Result: %d", test, result)
		}

		elapsed := time.Since(startTime)
		t.Logf("Test took %s\n", elapsed)
	}
}

func TestGetScoreFiveOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
	}

	result := GetScoreFiveOfAKind(testCards)
	expected := int(FiveKind)*CombatBase + 7

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreStraightFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
	}

	result := GetScoreStraightFlush(testCards)
	expected := int(StraightFlush)*CombatBase + 11

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreFiveGoldTiger(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("黑桃", "J"),
	}

	gold := 6

	result := GetScoreFiveGoldTiger(testCards, gold)
	expected := int(FiveOfGoldTiger)*CombatBase + 611

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreFourOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("愛心", "5"),
	}

	result := GetScoreFourOfAKind(testCards)
	expected := int(FourOfAKind)*CombatBase + 705

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreGoldFullHouse(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
	}

	gold := 7

	result := GetScoreGoldFullHouse(testCards, gold)
	expected := int(GoldFullHouse)*CombatBase + 705

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreFullHouse(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
	}

	result := GetScoreFullHouse(testCards)
	expected := int(FullHouse)*CombatBase + 705

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScorePairWithFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
	}

	result := GetScorePairWithFlush(testCards)
	expected := int(PairWithFlush)*CombatBase + 7060502

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("黑桃", "3"),
	}

	result := GetScoreFlush(testCards)
	expected := int(Flush)*CombatBase + 906050403

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreStraight(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "Q"),
		ConvertToCardValue("梅花", "K"),
		ConvertToCardValue("方塊", "A"),
	}

	result := GetScoreStraight(testCards)
	expected := int(Straight)*CombatBase + 14

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreGoldThreeOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),
	}

	gold := 7

	result := GetScoreGoldThreeOfAKind(testCards, gold)
	expected := int(GoldThreeOfAKind)*CombatBase + 71110

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreThreeOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	result := GetScoreThreeOfAKind(testCards)
	expected := int(ThreeOfAKind)*CombatBase + 31404

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreTwoPairs(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "A"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	result := GetScoreTwoPairs(testCards)
	expected := int(TwoPairs)*CombatBase + 140304

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreOnePair(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	result := GetScoreOnePair(testCards)
	expected := int(OnePair)*CombatBase + 3140402

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestGetScoreHighCard(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "7"),
	}

	result := GetScoreHighCard(testCards)
	expected := int(HighCard)*CombatBase + 1407060302

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}
