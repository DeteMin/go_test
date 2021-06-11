package main

import (
	"fmt"
	"github.com/imdario/mergo"
)

func main(){
	var m1 = map[string]string{
		"CN":"Level0",
		"US":"Level1",
	}

	var m2 = map[string]string{
		"CNAA":"Delay",
		"CNAB":"Delay",
	}

	if err := mergo.Merge(&m1, m2); err != nil {
		fmt.Printf("err:%v",err)
		return
	}
	fmt.Printf("m1:%v", m1)
}


