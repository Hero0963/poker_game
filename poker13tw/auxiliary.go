package poker13tw

import "sort"

func Combinations(arr []int, r int) [][]int {
	var res [][]int
	var backtrack func(start, currLen int, currComb []int)

	backtrack = func(start, currLen int, currComb []int) {
		if currLen == r {
			comb := make([]int, len(currComb))
			copy(comb, currComb)
			res = append(res, comb)
			return
		}

		for i := start; i < len(arr); i++ {
			currComb = append(currComb, arr[i])
			backtrack(i+1, currLen+1, currComb)
			currComb = currComb[:len(currComb)-1]
		}
	}

	backtrack(0, 0, []int{})
	return res
}

func InSlice(elem int, arr []int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
func GetSort2DIntArrayReverse(arr [][]int) [][]int {
	// Create a copy of the 2D array
	arrCopy := make([][]int, len(arr))
	for i, a := range arr {
		arrCopy[i] = make([]int, len(a))
		copy(arrCopy[i], a)
	}

	// Sort each internal 1D array in reverse order
	for _, a := range arrCopy {
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
	}

	// Sort the copied array in reverse order
	sort.Slice(arrCopy, func(i, j int) bool {
		return CompareIntSlicesReverse(arrCopy[i], arrCopy[j])
	})

	return arrCopy
}

func CompareIntSlicesReverse(a, b []int) bool {
	lenA, lenB := len(a), len(b)
	minLen := lenA
	if lenB < minLen {
		minLen = lenB
	}

	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return a[i] > b[i] // Reverse comparison order
		}
	}

	// If the common part is the same, longer slice is considered "greater"
	return lenA > lenB
}

func recoverOriginalPointByRKValueAndSuit(rkValue int, suitRecover int) int {
	return rkValue - 1 + 13*suitRecover
}

// ContainsInt checks if an int slice contains a specific value.
func ContainsInt(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// RemoveInt removes a specific value from an int slice and returns the modified slice.
func RemoveInt(slice []int, val int) []int {
	var result []int
	for _, v := range slice {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

// MapKeys returns keys of a map as a slice of integers.
func MapKeys(m map[int]bool) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// CartesianProduct returns the Cartesian product of multiple slices.
func CartesianProduct(slices ...[]int) [][]int {
	if len(slices) == 0 {
		return [][]int{}
	}
	if len(slices) == 1 {
		result := [][]int{}
		for _, val := range slices[0] {
			result = append(result, []int{val})
		}
		return result
	}

	rest := CartesianProduct(slices[1:]...)
	result := [][]int{}
	for _, val := range slices[0] {
		for _, r := range rest {
			newCombination := append([]int{val}, r...)
			result = append(result, newCombination)
		}
	}
	return result
}

func MergeSlices(slice1, slice2 []int) []int {
	result := append([]int{}, slice1...)
	result = append(result, slice2...)
	return result
}

func CombinationsPairs(arr [][]int, r int) [][][]int {
	var res [][][]int
	var backtrack func(start, currLen int, currComb [][]int)

	backtrack = func(start, currLen int, currComb [][]int) {
		if currLen == r {
			comb := make([][]int, len(currComb))
			for i, v := range currComb {
				comb[i] = make([]int, len(v))
				copy(comb[i], v)
			}
			res = append(res, comb)
			return
		}

		for i := start; i < len(arr); i++ {
			currComb = append(currComb, arr[i])
			backtrack(i+1, currLen+1, currComb)
			currComb = currComb[:len(currComb)-1]
		}
	}

	backtrack(0, 0, [][]int{})
	return res
}