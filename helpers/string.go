package helpers

import "sort"

func StringSortAndSearch(array []string, target string) (bool, int) {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	if index >= len(array) || array[index] != target {
		return false, index
	}
	return true, index
}
