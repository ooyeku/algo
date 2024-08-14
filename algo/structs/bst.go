package structs

import "sync"

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

// BinarySearchTree represents a binary search tree data structure. It has a mutex for thread safety
// and a root node. The less function is used to compare values for determining the position of nodes.
// The type also includes methods for inserting values, searching for values, and performing an
// in-order traversal of the tree.
type BinarySearchTree struct {
	mu   sync.Mutex
	root *TreeNode
	less func(a, b interface{}) bool
}

// NewBinarySearchTree creates a new instance of BinarySearchTree and returns it.
// The BinarySearchTree is initialized with the provided less function, which is used to determine the ordering of values in the tree.
// The less function should take two arguments of type interface{} and return a boolean indicating whether the first argument is less than the second argument.
// The BinarySearchTree struct is thread-safe and supports concurrent operations.
// To insert a value into the tree, use the Insert method.
// To search for a value in the tree, use the Search method.
// To traverse the tree in inorder, use the InOrder method.
// The tree is implemented using the TreeNode struct, which consists of a value, a left child, and a right child.
// The BinarySearchTree methods manipulate the tree by creating, inserting, searching, and traversing the TreeNode instances.
func NewBinarySearchTree(less func(a, b interface{}) bool) *BinarySearchTree {
	return &BinarySearchTree{
		less: less,
	}
}

// Insert appends a new node with the given value into the BinarySearchTree.
//
// The method locks the BinarySearchTree's mutex for exclusive access,
// defers its unlocking before returning.
//
// The method delegates the insertion operation to the insertNode helper method.
// The helper method performs a recursive search in the BinarySearchTree
// to find the appropriate position for the new node.
// If the current node is nil, the helper method creates a new TreeNode with the given value
// and returns it. If the given value is less than the current node's value,
// the helper method recursively inserts the value into the left subtree of the current node.
// Otherwise, it inserts the value into the right subtree. The helper method
// updates the left or right pointer of the current node accordingly and returns the current node.
//
// The method does not return any value. It updates the BinarySearchTree's root by assigning
// it the result of the insertNode operation.
func (bst *BinarySearchTree) Insert(value interface{}) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	bst.root = bst.insertNode(bst.root, value)
}

// insertNode inserts a new node with the given value into the binary search tree.
// If the tree is empty, it creates a new tree node with the given value as the root.
// Otherwise, it recursively inserts the new node to the left subtree if the value is less than the current node's value,
// or to the right subtree if the value is greater than or equal to the current node's value.
// The method returns the modified tree with the new node inserted.
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

// Search searches for the given value in the binary search tree. It acquires a lock to ensure
// thread safety, and then calls the private searchNode helper method to recursively search for
// the value starting from the root of the tree. If the value is found, it returns true, otherwise
// it returns false.
//
// This method blocks other goroutines from accessing or modifying the tree until the search operation
// is completed, to ensure that the tree remains consistent during the search.
//
// bst       - The BinarySearchTree instance to search in.
// value     - The value to search for in the tree.
//
// Returns true if the value is found in the tree, false otherwise.
func (bst *BinarySearchTree) Search(value interface{}) bool {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return bst.searchNode(bst.root, value)
}

// searchNode searches for a specific value in the Binary Search Tree.
// It recursively traverses the tree to find the node with the given value.
// It returns true if a node with the value is found, otherwise it returns false.
// The search starts at the given node, recursively following the left or right child
// depending on whether the value is less than or greater than the current node's value.
// If the given node is nil, the search returns false.
// The less method defined on the BinarySearchTree is used to compare values.
// The search is performed in a concurrent-safe manner using a mutex.
// Refer to the BinarySearchTree's Insert method for an example of how to use this method.
// The value parameter represents the value to be searched for in the tree.
// The method returns a boolean value indicating whether the value was found or not.
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

// InOrder performs an in-order traversal of the binary search tree and applies the given function
// to each value in ascending order. It locks the mutex before traversing and unlocks it afterwards.
func (bst *BinarySearchTree) InOrder(f func(value interface{})) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	bst.inOrderTraverse(bst.root, f)
}

// inOrderTraverse performs an in-order traversal of the binary search tree.
// It recursively traverses the tree in the ascending order of values and applies the given function to each value.
// The traversal starts at the specified node and follows the left child, then visits the current node,
// and finally follows the right child.
// The method does not return any value, but instead calls the provided function on each visited node's value.
// The method is used internally by the InOrder method.
// The node parameter represents the current node being visited during the traversal.
// The f parameter represents the function to apply to each visited node's value.
func (bst *BinarySearchTree) inOrderTraverse(node *TreeNode, f func(value interface{})) {
	if node != nil {
		bst.inOrderTraverse(node.Left, f)
		f(node.Value)
		bst.inOrderTraverse(node.Right, f)
	}
}
