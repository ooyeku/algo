// Copyright 2024 olayeku
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package structs

import "sync"

const (
	RED   = true
	BLACK = false
)

// RBNode represents a node in the Red-Black Tree.
type RBNode struct {
	Value  interface{}
	Color  bool
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
}

// RedBlackTree represents a thread-safe Red-Black Tree.
type RedBlackTree struct {
	mu   sync.Mutex
	root *RBNode
	less func(a, b interface{}) bool
}

// NewRedBlackTree creates a new Red-Black Tree.
func NewRedBlackTree(less func(a, b interface{}) bool) *RedBlackTree {
	return &RedBlackTree{
		less: less,
	}
}

// Insert inserts a new value into the Red-Black Tree.
func (rbt *RedBlackTree) Insert(value interface{}) {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()
	newNode := &RBNode{Value: value, Color: RED}
	if rbt.root == nil {
		rbt.root = newNode
	} else {
		rbt.insertNode(rbt.root, newNode)
	}
	rbt.fixInsert(newNode)
}

func (rbt *RedBlackTree) insertNode(root, node *RBNode) {
	if rbt.less(node.Value, root.Value) {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
		} else {
			rbt.insertNode(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
		} else {
			rbt.insertNode(root.Right, node)
		}
	}
}

func (rbt *RedBlackTree) fixInsert(node *RBNode) {
	for node != rbt.root && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == RED {
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					rbt.rotateLeft(node)
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				rbt.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == RED {
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					rbt.rotateRight(node)
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				rbt.rotateLeft(node.Parent.Parent)
			}
		}
	}
	rbt.root.Color = BLACK
}

func (rbt *RedBlackTree) rotateLeft(node *RBNode) {
	right := node.Right
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Parent = node.Parent
	if node.Parent == nil {
		rbt.root = right
	} else if node == node.Parent.Left {
		node.Parent.Left = right
	} else {
		node.Parent.Right = right
	}
	right.Left = node
	node.Parent = right
}

func (rbt *RedBlackTree) rotateRight(node *RBNode) {
	left := node.Left
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Parent = node.Parent
	if node.Parent == nil {
		rbt.root = left
	} else if node == node.Parent.Right {
		node.Parent.Right = left
	} else {
		node.Parent.Left = left
	}
	left.Right = node
	node.Parent = left
}

// Search searches for a value in the Red-Black Tree.
func (rbt *RedBlackTree) Search(value interface{}) bool {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()
	return rbt.searchNode(rbt.root, value)
}

func (rbt *RedBlackTree) searchNode(node *RBNode, value interface{}) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if rbt.less(value, node.Value) {
		return rbt.searchNode(node.Left, value)
	}
	return rbt.searchNode(node.Right, value)
}

// InOrder traverses the tree in order and applies the provided function to each node's value.
func (rbt *RedBlackTree) InOrder(f func(value interface{})) {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()
	rbt.inOrderTraverse(rbt.root, f)
}

func (rbt *RedBlackTree) inOrderTraverse(node *RBNode, f func(value interface{})) {
	if node != nil {
		rbt.inOrderTraverse(node.Left, f)
		f(node.Value)
		rbt.inOrderTraverse(node.Right, f)
	}
}
