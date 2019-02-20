package kmp

type TNode struct {
	next *TNode // next pointer used by TNode link list

	T       string // string for matching
	kmpNext []int  // the next array used by kmp algorithm
	cur     int    // the current index of char matched in S
	length  int    // the length of T

	matched int // the index of position of T in s
}

func NewTNode(pattern string) (pTN *TNode) {
	pTN = &TNode{
		T:       pattern,
		kmpNext: make([]int, len(pattern)),
		cur:     0,
		length:  len(pattern),
	}

	// init kmpNext by common trainNext function
	trainNext(pTN.T, pTN.kmpNext)
	return
}

// Init TNode
func (tn *TNode) Init() {
	tn.cur = 0
	tn.matched = -1
	tn.next = nil
}

// Match reports whether the string s contains any match of pattern T.
func (tn *TNode) Match(s string, pos int, length int) int {
	i := pos
	tn.cur = 0
	for i < length && tn.cur < tn.length {
		if tn.cur == -1 || s[i] == tn.T[tn.cur] {
			i++
			tn.cur++
		} else {
			tn.cur = tn.kmpNext[tn.cur]
		}
	}

	if tn.cur == tn.length {
		return i - tn.cur
	}

	return -1
}

// TryMatch try to match T at pos of s.
func (tn *TNode) TryMatch(s string, pos, length int) bool {
	for tn.cur < tn.length {
		if tn.cur == -1 || s[pos] == tn.T[tn.cur] {
			tn.cur++
			break
		} else {
			tn.cur = tn.kmpNext[tn.cur]
		}
	}

	if tn.cur == tn.length {
		tn.matched = pos - tn.cur + 1
		return true
	}

	return false
}

type KMPMatchMachine struct {
	// Strings for matching, it is a array of TNode pointer
	patterns4m []*TNode
	// Number of strings to matching, it is the length of patterns4m
	npatterns int
}

func NewKMPMatchMachine(patterns ...string) (pMM *KMPMatchMachine) {
	pMM = &KMPMatchMachine{
		npatterns: len(patterns),
	}
	pMM.patterns4m = make([]*TNode, 0, pMM.npatterns)

	// init patterns4m field
	for _, pattern := range patterns {
		pTN := NewTNode(pattern)
		pMM.patterns4m = append(pMM.patterns4m, pTN)
	}

	return
}

// readyNodes
func (kmm *KMPMatchMachine) readyNodes() (patternsList *TNode) {
	patternsList = nil
	for _, n := range kmm.patterns4m {
		n.Init()
		n.next = patternsList
		patternsList = n
	}

	return
}

// MatchAll reports whether the string s contains any match of patterns.
func (kmm *KMPMatchMachine) MatchAll(s string) (matched int) {
	matched = 0
	length := len(s)
	curlist := kmm.readyNodes()
	var nextlist *TNode

	for i := range s {
		node := curlist
		for node != nil {
			nextNode := node.next
			if !node.TryMatch(s, i, length) {
				node.next = nextlist
				nextlist = node
			} else {
				matched++
			}
			node = nextNode
		}

		// if s match all patterns , finish Match
		if nextlist == nil {
			break
		} else {
			curlist = nextlist
			nextlist = nil
		}
	}

	return
}
