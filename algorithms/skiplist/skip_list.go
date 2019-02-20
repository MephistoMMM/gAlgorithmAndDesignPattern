// Copyright © 2019 Mephis Pheies <mephistommm@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package skiplist

import (
	"math/rand"
)

// SkipList is interface of skip list
type SkipList interface {
	Insert(key int, value interface{})
	Delete(key int)
	Search(key int) interface{}
}

type node struct {
	key   int
	value interface{}

	// list to hold references to node of different level
	forward []*node
}

// newNode create and init a new Node
func newNode(level int, key int, value interface{}) *node {
	node := &node{
		key:     key,
		value:   value,
		forward: make([]*node, level+1),
	}

	return node
}

// skipList struct of skip list
type skipList struct {
	// Maximum level for this skip list
	maxLevel int

	// p is the fraction of the nodes with level
	// i references also having level i+1 references
	p float64

	header *node
	// current level of skip list
	level int
}

func NewSkipList(maxLevel int, p float64) SkipList {
	return &skipList{
		maxLevel: maxLevel,
		p:        p,
		header:   newNode(maxLevel, -1, nil),
		level:    0,
	}
}

// randomLevel create random level for node
func (sl *skipList) randomLevel() int {
	level := 0
	for rand.Float64() < sl.p && level < sl.maxLevel {
		level++
	}

	return level
}

// skipToNode go to the correct position of the node with key
func (sl *skipList) skipToNode(key int) (current *node, update []*node) {
	update = make([]*node, 0, sl.level+1)
	current = sl.header

	// start from highest level of skip list
	// move the current reference forward while key
	// is greater than key of node next to current
	// Otherwise inserted current in update and
	// move one level down and continue search
	for i := sl.level; i > -1; i-- {
		for current.forward[i] != nil &&
			current.forward[i].key < key {
			current = current.forward[i]
		}

		update = append(update, current)
	}

	// reached level 0 and forward reference to
	// right, which is desired position to
	// insert key.
	current = current.forward[0]
	return
}

// Insert insert given key and value in skip list
func (sl *skipList) Insert(key int, value interface{}) {
	current, update := sl.skipToNode(key)

	// if current is nil that means we have reached
	// to end of the level or current's key is not equal
	// to key to insert that means we have to insert
	// node between update[0] and current node
	if current == nil || current.key != key {
		// Generate a random level for node
		rlevel := sl.randomLevel()

		// If random level is greater than list's current
		// level (node with highest level inserted in
		// list so far), initialize update value with reference
		// to header for further use
		if rlevel > sl.level {
			updat := make([]*node, rlevel-sl.level)
			for i := range updat {
				updat[i] = sl.header
			}
			update = append(updat, update...)

			// TODO config sl.header.forward length dynamically
			sl.level = rlevel
		}

		// create new node with random level generated
		n := newNode(rlevel, key, value)

		// insert node by rearranging references
		for i := 0; i < rlevel+1; i++ {
			n.forward[i] = update[sl.level-i].forward[i]
			update[sl.level-i].forward[i] = n
		}
	}
}

// Delete ...
func (sl *skipList) Delete(key int) {
	current, update := sl.skipToNode(key)

	// If current node is target node
	if current != nil && current.key == key {

		for i := 0; i < sl.level+1; i++ {

			// If at level i, next node is not target
			// node, break the loop, no need to move
			// further level
			if update[sl.level-i].forward[i] != current {
				break
			}
			update[sl.level-i].forward[i] = current.forward[i]
		}

		// Remove levels having no elements
		for sl.level > 0 && sl.header.forward[sl.level] == nil {
			sl.level--
		}
	}
}

// Search ...
func (sl *skipList) Search(key int) interface{} {
	current := sl.header

	// start from highest level of skip list
	// move the current reference forward while key
	// is greater than key of node next to current
	// Otherwise inserted current in update and
	// move one level down and continue search
	for i := sl.level; i > -1; i-- {
		for current.forward[i] != nil &&
			current.forward[i].key < key {
			current = current.forward[i]
		}

	}

	// reached level 0 and forward reference to
	// right, which is desired position to
	// insert key.
	current = current.forward[0]

	if current != nil && current.key == key {
		return current.value
	}

	return nil
}
