package structandinterface

import "math"

// Shape 形状接口
type Shape interface{
	Area() float64
}

//Rectangle 矩形结构体
type Rectangle struct{
	w float64
	h float64
}

//Circle 圆形结构体
type Circle struct{
	Radius float64
}

//Triangle 三角形结构体
type Triangle struct{
	a float64
	b float64
}

//Area : 计算矩形面积
func (r Rectangle)Area()float64{
	return r.w*r.h
}

//Area : 计算圆的面积
func (c Circle) Area()float64{
	return math.Pi*c.Radius*c.Radius
}

//Area ： 计算三角形面积
func (t Triangle)Area()float64{
	return t.a*t.b/2
}