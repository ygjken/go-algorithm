package rbtree

import (
	"fmt"
)

type Node struct {
	Key    int
	Color  uint8 //0:red, 1:black
	Parent *Node
	Left   *Node
	Right  *Node
}

var RBnil = &Node{
	Key:    -10000,
	Color:  1,
	Parent: nil,
	Right:  nil,
	Left:   nil,
}

func LeftRotate(T *Node, x *Node) *Node {
	var r *Node // tree root
	var y *Node // the right child of node x

	r = T
	y = x.Right

	// Change members of y.Left and x.Right
	x.Right = y.Left
	if y.Left != RBnil {
		y.Left.Parent = x
	}

	// Change parent of y
	y.Parent = x.Parent

	// Change members of x.Parent
	if x.Parent == RBnil {
		r = y
	} else {
		if x == x.Parent.Left {
			x.Parent.Left = y
		} else {
			x.Parent.Right = y
		}
	}
	// changed children of node x
	// changed right child of node y
	// changed parent of y

	y.Left = x
	x.Parent = y

	return r
}

func RightRotate(T *Node, x *Node) *Node {
	var r *Node
	var y *Node

	r = T
	y = x.Left

	x.Left = y.Right
	if y.Right != RBnil {
		y.Right.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == RBnil {
		r = y
	} else {
		if x == x.Parent.Right {
			x.Parent.Right = y
		} else {
			x.Parent.Left = y
		}
	}

	y.Right = x
	x.Parent = y

	return r
}

func Walk(x *Node) {
	if x != RBnil {
		Walk(x.Left)
		fmt.Printf(" %d", x.Key)
		Walk(x.Right)
	}
}

func TreeInsert(T *Node, x *Node) *Node {
	var r, y, z *Node

	r = T     // root
	y = RBnil // before attention variable
	z = r     // attention variable

	for z != RBnil {
		y = z
		if x.Key < z.Key {
			z = z.Left
		} else {
			z = z.Right
		}
	}

	x.Parent = y

	if y == RBnil {
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

func RBTrerInsert(T *Node, x *Node) *Node {
	var r *Node //root
	var y *Node

	r = TreeInsert(T, x)

	// agjust tree
	// node x color is 0(red)
	for x != r && x.Parent.Color == 0 {

		if x.Parent == x.Parent.Parent.Left {
			y = x.Parent.Parent.Right
			if y.Color == 0 {
				x.Parent.Color = 1
				y.Color = 1
				x.Parent.Parent.Color = 0
				x = x.Parent.Parent // switch node we forcus
			} else {
				// left-right -> double rotation
				if x == x.Parent.Right {
					x = x.Parent
					r = LeftRotate(r, x)
				}
				// left-left -> single rotation?
				x.Parent.Color = 1
				x.Parent.Parent.Color = 0
				r = RightRotate(r, x.Parent.Parent)
			}
		} else {
			y = x.Parent.Parent.Left
			if y.Color == 0 {
				x.Parent.Color = 1
				y.Color = 1
				x.Parent.Parent.Color = 0
				x = x.Parent.Parent // switch node we forcus
			} else {
				// right-left -> double rotation
				if x == x.Parent.Left { // right-left
					x = x.Parent
					r = RightRotate(r, x)
				}
				// right-right -> sihle rotation?
				x.Parent.Color = 1
				x.Parent.Parent.Color = 0
				r = LeftRotate(r, x.Parent.Parent)

			}
		}
	}
	r.Color = 1

	return r
}
