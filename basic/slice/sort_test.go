package slice

import (
	"sort"
	"testing"
)

type SortStruct struct {
	person string
	height int
	weight int
}

func TestSortSlice(t *testing.T) {
	sortStructs := []*SortStruct{
		{person: "xx", height: 3, weight: 2},
		{person: "xx", height: 1, weight: 2},
		{person: "xx", height: 2, weight: 2},
	}

	sort.Slice(sortStructs, func(i, j int) bool {
		return sortStructs[i].height < sortStructs[j].height
	})

	for _, sortStruct := range sortStructs {
		t.Log(sortStruct)
	}
}

func TestSliceFunc(t *testing.T) {
	s := []int{1, 2, 3, 4}
	changeSlice(s)
	t.Log(s)
}

func changeSlice(s []int) {
	s[0] = 99
}
