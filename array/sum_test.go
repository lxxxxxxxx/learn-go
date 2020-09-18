package array

import (
	"testing"
	"reflect"
)

func Sum(slice []int) (sum int) {
	for _, num := range slice {
		sum += num
	}
	return
}

func AllSum(sl)

func TestSum(t *testing.T) {

	assert := func(want, got int) {
		t.Helper()
		if got != want {
			t.Errorf("want %d ,but got %d", want, got)
		}
	}

	t.Run("array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}

		want := 15
		got := Sum(arr)

		assert(want, got)
	})

	t.Run("slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7}

		want := 28
		got := Sum(slice)

		assert(want, got)
	})
}

func TestAllSum(t *testing.T){
	want := AllSum([]int{1,2,3},[]int[4,5,6])
	got := []int{6,15}

	if !reflect.DeepEqual(want,got){
		t.Errorf("want %v ,but got %v",want,got)
	}
}
