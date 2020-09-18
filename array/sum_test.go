package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assert := func(want, got int) {
		t.Helper()
		if got != want {
			t.Errorf("want %d ,but got %d", want, got)
		}
	}

	checkSum := func(t *testing.T,want []int,got []int){
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v ,but got %v", want, got)
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

	t.Run("test all sum", func(t *testing.T) {
		// want := AllSum([]int{1, 2, 3}, []int{4, 5, 6})
		want := AllSum1([]int{1, 2, 3}, []int{4, 5, 6})
		got := []int{6, 15}

		checkSum(t,want,got)
	})

	t.Run("test sum all tail",func(t* testing.T){
		got :=SumAllTial([]int{1,2,3},[]int{2,5,8,4},[]int{3,4,7,5})
		want:=[]int{5,17,16}

		checkSum(t,want,got)
	})

	t.Run("test empty slice",func(t* testing.T){
		got:=SumAllTial([]int{},[]int{1,2,5})
		want:=[]int{0,7}

		checkSum(t,want,got)
	})
}
