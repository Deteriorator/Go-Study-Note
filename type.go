/*
 * @program: Go-Study-Note
 * @author: Leon
 * @create: 2020-09-09 23:11
 */
package main

// var 定义变量
var x int
var f float32 = 1.6
var s = "abc"

// ---------------------

// 一次定义多个变量
var t, y, z int
var n, m = "abc", 123

var (
	a int
	b float32
)

func main() {
	// 函数内部使用 :=
	x := 123
	e, r := 0x1234, "Hello, World!"
	println(x, e, r)
}
