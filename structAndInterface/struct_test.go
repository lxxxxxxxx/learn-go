package structandinterface

import "testing"

func checkRes(t *testing.T,want,got float64){
	if got != want{
		t.Errorf("want '%.2f',but got '%.2f'",want,got)
	}
}

func checkArea(t *testing.T,name string,shape Shape,want float64){
	got:=shape.Area()

	if want != got{
		t.Errorf("%#v ,want '%.2f',but got '%.2f'",name,want,got)
	}
}


func TestArea(t *testing.T){
	t.Run("test RectangleArea",func(t *testing.T){
		rect :=Rectangle{2.0,2.0}
		got := rect.Area()
		want := 4.0
	
		checkRes(t,want,got)
	})

	t.Run("test circles area",func(t *testing.T){
		circle := Circle{10}
		want :=314.1592653589793

		checkArea(t,"circle",circle,want)
	})

	areaTests := []struct{
		name string
		shape Shape
		want1 float64
	}{
		{name:"Rectangle",shape:Rectangle{w:10.0,h:20.0},want1:200.0},
		{name:"Circle",shape:Circle{Radius:10.0},want1:314.1592653589793},
		{name:"Triangle",shape:Triangle{a:12,b:6},want1:36.0},
	}

	for _,tt := range areaTests{
		checkArea(t,tt.name,tt.shape,tt.want1)
	}

}
