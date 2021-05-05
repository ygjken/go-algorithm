package main

import (
	"fmt"

	"github.com/ygjken/go-algorithm/pkg/binarytree"
	"github.com/ygjken/go-algorithm/pkg/readdata"
)

func main() {
	var tree *binarytree.Node

	data, n := readdata.LoadData()
	fmt.Println("data > ", data)

	// create a binary tree
	for i := 0; i < n; i++ {
		// allocate object
		x := binarytree.Node{
			Key:    data[i],
			Parent: nil,
			Left:   nil,
			Right:  nil,
		}

		// tree insert
		tree = binarytree.TreeInsert(tree, &x)
		fmt.Printf("tree walk >")

		// tree walk
		binarytree.InorderTreeWalk(tree)
		fmt.Printf("\n")

		// find tree maximum and minimum
		y := binarytree.TreeMaximum(tree)
		fmt.Println("maximum > ", y.Key)
		y = binarytree.TreeMinimum(tree)
		fmt.Println("minimum > ", y.Key)

		fmt.Printf("\n")
	}

}
