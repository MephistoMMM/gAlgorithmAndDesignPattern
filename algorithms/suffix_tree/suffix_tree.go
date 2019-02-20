// Copyright © 2018 Mephis Pheies <mephistommm@gmail.com>
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
// Package suftree is a golang program to implement Ukkonen's Suffix Tree Construction
package suftree

import "fmt"

type SuffixTreeNode struct {
	children map[byte]*SuffixTreeNode
	// pointer to other node via suffix link
	suffixLink *SuffixTreeNode

	// (start, end) interval specifies the edge, by which the
	// node is connected to its parent node. Each edge will
	// connect two nodes, one parent and one child, and
	// (start, end) interval of a given edge will be stored
	// in the child node. Lets say there are two nods A and B
	// connected by an edage with indices (5, 8) then this
	// indices (5, 8) will be stored in node B.
	start int
	end   *int

	// for leaf nodes, it stores the index of suffix for
	// the path from root to leaf
	suffixIndex int
}

type Node = SuffixTreeNode

func newNode(root *Node, start int, end *int) *Node {
	node := &Node{
		// For root node, suffixLink will be set to nil
		// For internal nodes, suffixLink will be set to root
		// by default in current extension and may change in
		// next extension
		suffixLink: root,
		start:      start,
		end:        end,

		// suffixIndex will be set to -1 by default and
		// actual suffix index will be set later for leaves
		// at the end of all phases
		suffixIndex: -1,

		children: make(map[byte]*Node, 16),
	}

	return node
}

func (node *Node) EdgeLength() int {
	return *node.end - node.start + 1
}

type SuffixTree struct {
	// Input string
	text string
	// Pointer to root node
	root *Node

	// lastNewNode will point to newly created internal node,
	// waiting for it's suffix link to be set, which might get
	// a new suffix link (other than root) in next extension of
	// same phase. lastNewNode will be set to nil when last
	// newly created internal node (if there is any) got it's
	// suffix link reset to new internal node created in next
	// extension of same phase.
	lastNewNode *Node
	activeNode  *Node

	// activeEdge is represented as input string character
	// index (not the character itself)
	activeEdge   int // -1
	activeLength int // 0

	// remainingSuffixCount tells how many suffixes yet to
	// be added in tree
	remainingSuffixCount int // 0
	leafEnd              int // -1

	// Length of input string
	size int // -1

	rootEnd *int
}

type Tree = SuffixTree

func (tree *Tree) walkDown(currNode *Node) bool {
	// activePoint change for walk down (APCFWD) using
	// Skip/Count Trick (Trick 1). If activeLength is greater
	// than current edge length, set next internal node as
	// activeNode and adjust activeEdge and activeLength
	// accordingly to represent same activePoint
	edgeLength := currNode.EdgeLength()
	if tree.activeLength >= edgeLength {
		tree.activeEdge += edgeLength
		tree.activeLength -= edgeLength
		tree.activeNode = currNode
		return true
	}
	return false
}

