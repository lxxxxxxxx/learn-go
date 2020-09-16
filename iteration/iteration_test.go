package iteration

import (
	"fmt"
	"testing"
	"strings"
)

func Repeat(str string,n int) string{
	var repeated string
	for i:=0;i<n;i++{
		repeated+=str
	}
	return repeated
}

func TestRepeat(t* testing.T){
	got :=Repeat("a",10);
	want:="aaaaaaaaaa"
	if got != want{
		t.Errorf("want '%q',but got '%q'",want,got)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i:=0;i<b.N;i++{
		Repeat("a",100)
	}
}


func ExampleRepeat(){
	repeated := Repeat("hello,",5)
	fmt.Println(repeated)
}