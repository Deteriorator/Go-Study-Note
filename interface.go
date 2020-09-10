/*
 * @program: Go-Study-Note
 * @author: Leon
 * @create: 2020-09-10 22:59
 */
package main

import "fmt"

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func Print(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

func main() {
	var t Printer = &User{1, "Tom"}
	t.Print()
	Print(1)
	Print("hello, world!")
}
