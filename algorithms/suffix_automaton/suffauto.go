// Package sufauto provide the algorithm about suffix automaton
//
// Author: Mephis Pheies <mephistommm@gmail.com>
package sufauto

// size of char set
const charsetSize = 26

// Node represent state in SAM
type Node struct {
	// translate arc
	arc      [charsetSize]*Node
	sufflink *Node
	max      int
	// posCnt represents the size of its endpos
	posCnt int
}

// Sufflink suffix link of node
func (n *Node) Sufflink() *Node {
	return n.sufflink
}

// Max the length of the longest string in state
func (n *Node) Max() int {
	return n.max
}

// Min the length of the shortest string in state
func (n *Node) Min() int {
	return n.sufflink.max + 1
}

// PosCnt the size of endpos(n)
func (n *Node) PosCnt() int {
	return n.posCnt
}

// SuffixAutomaton SAM
type SuffixAutomaton struct {
	start *Node
	last  *Node
}

// Init SAM
func (sam *SuffixAutomaton) Init() {
	sam.start = &Node{}
	sam.last = sam.start
}

// StartNode get start node
func (sam *SuffixAutomaton) StartNode() *Node {
	return sam.start
}

// LastNode get last node
func (sam *SuffixAutomaton) LastNode() *Node {
	return sam.last
}

// Extend add a new char to extend the sam
func (sam *SuffixAutomaton) Extend(c int) *Node {
	// node u represent the new whole string
	u, v := &Node{max: sam.last.max + 1}, sam.last

	// let u be accessed by node who has no 'c' arc in the path of suffix link from last node
	// notice that v will reach nil
	for ; v != nil && v.arc[c] == nil; v = v.sufflink {
		v.arc[c] = u
	}

	// if v reach nil, let the sufflink of v point to start(gen of prefix tree)
	if v == nil {
		u.sufflink = sam.start
	} else if v.arc[c].max == v.max+1 {
		// let the sufflink of u point to the node accessed by v.arc[c] directly
		u.sufflink = v.arc[c]
	} else {
		// split node, n is new node and o is the old one
		n, o := &Node{max: v.max + 1}, v.arc[c]
		// copy output arc to new node
		n.arc = o.arc
		// let the sufflink of new node point to the sufflink of origin node
		n.sufflink = o.sufflink
		// let the sufflink of old node and u point to new node n
		u.sufflink = n
		o.sufflink = u.sufflink

		// let n be pointed to by nodes who has arc to old node in the path of suffix link
		for ; v != nil && v.arc[c] == nil; v = v.sufflink {
			v.arc[c] = n
		}
	}

	// update the last node
	sam.last = u
	return u
}

// SAWithTopo SAM
type SAWithTopo struct {
	start *Node
	last  *Node

	// To enumerate all nodes conveniently, put node into pool
	// curr points to behind of the current last node
	curr    int
	pool    []Node
	strSize int

	// store the result of topo sort (accroding to 'max' from small to large)
	topo []*Node
}

// Init SAM
func (sam *SAWithTopo) Init(strSize int) {
	sam.pool = make([]Node, 2*sam.strSize+1)
	sam.topo = nil
	sam.strSize = strSize
	sam.curr = 0
	sam.start = sam.newNode(sam.curr, 0, false)
	sam.last = sam.start
	sam.curr++
}

// newNode take out and init a node from pool
func (sam *SAWithTopo) newNode(pos int, max int, newSuffix bool) *Node {
	n := &sam.pool[pos]
	n.max = max
	if newSuffix {
		n.posCnt = 1
	}

	return n
}

// StartNode get start node
func (sam *SAWithTopo) StartNode() *Node {
	return sam.start
}

// LastNode get last node
func (sam *SAWithTopo) LastNode() *Node {
	return sam.last
}

// Extend add a new char to extend the sam
func (sam *SAWithTopo) Extend(c int) *Node {
	sam.topo = nil

	// node u represent the new whole string
	u, v := sam.newNode(sam.curr, sam.last.max+1, true), sam.last
	sam.curr++

	// notice that v will reach nil
	for ; v != nil && v.arc[c] == nil; v = v.sufflink {
		v.arc[c] = u
	}

	if v == nil {
		u.sufflink = sam.start
	} else if v.arc[c].max == v.max+1 {
		u.sufflink = v.arc[c]
	} else {
		// split node, n is new node and o is the old one
		n, o := sam.newNode(sam.curr, sam.last.max+1, false), v.arc[c]
		sam.curr++
		n.arc = o.arc
		n.sufflink = o.sufflink
		u.sufflink = n
		o.sufflink = u.sufflink

		for ; v != nil && v.arc[c] == nil; v = v.sufflink {
			v.arc[c] = n
		}
	}

	sam.last = u
	return u
}

// Toposort sort sam by topo sort
func (sam *SAWithTopo) Toposort() []*Node {
	if sam.topo != nil {
		return sam.topo
	}

	buc := make([]int, sam.strSize)

	// keep the max value for cleaning buc
	max := 0

	// normal counting sort
	for i := range sam.pool[:sam.curr] {
		v := &sam.pool[i]
		if max < v.max {
			max = v.max
		}
		buc[v.max]++
	}
	for i := 1; i <= max; i++ {
		buc[i] += buc[i-1]
	}

	if len(sam.topo) != sam.curr {
		sam.topo = make([]*Node, sam.curr)
	}
	for i := range sam.pool[:sam.curr] {
		v := &sam.pool[i]
		buc[v.max]--
		sam.topo[buc[v.max]] = v
	}

	return sam.topo
}

// Calc calculate the endpos number of nodes
func (sam *SAWithTopo) Calc() {
	topo := sam.Toposort()

	// Recursive accroding to the order of topo from children to parents
	for i := len(topo) - 1; i > 0; i-- {
		v := topo[i]
		v.sufflink.posCnt += v.posCnt
	}
}
