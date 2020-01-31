package main

import (
	"fmt"
	"sync"
)

type shape interface {
	Diameter()
	Area(v int)
}

type ImplShape struct {
	pie float32
}

var (
	singleton *ImplShape
	once sync.Once
)

func InitSingleton() *ImplShape {
	once.Do(func() {
		singleton = &ImplShape{pie: 3.14}
	})
	return singleton
}

func (i *ImplShape) Area(v int)  {
	fmt.Println(i.pie * float32(v))
}



func main()  {
	InitSingleton()
	//o := &ImplShape{}
	singleton.Area(20)
}