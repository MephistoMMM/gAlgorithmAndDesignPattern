// Package kmp provide kmp algorithm
//
// Author: Mephis Pheies <mephistommm@gmail.com>
package kmp

func trainNext(T string, next []int) {
	i, lenOfT := 0, len(T)-1
	next[i] = -1
	j := -1
	for i < lenOfT {
		if j == -1 || T[i] == T[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
}

// Index match T in S by normal kmp algorithm
func Index(S, T string, pos int) int {
	lenOfS, lenOfT := len(S), len(T)

	next := make([]int, lenOfT)
	trainNext(T, next)

	i := pos
	j := 0
	for i < lenOfS && j < lenOfT {
		if j == -1 || S[i] == T[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == lenOfT {
		return i - j
	}

	return -1
}

type lazyNext struct {
	i    int
	j    int
	T    string
	next []int
}

func (l *lazyNext) Init(T string) {
	l.T = T
	l.next = make([]int, len(T))
	l.next[0] = -1
	l.i = 0
	l.j = -1
}

func (l *lazyNext) Next(k int) int {
	// k should be between 0 and lenOfT-1
	if k <= l.i {
		return l.next[k]
	}

	for l.i < k {
		if l.j == -1 || l.T[l.i] == l.T[l.j] {
			l.i++
			l.j++
			l.next[l.i] = l.j
		} else {
			l.j = l.next[l.j]
		}
	}

	return l.next[k]
}

// LazyIndex match T in S by lazy kmp algorithm
// next array is trained while matching
func LazyIndex(S, T string, pos int) int {
	lenOfS, lenOfT := len(S), len(T)
	var l lazyNext
	l.Init(T)

	i := pos
	j := 0
	for i < lenOfS && j < lenOfT {
		if j == -1 || S[i] == T[j] {
			i++
			j++
		} else {
			j = l.Next(j)
		}
	}

	if j == lenOfT {
		return i - j
	}

	return -1
}
