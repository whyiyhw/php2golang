package main

import (
	"testing"
	"fmt"
)

func TestFirstTry(t *testing.T) {
	t.Log("this is first test")
}

func TestBuildData(t *testing.T) {
	a, b := 1, 1
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		fmt.Println(b)
		a, b = b, a+b
	}
}

func TestIota(t *testing.T) {
	const (
		Monday = iota
		Tuesday
		Wednesday
	)
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	t.Log(Wednesday, Monday, Tuesday)
	a := 7
	t.Log(Readable&a == Readable, Writable&a == Writable, Executable&a == Executable)
}

func TestTypeSwitch(t *testing.T) {
	type myInt int
	var a int32
	var b int64
	a = 1
	b = 2
	a = int32(b)
	t.Log(a, b)
	var c myInt
	c = myInt(a)
	t.Log(c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Logf("%T *%s*", s, s)
	t.Log(len(s)) //0
}
