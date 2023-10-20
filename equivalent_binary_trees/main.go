package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type PrintableChannel struct {
	ch chan int
}

func (p PrintableChannel) String() string {
	result := "{ "
	for {
		num, open := <-p.ch
		if !open {
			result = result[:len(result)-2] + " }"
			break
		}
		result += fmt.Sprintf("%d, ", num)
	}

	return result
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	t1ch := make(chan int)
	t2ch := make(chan int)
	go func() {
		Walk(t1, t1ch)
		close(t1ch)
	}()

	go func() {
		Walk(t2, t2ch)
		close(t2ch)
	}()

	for {
		x, b := <-t1ch
		if !b {
			break
		}
		y := <-t2ch
		if x != y {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()
	fmt.Println(PrintableChannel{ch})

	fmt.Println(Same(tree.New(5), tree.New(5)))
	fmt.Println(Same(tree.New(1), tree.New(5)))
}
