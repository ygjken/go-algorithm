package binarytree

import "fmt"

type Node struct {
	Key    int
	Parent *Node
	Left   *Node
	Right  *Node
}

func TreeInsert(T *Node, x *Node) *Node {
	var r, y, z *Node

	r = T   // root
	y = nil // before attention variable
	z = r   // attention variable

	for z != nil {
		y = z
		if x.Key < z.Key {
			z = z.Left
		} else {
			z = z.Right
		}
	}
	x.Parent = y
	if y == nil {
		r = x
	} else {
		if x.Key < y.Key {
			y.Left = x
		} else {
			y.Right = x
		}
	}

	return r
}

func InorderTreeWalk(x *Node) {
	if x != nil {
		InorderTreeWalk(x.Left)
		fmt.Printf(" %d", x.Key)
		InorderTreeWalk(x.Right)
	}
}

func TreeMaximum(x *Node) *Node {
	if x != nil {
		for x.Right != nil {
			x = x.Right
		}
	}
	return x
}

func TreeMinimum(x *Node) *Node {
	if x != nil {
		for x.Left != nil {
			x = x.Left
		}
	}
	return x
}

func TreeSearch(x *Node, a int) *Node {
	for x != nil && x.Key != a {
		if a < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

func TreeDelete(T *Node, z *Node) *Node {
	// the order of replacement
	// z <- y <- x
	var y *Node
	var x *Node
	var r *Node // root

	r = T // the node under search

	// specify node y
	if z.Left == nil || z.Right == nil {
		y = z
	} else {
		y = TreeMinimum(z.Right)
	}

	// specify node x
	if y.Left != nil {
		x = y.Left
	} else {
		x = y.Right
	}

	// when don't have node x
	if x != nil {
		x.Parent = y.Parent
	}

	// is node y root
	if y.Parent == nil {
		r = x
	} else {
		// is y on parent's left or right
		if y == y.Parent.Left {
			y.Parent.Left = x
		} else {
			y.Parent.Right = x
		}
	}

	if y != z {
		z.Key = y.Key
	}

	return r
}

func TreeHeight(x *Node) int {
	var h, h_tmp int

	h = 0
	if x != nil {
		h = TreeHeight(x.Left)
		h_tmp = TreeHeight(x.Right)
		if h_tmp > h {
			h = h_tmp
		}
		h++
	}

	return h
}
