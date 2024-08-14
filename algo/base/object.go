package base

// Object is an interface that represents a generic object. It provides methods for comparison,
// equality check, hashing, string representation, and getting the size of the object.
type Object interface {
	Compare(Object) int
	Equals(Object) bool
	Hash() int
	String() string
	Size() int
}
