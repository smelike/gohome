package main

// Reference: https://lowcode.life/a-tour-of-go-exercises-equivalent-binary-trees/
// There is still one (known) issue with this solution, it doesn’t deal
// with binary trees with a different number of nodes.
import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Left tree - value - right tree
	var worker func(t *tree.Tree)

	worker = func(t *tree.Tree) {
		if t.Left != nil {
			worker(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			worker(t.Right)
		}
	}
	worker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	// 使用 Walk 函数将 binary tree 中的值，都写入到 channel 中。
	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Same:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Different:", Same(tree.New(1), tree.New(2)))
}
