package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a
	fmt.Println("pointer address &a: " + pa)
	fmt.Println("pointer value *pa: " + *pa)

	m := map[aaa:true bbb:true]
	v, ok := m["aaa"]
	fmt.Println("aaa: " + v + " : " + ok)
	fmt.Println(m)
}