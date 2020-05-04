package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a
	fmt.Println("pointer address &a:", pa)
	fmt.Println("pointer value *pa:", *pa)

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
}