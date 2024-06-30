package sorting_test

import (
	"algos/sorting"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{13, 46, 24, 52, 20, 9}
	want := []int{9, 13, 20, 24, 46, 52}
	got := sorting.BubbleSort(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}
