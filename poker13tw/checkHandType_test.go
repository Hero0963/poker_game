package poker13tw

import (
	"testing"
	"time"
)

func TestIsFiveOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),
	}

	result := IsFiveOfAKind(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestIsStraightFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
	}

	result := IsStraightFlush(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestIsFiveOfGoldTiger(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("黑桃", "J"),
	}

	testGold := 6

	result := IsFiveOfGoldTiger(testCards, testGold)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestIsFourOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("愛心", "5"),
	}

	result := IsFourOfAKind(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestIsGoldFullHouse(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("愛心", "5"),
	}

	testGold := 7

	result := IsGoldFullHouse(testCards, testGold)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestIsFullHouse(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
	}

	result := IsFullHouse(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsPairWithFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
	}

	result := IsPairWithFlush(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsFlush(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("黑桃", "3"),
	}

	result := IsFlush(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsStraight(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "Q"),
		ConvertToCardValue("梅花", "K"),
		ConvertToCardValue("方塊", "A"),
	}

	result := IsStraight(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsGoldThreeOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),
	}

	testGold := 7

	result := IsGoldThreeOfAKind(testCards, testGold)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsThreeOfAKind(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	result := IsThreeOfAKind(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsTwoPairs(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "A"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	result := IsTwoPairs(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsOnePair(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	result := IsOnePair(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}

func TestIsHighCard(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "7"),
	}

	result := IsHighCard(testCards)
	expected := true

	if result != expected {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)
}
