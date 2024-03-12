package main

import "fmt"

type Human struct {
	id int
	a  Action
}
type Action struct {
	move  int
	sleep int
	eat   int
}

func (h *Human) In() {
	h.id = 1
}
func main() {
	h := &Human{}
	h.In()
	h.a.eat = 2
	h.a.move = 3
	h.a.sleep = 4

	fmt.Println(h)
}
