package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a
	fmt.Println("pointer address &a:", pa)
	fmt.Println("pointer value *pa:", *pa)
	fmt.Println("--")

	m := map[string]bool{"aaa": true, "bbb": true}
	fmt.Println(m)
	v, ok := m["aaa"]
	fmt.Println("aaa:", v, ":", ok)
	v, ok = m["ccc"]
	fmt.Println("ccc:", v, ":", ok)
	v = m["bbb"]
	fmt.Println("bbb:", v)
	v = m["ddd"]
	fmt.Println("ddd:", v)
	fmt.Println("--")

	u := user{name: "eee fff", score: 100}
	fmt.Println(u)
	fmt.Printf("before hit() score: %d\n", u.score)
	u.hit()
	fmt.Printf("after hit() score: %d\n", u.score)
	fmt.Println("-")
	fmt.Printf("before show() score: %d\n", u.score)
	u.show()
	fmt.Printf("after show() score: %d\n", u.score)
}

type user struct {
	name string
	score int
}

func (u user) show() {
	fmt.Printf("before inside show() score: %d\n", u.score)
	fmt.Printf("name: %s, score: %d\n", u.name, u.score)
	u.score++
	fmt.Printf("after inside show() score: %d\n", u.score)
}

func (u *user) hit() {
	fmt.Printf("before inside hit() score: %d\n", u.score)
	u.score++
	fmt.Printf("after inside hit() score: %d\n", u.score)
}