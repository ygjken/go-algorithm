package binarytree

type RBNode struct {
	Key    int
	Color  uint8 //0:red, 1:black
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
}

func InitRBnil() *RBNode {
	var n *RBNode
	n.Key = -10000
	n.Color = 1

	return n
}

var RBnil = InitRBnil()

func RBTreeLeftRotate(T *RBNode, x *RBNode) *RBNode {
	var r *RBNode // tree root
	var y *RBNode // the right child of node x

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
