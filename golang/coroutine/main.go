package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("main panic:", e)
			panic(e)
		}
	}()
	next, _ := coro_Pull(func(yield func(string) bool) {
		yield("hello")
		panic("world")
	})
	for {
		fmt.Println(next())
	}
}

func coro_Pull[V any](push func(yield func(V) bool)) (pull func() (V, bool), stop func()) {
	copush := func(more bool, yield func(V) bool) V {
		if more {
			push(yield)
		}
		var zero V
		return zero
	}
	resume := coro_New(copush)
	pull = func() (V, bool) {
		return resume(true)
	}
	stop = func() {
		resume(false)
	}
	return pull, stop
}

type msg[T any] struct {
	panic any
	val   T
}

func coro_New[In, Out any](f func(In, func(Out) In) Out) (resume func(In) (Out, bool)) {
	cin := make(chan In)
	cout := make(chan msg[Out])
	running := true
	resume = func(in In) (out Out, ok bool) {
		if !running {
			return
		}
		cin <- in
		m := <-cout
		if m.panic != nil {
			panic(m.panic)
		}
		return m.val, running
	}
	yield := func(out Out) In {
		cout <- msg[Out]{val: out}
		return <-cin
	}
	go func() {
		defer func() {
			if running {
				running = false
				cout <- msg[Out]{panic: recover()}
			}
		}()
		out := f(<-cin, yield)
		running = false
		cout <- msg[Out]{val: out}
	}()
	return resume
}
