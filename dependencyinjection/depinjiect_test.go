package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "lx") //这里用取地址的原因是：bytes.Buffer没有实现io.Writer接口，而*bytes.Buffer才实现了这个接口，所以需要传入指针

	got := buffer.String()
	want := "Hello lx"

	if got != want {
		t.Errorf("want '%s', but got '%s'", want, got)
	}
}
