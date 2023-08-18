package poker13tw

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestGetFiveOfAKindFromPokerHand(t *testing.T) {
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

		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "4"),
	}

	result := GetFiveOfAKindFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),
	}

	e2 := []int{
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("梅花", "4"),
	}

	sort.Sort(sort.Reverse(sort.IntSlice(e2)))
	sort.Sort(sort.Reverse(sort.IntSlice(e1)))
	expected = append(expected, e1)
	expected = append(expected, e2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetStraightFlushFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),

		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("方塊", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "T"),
		ConvertToCardValue("方塊", "J"),
	}

	result := GetStraightFlushFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "8"),
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
	}

	e2 := []int{
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "8"),
	}

	e3 := []int{
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "8"),
		ConvertToCardValue("方塊", "9"),
	}

	e4 := []int{
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "T"),
	}

	e5 := []int{
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "T"),
		ConvertToCardValue("方塊", "J"),
	}

	sort.Sort(sort.Reverse(sort.IntSlice(e1)))
	sort.Sort(sort.Reverse(sort.IntSlice(e2)))
	sort.Sort(sort.Reverse(sort.IntSlice(e3)))
	sort.Sort(sort.Reverse(sort.IntSlice(e4)))
	sort.Sort(sort.Reverse(sort.IntSlice(e5)))
	expected = append(expected, e5)
	expected = append(expected, e4)
	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetFiveOfGoldTigerFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("黑桃", "A"),

		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "3"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("黑桃", "8"),
	}

	gold := 6

	result := GetFiveOfGoldTigerFromPokerHand(testCards, gold)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "6"),
	}

	sort.Sort(sort.Reverse(sort.IntSlice(e1)))
	expected = append(expected, e1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetFourOfAKindFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("黑桃", "A"),

		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "3"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("黑桃", "8"),
	}

	result := GetFourOfAKindFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("梅花", "7"),
		ConvertToCardValue("方塊", "7"),
	}

	sort.Sort(sort.Reverse(sort.IntSlice(e1)))
	expected = append(expected, e1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetGoldFullHouseFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("愛心", "5"),

		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("梅花", "T"),
	}

	gold := 7

	result := GetGoldFullHouseFromPokerHand(testCards, gold)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
	}

	e2 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("黑桃", "5"),
	}

	e3 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("愛心", "5"),
	}

	e4 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("黑桃", "5"),
	}

	e5 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("愛心", "5"),
	}

	e6 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("愛心", "5"),
	}

	e7 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),

		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("梅花", "T"),
	}

	expected = append(expected, e7)
	expected = append(expected, e6)
	expected = append(expected, e5)
	expected = append(expected, e4)
	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", sortedExpected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetFullHouseFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("梅花", "T"),
	}

	result := GetFullHouseFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("梅花", "5"),
		ConvertToCardValue("方塊", "5"),
	}

	e2 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("梅花", "T"),
	}

	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", sortedExpected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetPairWithFlushFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("方塊", "2"),
		ConvertToCardValue("方塊", "5"),
		ConvertToCardValue("方塊", "7"),
	}

	result := GetPairWithFlushFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
	}
	e2 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "J"),
	}

	e3 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("黑桃", "J"),
	}

	e4 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "J"),
	}

	expected = append(expected, e4)
	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetFlushFromPokerHand(t *testing.T) {
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
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "7"),
	}

	result := GetFlushFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "9"),
		ConvertToCardValue("黑桃", "6"),
		ConvertToCardValue("黑桃", "5"),
		ConvertToCardValue("黑桃", "4"),
		ConvertToCardValue("黑桃", "3"),
	}
	e2 := []int{
		ConvertToCardValue("方塊", "8"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "7"),
	}

	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetStraightFromPokerHand(t *testing.T) {
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

	result := GetStraightFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "Q"),
		ConvertToCardValue("梅花", "K"),
		ConvertToCardValue("方塊", "A"),
	}
	e2 := []int{
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "Q"),
		ConvertToCardValue("梅花", "K"),
		ConvertToCardValue("方塊", "A"),
	}

	e3 := []int{
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "Q"),
		ConvertToCardValue("梅花", "K"),
	}

	e4 := []int{
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("愛心", "Q"),
		ConvertToCardValue("梅花", "K"),
	}

	expected = append(expected, e4)
	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetGoldThreeOfAKindFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "T"),
		ConvertToCardValue("黑桃", "J"),
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("方塊", "7"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
	}

	gold := 7

	result := GetGoldThreeOfAKindFromPokerHand(testCards, gold)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "7"),
		ConvertToCardValue("愛心", "7"),
		ConvertToCardValue("方塊", "7"),
	}

	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetThreeOfAKindFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("黑桃", "3"),
		ConvertToCardValue("方塊", "6"),
		ConvertToCardValue("方塊", "9"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "T"),
		ConvertToCardValue("黑桃", "2"),
		ConvertToCardValue("梅花", "4"),
	}

	result := GetThreeOfAKindFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	e2 := []int{
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("黑桃", "3"),
	}

	e3 := []int{
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("黑桃", "3"),
	}

	e4 := []int{
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("黑桃", "3"),
	}

	e5 := []int{
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("梅花", "4"),
	}

	expected = append(expected, e5)
	expected = append(expected, e4)
	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetTwoPairsFromPokerHand(t *testing.T) {
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

	result := GetTwoPairsFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "A"),
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	e2 := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "A"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "4"),
	}

	e3 := []int{
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
		ConvertToCardValue("方塊", "4"),
		ConvertToCardValue("愛心", "4"),
	}

	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetOnePairFromPokerHand(t *testing.T) {
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

	result := GetOnePairFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("黑桃", "2"),
	}

	e2 := []int{
		ConvertToCardValue("梅花", "3"),
		ConvertToCardValue("方塊", "3"),
	}

	e3 := []int{
		ConvertToCardValue("愛心", "4"),
		ConvertToCardValue("方塊", "4"),
	}

	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}

func TestGetHighCardFromPokerHand(t *testing.T) {
	startTime := time.Now()

	testCards := []int{
		ConvertToCardValue("黑桃", "A"),
		ConvertToCardValue("愛心", "2"),
		ConvertToCardValue("愛心", "3"),
		ConvertToCardValue("梅花", "6"),
		ConvertToCardValue("方塊", "7"),
	}

	result := GetHighCardFromPokerHand(testCards)
	expected := [][]int{}
	e1 := []int{
		ConvertToCardValue("黑桃", "A"),
	}

	e2 := []int{
		ConvertToCardValue("愛心", "2"),
	}

	e3 := []int{
		ConvertToCardValue("愛心", "3"),
	}
	e4 := []int{
		ConvertToCardValue("梅花", "6"),
	}

	e5 := []int{
		ConvertToCardValue("方塊", "7"),
	}
	expected = append(expected, e5)
	expected = append(expected, e4)
	expected = append(expected, e3)
	expected = append(expected, e2)
	expected = append(expected, e1)
	sortedExpected := GetSort2DIntArrayReverse(expected)

	if !reflect.DeepEqual(result, sortedExpected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	elapsed := time.Since(startTime)
	t.Logf("Test took %s\n", elapsed)

}