func (tree *Tree) extend(pos int) {
	// Extension Rule 1, this takes care of extending all
	// leaves created so far in tree
	tree.leafEnd = pos

	// Increment remainingSuffixCount indicating that a
	// new suffix added to the list of suffixes yet to be
	// added in tree
	tree.remainingSuffixCount++

	// set lastNewNode to nil while starting a new phase,
	// indicating there is no internal node wating for
	// it's suffix link reset in current phase
	tree.lastNewNode = nil

	// Add all suffixes (yet to be added) one by one in tree
	for tree.remainingSuffixCount > 0 {
		if tree.activeLength == 0 {
			// APCFALZ
			tree.activeEdge = pos
		}

		char := tree.text[pos]
		edgeStartChar := tree.text[tree.activeEdge]

		// There is no outgoing edge starting with
		// activeEdge from activeNode
		if tree.activeNode.children[edgeStartChar] == nil {
			// Extension Rule 2 (A new leaf edge gets created)
			tree.activeNode.children[edgeStartChar] = newNode(tree.root, pos, &tree.leafEnd)

			// A new leaf edge is created in above line starting
			// from an existing node (the current activeNode), and
			// if there is any internal node waiting for it's suffix
			// link get reset, point the suffix link from that last
			// internal node to current activeNode. Then set lastNewNode
			// to nil indicating no more node wating for suffix link
			// reset
			if tree.lastNewNode != nil {
				tree.lastNewNode.suffixLink = tree.activeNode
				tree.lastNewNode = nil
			}
		} else {
			// There is an outgoing edge staring with activeEdge
			// from activeNode

			// Get the next node at the end of edge starting
			// with activeEdge
			next := tree.activeNode.children[tree.text[tree.activeEdge]]
			if tree.walkDown(next) {
				// Start from next node (the new activeNode)
				continue
			}

			existChar := tree.text[next.start+tree.activeLength]
			// Extension Rule 3 (current character being processed
			// is already on the edge)
			if existChar == char {
				// If a newly created node waiting for it's
				// suffix link to be set, then set suffix link
				// of that waiting node to current active node
				if tree.lastNewNode != nil && tree.activeNode != tree.root {
					tree.lastNewNode.suffixLink = tree.activeNode
					tree.lastNewNode = nil
				}

				// APCFER3
				tree.activeLength++
				// STOP all further processing in this phase
				// and move on to new phase
				break
			}

			// We will be here when activePoint is in middle of
			// the edge being traversed and current character
			// being processed is not on the edge (we fall off
			// the tree). In this case, we add a new internal node
			// and a new leaf edge going out of that new node. This
			// is Extension Rule 2, where a new leaf edge and a new
			// internal node get created
			splitEnd := next.start + tree.activeLength - 1

			// New internal node
			split := newNode(tree.root, next.start, &splitEnd)
			tree.activeNode.children[edgeStartChar] = split

			// New leaf coming out of new internal node
			split.children[char] = newNode(tree.root, pos, &tree.leafEnd)
			next.start += tree.activeLength
			split.children[existChar] = next

			// We got a new internal node here. If there is any
			// internal node created in last extensions of same
			// phase which is still waiting for it's suffix link
			// reset, do it now.
			if tree.lastNewNode != nil {
				// suffixLink of lastNewNode points to current newly
				// created internal node
				tree.lastNewNode.suffixLink = split
			}

			// Make the current newly created internal node waiting
			// for it's suffix link reset (which is pointing to root
			// at present). If we come across any other internal node
			// (existing or newly created) in next extension of same
			// phase, when a new leaf edge gets added (i.e. when
			// Extension Rule 2 applies is any of the next extension
			// of same phase) at that point, suffixLink of this node
			// will point to that internal node.
			tree.lastNewNode = split
		}

		// One suffix got added in three, decrement the count of
		// suffixes yet to be added.
		tree.remainingSuffixCount--
		if tree.activeNode == tree.root && tree.activeLength > 0 {
			tree.activeLength--
			tree.activeEdge = pos - tree.remainingSuffixCount + 1
		} else if tree.activeNode != tree.root {
			// APCFER2C2
			tree.activeNode = tree.activeNode.suffixLink
		}
	}
}

func (tree *Tree) Edge(i, j int) string {
	return tree.text[i : j+1]
}

func (tree *Tree) PrintSuffixIndex() {
	tree.PrintSuffixIndexByDFS(tree.root)
}

// PrintSuffixIndexByDFS print the suffix in DFS manner
// Each edge along with it's suffix index will be printed
func (tree *Tree) PrintSuffixIndexByDFS(n *Node) {
	if n == nil {
		return
	}

	if n.start != -1 {
		fmt.Print(tree.Edge(n.start, *n.end))
	}

	leaf := true
	for _, v := range n.children {
		if leaf && n.start != -1 {
			fmt.Printf(" [%d]\n", n.suffixIndex)
		}

		// Current node is not a leaf as it has outgoing
		// edges from it
		leaf = false
		tree.PrintSuffixIndexByDFS(v)
	}

	if leaf {
		// Print suffixIndex
		fmt.Printf(" [%d]\n", n.suffixIndex)
	}
}

// SetSuffixIndexByDFS set suffix index
func (tree *Tree) setSuffixIndexByDFS(n *Node, labelHeight int) {
	if n == nil {
		return
	}

	leaf := 1
	for _, v := range n.children {
		// Current node is not a leaf as it has outgoing
		// edges from it
		leaf = 0
		tree.setSuffixIndexByDFS(v, labelHeight+v.EdgeLength())
	}

	if leaf == 1 {
		n.suffixIndex = len(tree.text) - labelHeight
	}
}

func (tree *Tree) FreeSuffixTreeByPostOrder(n *Node) {
	if n == nil {
		return
	}

	for k, v := range n.children {
		tree.FreeSuffixTreeByPostOrder(v)
		delete(n.children, k)
	}

	if n.suffixIndex == -1 {
		n.end = nil
	}
	n = nil
}

// BuildSuffixTree build the suffix tree and print the edge labels along with
// suffixIndex. suffixIndex for leaf edges will be >= 0 and
// for non-leaf edges will be -1
func BuildSuffixTree(text string) *Tree {
	tree := &Tree{
		text:                 text,
		rootEnd:              new(int),
		activeEdge:           -1,
		activeLength:         0,
		remainingSuffixCount: 0,
		leafEnd:              -1,
	}
	*tree.rootEnd = -1

	// Root is a special node with start and end indices as -1,
	// as it has no parent from where and edge comes to root
	tree.root = newNode(nil, -1, tree.rootEnd)

	tree.activeNode = tree.root
	for i := 0; i < len(text); i++ {
		tree.extend(i)
	}
	tree.setSuffixIndexByDFS(tree.root, 0)

	return tree
}
