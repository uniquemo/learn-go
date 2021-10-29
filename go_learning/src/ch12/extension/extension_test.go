package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Println("旺")
}

func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("John")
}

type Cat struct {
	Pet
}

func (c *Cat) Speak() {
	fmt.Println("喵")
}

func TestCat(t *testing.T) {
	cat := new(Cat)
	cat.SpeakTo("Mary")
}
