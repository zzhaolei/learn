package main

import (
	"fmt"
	"reflect"
	"time"
)

type clientInterface interface {
	LearnGo()
	LearnC()
}

type client1 struct{}

func (c *client1) LearnGo() {}
func (c *client1) LearnC()  {}

type client2 struct{}

func (c *client2) LearnGo() {}
func (c *client2) LearnC()  {}

type S[T clientInterface] struct {
	value T
}

// 测试函数
func (s S[T]) test(v T) {
	fmt.Println("==========", reflect.TypeOf(s.value), reflect.TypeOf(v))
}

func NewS[T clientInterface](client T) S[T] {
	return S[T]{
		value: client,
	}
}

func main() {
	one := client1{}
	two := client2{}

	s1 := NewS[*client1](&one)
	s1.test(&one)

	s2 := NewS[*client2](&two)
	s2.test(&two)

	timer := time.NewTimer(time.Second * 10)
	timer.Reset(time.Second * 1)
	<-timer.C
}
