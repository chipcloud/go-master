package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()

	l.PushFront(1)

	l.PushBack(2)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
