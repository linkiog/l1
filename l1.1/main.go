package main

import "fmt"

type Human struct {
	Name string
	Age  int
}
type Action struct {
	Human
}

func (h *Human) Speak() {
	fmt.Printf("Hello my name is %s. I'm %d\n", h.Name, h.Age)
}
func (a *Action) Run() {
	fmt.Printf("%s is running\n", a.Name)
}

func main() {
	action := Action{
		Human: Human{
			Name: "linkiog",
			Age:  22,
		},
	}
	action.Speak()
	action.Run()
}
