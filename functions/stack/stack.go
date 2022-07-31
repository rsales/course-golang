package main

import (
	"runtime/debug"
)

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}

//=====================
// goroutine 1 [running]:
// runtime/debug.Stack()
// 	/opt/homebrew/Cellar/go/1.18/libexec/src/runtime/debug/stack.go:24 +0x68
// runtime/debug.PrintStack()
// 	/opt/homebrew/Cellar/go/1.18/libexec/src/runtime/debug/stack.go:16 +0x20
// main.f3(...)
// 	/Users/rafaelsales/Development/study/course-golang/functions/stack/stack.go:8
// main.f2(...)
// 	/Users/rafaelsales/Development/study/course-golang/functions/stack/stack.go:12
// main.f1(...)
// 	/Users/rafaelsales/Development/study/course-golang/functions/stack/stack.go:16
// main.main()
// 	/Users/rafaelsales/Development/study/course-golang/functions/stack/stack.go:20 +0x2c
// start here and ended here too 👆👇
//=====================
