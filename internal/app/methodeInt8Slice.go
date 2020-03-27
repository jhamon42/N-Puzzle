package main

import (
	"sort"
)

// Len : size
func (s Int8Slice) Len() int {
	return len(s)
}

// Less : for sort fonc
func (s Int8Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap : swap 2 elem for sort
func (s Int8Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Trim : trim 0
func (s Int8Slice) Trim() Int8Slice {
	var trim Int8Slice
	for _, key := range s {
		if key != 0 {
			trim = append(trim, key)
		}
	}
	return trim
}

// Index : index of value
func (s Int8Slice) Index(comp int8) int {
	for index, key := range s {
		if key == comp {
			return index
		}
	}
	return -1
}

// Sort is a convenience method.
func (s Int8Slice) Sort() {
	sort.Sort(s)
}
