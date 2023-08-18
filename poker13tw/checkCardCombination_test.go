package poker13tw

import (
	"testing"
	"time"
)

func TestHasFiveOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),

		ConvertToCardValue("方塊", "A"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("方塊", "8"),
	}

	result := HasFiveOfAKind(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestHasStraightFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),

		ConvertToCardValue("方塊", "A"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("方塊", "8"),
	}

	result := HasStraightFlush(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestHasFiveOfGoldTiger(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "3"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("黑桃", "8"),
	}
	gold := 6
	result := HasFiveOfGoldTiger(testCards, gold)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasFourOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("愛心", "5"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("愛心", "6"),
		ConvertToCardValue("方塊", "2"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("愛心", "4"),
	}
	result := HasFourOfAKind(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasGoldFullHouse(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "4"),
	}
	gold := 7
	result := HasGoldFullHouse(testCards, gold)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasFullHouse(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "4"),
	}
	result := HasFullHouse(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasPairWithFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("愛心", "6"),
		ConvertToCardValue("愛心", "5"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "4"),
	}
	result := HasPairWithFlush(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("黑桃", "3"),
		ConvertToCardValue("方塊", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
	}
	result := HasFlush(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasStraight(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "Q"),
		ConvertToCardValue("梅花", "K"),
		ConvertToCardValue("方塊", "A"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
	}
	result := HasStraight(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasGoldThreeOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
	}
	gold := 7
	result := HasGoldThreeOfAKind(testCards, gold)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasThreeOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
	}
	result := HasThreeOfAKind(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasTwoPairs(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "A"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
	}
	result := HasTwoPairs(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasOnePair(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
	}
	result := HasOnePair(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestHasHighCard(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "7"),
	}
	result := HasHighCard(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}
