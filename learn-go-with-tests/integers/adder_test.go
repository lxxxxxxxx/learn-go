package integers

import (
	"fmt"
	"testing"
)

func Add(a,b int) int{
	return a+b
}

func TestAdder(t *testing.T){
	sum := 4
	got := Add(2,2)

	if sum != got{
		t.Errorf("expected '%d',but got '%d'",sum,got)
	}
}

func ExampleAdd(){
	sum := Add(1,3)
	fmt.Println(sum)
	//Output:4
}