package poker13tw

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCardSuit(t *testing.T) {
	tests := []struct {
		card     int
		expected string
	}{
		{1, "黑桃"},
		{14, "愛心"},
		{27, "方塊"},
		{40, "梅花"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Card %d", test.card), func(t *testing.T) {
			startTime := time.Now()

			result := GetCardSuit(test.card)
			if result != test.expected {
				t.Errorf("For card %d, expected %s, but got %s", test.card, test.expected, result)
			}

			elapsed := time.Since(startTime)
			t.Logf("Test took %s\n", elapsed)
		})
	}
}

func TestGetCardRankt(t *testing.T) {
	tests := []struct {
		card     int
		expected string
	}{
		{1, "2"},
		{13, "A"},
		{25, "K"},
		{40, "2"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Card %d", test.card), func(t *testing.T) {
			startTime := time.Now()

			result := GetCardRank(test.card)
			if result != test.expected {
				t.Errorf("For card %d, expected %s, but got %s", test.card, test.expected, result)
			}

			elapsed := time.Since(startTime)
			t.Logf("Test took %s\n", elapsed)
		})
	}
}

func TestConvertToCardValue(t *testing.T) {
	tests := []struct {
		suit     string
		rank     string
		expected int
	}{
		{"黑桃", "2", 1},
		{"黑桃", "Q", 11},
		{"黑桃", "A", 13},
		{"愛心", "9", 21},
	}

	for _, test := range tests {
		startTime := time.Now()
		result := ConvertToCardValue(test.suit, test.rank)
		if result != test.expected {
			t.Errorf("Test failed: %+v, Result: %d", test, result)
		}

		elapsed := time.Since(startTime)
		t.Logf("Test took %s\n", elapsed)
	}
}

func TestConvertRankToPoint(t *testing.T) {
	tests := []struct {
		rank     string
		expected int
	}{
		{"2", 2},
		{"Q", 12},
		{"A", 14},
		{"9", 9},
	}

	for _, test := range tests {
		startTime := time.Now()
		result := ConvertRankToPoint(test.rank)
		if result != test.expected {
			t.Errorf("Test failed: %+v, Result: %d", test, result)
		}

		elapsed := time.Since(startTime)
		t.Logf("Test took %s\n", elapsed)
	}

}

func TestConvertCardToPoint(t *testing.T) {
	tests := []struct {
		card     int
		expected int
	}{
		{1, 2},
		{13, 14},
		{26, 14},
		{43, 5},
	}

	for _, test := range tests {
		startTime := time.Now()
		result := ConvertCardToPoint(test.card)
		if result != test.expected {
			t.Errorf("Test failed: %+v, Result: %d", test, result)
		}

		elapsed := time.Since(startTime)
		t.Logf("Test took %s\n", elapsed)
	}

}
