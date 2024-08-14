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

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

// BinarySearchTree represents a thread-safe binary search tree.
type BinarySearchTree struct {
	mu   sync.Mutex
	root *TreeNode
	less func(a, b interface{}) bool
}

// NewBinarySearchTree creates a new binary search tree.
func NewBinarySearchTree(less func(a, b interface{}) bool) *BinarySearchTree {
	return &BinarySearchTree{
		less: less,
	}
}

// Insert inserts a new value into the binary search tree.
func (bst *BinarySearchTree) Insert(value interface{}) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	bst.root = bst.insertNode(bst.root, value)
}

func (bst *BinarySearchTree) insertNode(node *TreeNode, value interface{}) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if bst.less(value, node.Value) {
		node.Left = bst.insertNode(node.Left, value)
	} else {
		node.Right = bst.insertNode(node.Right, value)
	}
	return node
}

// Search searches for a value in the binary search tree.
func (bst *BinarySearchTree) Search(value interface{}) bool {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return bst.searchNode(bst.root, value)
}

func (bst *BinarySearchTree) searchNode(node *TreeNode, value interface{}) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if bst.less(value, node.Value) {
		return bst.searchNode(node.Left, value)
	}
	return bst.searchNode(node.Right, value)
}

// InOrder traverses the tree in order and applies the provided function to each node's value.
func (bst *BinarySearchTree) InOrder(f func(value interface{})) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	bst.inOrderTraverse(bst.root, f)
}

func (bst *BinarySearchTree) inOrderTraverse(node *TreeNode, f func(value interface{})) {
	if node != nil {
		bst.inOrderTraverse(node.Left, f)
		f(node.Value)
		bst.inOrderTraverse(node.Right, f)
	}
}
