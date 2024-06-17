package main

import (
	"syscall/js"
)

func fibImpl(n int) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fibImpl(n-1) + fibImpl(n-2)
}

func Fibonacci(_ js.Value, args []js.Value) any {
	if len(args) != 1 {
		return -1
	}
	n := args[0].Int()
	return fibImpl(n)
}

func main() {
	//log := js.Global().Get("console").Get("log")
	//log.Invoke("HELLO WORLD")

	done := make(chan int, 0)
	js.Global().Set("fibonacci", js.FuncOf(Fibonacci))
	<-done
}
