package main

import "fmt"

type A struct {
	name string
	aSex *B
	age *string
}

type B struct {
	sex string
}

func main(){
	age := "12"
	a := A{
		"zhang",
		&B{
			"man",
		},
		&age,
	}

	b := a

	b.aSex.sex = "woman"
	bage :="17"
	b.age = &bage
	fmt.Printf("a.sex:%s a.age:%v b.sex:%s b.age:%v", a.aSex.sex, *a.age, b.aSex.sex, *b.age)
}
