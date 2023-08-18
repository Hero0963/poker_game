package poker13tw

import (
	"reflect"
	"testing"
	"time"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		arr      []int
		r        int
		expected [][]int
	}{
		{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for _, test := range tests {
		startTime := time.Now()
		result := Combinations(test.arr, test.r)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test failed: %+v, Result: %+v", test, result)
		}

		elapsed := time.Since(startTime)
		t.Logf("Test took %s\n", elapsed)
	}
}

func TestInSlice(t *testing.T) {
	tests := []struct {
		elem     int
		arr      []int
		expected bool
	}{
		{2, []int{1, 2, 3}, true},
		{5, []int{4, 6, 8}, false},
		{0, []int{}, false},
	}

	for _, test := range tests {
		result := InSlice(test.elem, test.arr)
		if result != test.expected {
			t.Errorf("Test failed: elem=%d, arr=%v, expected=%v, got=%v", test.elem, test.arr, test.expected, result)
		}
	}
}

func TestCompareIntSlicesReverse(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, false},
		{[]int{4, 5, 6}, []int{7, 8, 9}, false},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
	}

	for _, test := range tests {
		result := CompareIntSlicesReverse(test.a, test.b)
		if result != test.expected {
			t.Errorf("Test failed: a=%v, b=%v, expected=%v, got=%v", test.a, test.b, test.expected, result)
		}
	}
}

func TestSort2DIntArrayReverse(t *testing.T) {
	tests := []struct {
		arr      [][]int
		expected [][]int
	}{
		{
			[][]int{
				{29, 3, 3, 16, 42},
				{6, 6, 19, 45, 32},
			},
			[][]int{
				{45, 32, 19, 6, 6},
				{42, 29, 16, 3, 3},
			},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		startTime := time.Now()
		sortedArr := GetSort2DIntArrayReverse(test.arr)
		if !reflect.DeepEqual(sortedArr, test.expected) {
			t.Errorf("Test failed: expected=%v, got=%v", test.expected, sortedArr)
		}
		elapsed := time.Since(startTime)
		t.Logf("Test took %s\n", elapsed)
	}
}
