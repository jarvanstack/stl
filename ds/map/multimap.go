package treemap

import (
	"stl/ds/rbtree"
	"stl/utils/comparator"
	"stl/utils/sync"
	"stl/utils/visitor"
)

// MultiMap uses RbTress for internal data structure, and keys can bee repeated.
type MultiMap struct {
	 //TODO: Complete me!
}

//NewMultiMap creates a new MultiMap
func NewMultiMap(opts ...Option) *MultiMap {
	 //TODO: Complete me!
}

//Insert inserts a key-value to the MultiMap
func (mm *MultiMap) Insert(key, value interface{}) {
	 //TODO: Complete me!
}

//Get returns the first node's value by the passed key if the key is in the MultiMap, otherwise returns nil
func (mm *MultiMap) Get(key interface{}) interface{} {
	 //TODO: Complete me!
}

//Erase erases the key in the MultiMap
func (mm *MultiMap) Erase(key interface{}) {
	 //TODO: Complete me!
}

//Find finds the node by the passed key in the MultiMap and returns its iterator
func (mm *MultiMap) Find(key interface{}) *MapIterator {
	 //TODO: Complete me!
}

//LowerBound find the first node that its key is equal or greater than the passed key in the MultiMap, and returns its iterator
func (mm *MultiMap) LowerBound(key interface{}) *MapIterator {
	 //TODO: Complete me!
}

//Begin returns the first node's iterator
func (mm *MultiMap) Begin() *MapIterator {
	 //TODO: Complete me!
}

//First returns the first node's iterator
func (mm *MultiMap) First() *MapIterator {
	 //TODO: Complete me!
}

//Last returns the last node's iterator
func (mm *MultiMap) Last() *MapIterator {
	 //TODO: Complete me!
}

//Clear clears the MultiMap
func (mm *MultiMap) Clear() {
	 //TODO: Complete me!
}

// Contains returns true if the passed value is in the MultiMap. otherwise returns false.
func (mm *MultiMap) Contains(value interface{}) bool {
	 //TODO: Complete me!
}

// Size returns the amount of elements in the MultiMap
func (mm *MultiMap) Size() int {
	 //TODO: Complete me!
}

// Traversal traversals elements in the MultiMap, it will not stop until to the end of the MultiMap or the visitor returns false
func (mm *MultiMap) Traversal(visitor visitor.KvVisitor) {
	 //TODO: Complete me!
}
