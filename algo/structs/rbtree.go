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

// RED is a constant of type bool representing the color red.
// BLACK is a constant of type bool representing the color black.
const (
	RED   = true
	BLACK = false
)

// RBNode represents a node in a Red-Black Tree.
// It contains the value, color, left child, right child, and parent of the node.
type RBNode struct {
	Value  interface{}
	Color  bool
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
}

// RedBlackTree represents a red-black tree data structure.
//
// Red-black trees are self-balancing binary search trees that provide
// efficient insertion, deletion, and search operations. Each node of
// the tree is colored either red or black which ensures that the tree
// remains balanced in terms of the number of nodes on each side of each
// node.
//
// The RedBlackTree type consists of a mutex to handle concurrent access,
// a root node representing the top of the tree, and a "less" function
// used to compare the values of the nodes. The "less" function should
// return true if the first value is less than the second value.
type RedBlackTree struct {
	mu   sync.Mutex
	root *RBNode
	less func(a, b interface{}) bool
}

// NewRedBlackTree creates a new Red-Black Tree with the specified comparison function.
// The comparison function takes two interface{} values, a and b, and returns true if a is less than b.
// The Red-Black Tree is returned as a pointer to a RedBlackTree struct.
func NewRedBlackTree(less func(a, b interface{}) bool) *RedBlackTree {
	return &RedBlackTree{
		less: less,
	}
}

// Insert adds a new node with the specified value to the red-black tree. It
// acquires a lock to ensure thread safety during the insertion operation. If
// the tree is empty, the new node becomes the root. Otherwise, the method
// delegates the insertion to the insertNode function to find the appropriate
// position for the new node. After the insertion, the fixInsert function is
// called to restore the red-black tree properties.
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

// insertNode inserts a new node into the RedBlackTree.
// It recursively inserts the node in the appropriate position based on its value.
// If the value of the new node is less than the value of the current root,
// it goes to the left sub-tree. Otherwise, it goes to the right sub-tree.
// If the left/right child is nil, assign the new node to it and set its parent to the root.
// Otherwise, recursively call insertNode on the respective child node.
// The method is called internally by the Insert method.
// Parameters:
// - root: The root node of the sub-tree where the new node will be inserted.
// - node: The new node to be inserted into the sub-tree.
// Complexity: O(log n), where n is the number of nodes in the tree.
// The method does not return anything.
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

// fixInsert fixes the Red-Black Tree after inserting a node by ensuring that all
// the properties of a Red-Black Tree are maintained. It starts from the newly
// inserted node and traverses up the tree, performing rotations and color changes
// as necessary. The fixInsert method is called internally by the Insert method.
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

// rotateLeft performs a left rotation on the specified node in a Red-Black Tree.
// The right child of the node becomes its parent, and the left child of the right child becomes the right child of the original node.
// The node's parent and the original right child's parent are adjusted accordingly.
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

// rotateRight performs a right rotation on the Red-Black Tree with the specified node as the pivot.
// The left child of the input node becomes the new parent and the input node becomes the right child of the new parent.
// The right child of the new parent becomes the left child of the input node, and the input node becomes the parent of the right child.
// If the node has a parent, the new parent will replace the node as its parent's left or right child accordingly.
// If the node is the root of the tree, the new parent becomes the root.
// Any child or parent references affected during the rotation are adjusted accordingly.
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

// Search searches for a value in the Red-Black Tree and returns true if the value is found, false otherwise. The search
// operation is performed by calling the `searchNode` method recursively on the root node of the Red-Black Tree.
// The method acquires a lock and releases it before returning.
// The time complexity of the search operation is O(log N), where N is the number of nodes in the tree.
func (rbt *RedBlackTree) Search(value interface{}) bool {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()
	return rbt.searchNode(rbt.root, value)
}

// searchNode searches for a node with the given value in the Red-Black Tree.
// It starts the search from the provided node and recursively traverses the tree,
// comparing the values and moving left or right based on the comparison.
// Returns true if the node with the given value is found, false otherwise.
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

// InOrder traverses the RedBlackTree in an in-order manner and applies the provided function f to each node's value.
// The traversal starts from the root node and proceeds recursively to the left subtree, then the current node,
// and finally the right subtree. The function f is called on each node's value during traversal.
// The RedBlackTree's mutex is acquired before traversal and released after traversal is completed.
// This method can be used to perform an in-order traversal and apply a function to each value in the tree.
func (rbt *RedBlackTree) InOrder(f func(value interface{})) {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()
	rbt.inOrderTraverse(rbt.root, f)
}

// inOrderTraverse performs an in-order traversal of the Red-Black Tree starting from the given node.
// It applies the provided function to each value in the tree in ascending order.
func (rbt *RedBlackTree) inOrderTraverse(node *RBNode, f func(value interface{})) {
	if node != nil {
		rbt.inOrderTraverse(node.Left, f)
		f(node.Value)
		rbt.inOrderTraverse(node.Right, f)
	}
}
