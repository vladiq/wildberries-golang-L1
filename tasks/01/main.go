package main

import "fmt"

type Human struct {
	name     string
	heightCm int
}

func (h *Human) Sing() {
	fmt.Printf("I'm only %s after all.\n", h.name)
}

func (h *Human) GrowOneCm() {
	h.heightCm += 1
}

func (h *Human) Yell() {
	fmt.Println("AAAAAAAAAAAAAAAaaaa")
}

type Action struct {
	Human
	runningSpeed int
}

func (a Action) Run() {
	fmt.Printf("Running at the speed of %d km/h!\n", a.runningSpeed)
}

func (a Action) TellHeight() {
	fmt.Printf("I'm %d cm tall.\n", a.heightCm)
}

func (a Action) Yell() {
	fmt.Println("OOOOOOOOOOOOOOOOooooooooo")
}

func main() {
	a := Action{
		Human{
			name:     "Joseph",
			heightCm: 185,
		},
		10,
	}
	a.Sing()
	a.Yell()
	a.Human.Yell()
	a.Run()
	a.TellHeight()
	a.GrowOneCm()
	a.TellHeight()
}
