package kmp

import (
	"reflect"
	"testing"
)

func TestTrainNext(t *testing.T) {
	T := "ababaaababaa"
	next := make([]int, len(T))
	res := []int{-1, 0, 0, 1, 2, 3, 1, 1, 2, 3, 4, 5}

	trainNext(T, next)
	if !reflect.DeepEqual(next, res) {
		t.Errorf("next should be %v, but get %v.\n", res, next)
	}
}

func TestIndex(t *testing.T) {
	T := "ababaaababaa"
	S := "ababababababaaababaabbababbb"

	index := Index(S, T, 0)

	if index != 8 {
		t.Errorf("index should be 8, but get %v.\n", index)
	}

	S = "ababaaababaaababaaababaabbababbb"
	index = Index(S, T, 10)
	if index != 12 {
		t.Errorf("index should be 12, but get %v.\n", index)
	}

	S = "babaabbababbbbabaabbababbbbabaabbababbb"
	index = Index(S, T, 10)
	if index != -1 {
		t.Errorf("index should be -1, but get %v.\n", index)
	}
}

func TestLazyIndex(t *testing.T) {
	T := "ababaaababaa"
	S := "abcdefghijklmnopqrstuvwxyz"

	index := LazyIndex(S, T, 0)
	if index != -1 {
		t.Errorf("index should be -1, but get %v.\n", index)
	}

	S = "ababababababaaababaabbababbb"
	index = LazyIndex(S, T, 0)
	if index != 8 {
		t.Errorf("index should be 8, but get %v.\n", index)
	}

	S = "ababaaababaaababaaababaabbababbb"
	index = LazyIndex(S, T, 10)
	if index != 12 {
		t.Errorf("index should be 12, but get %v.\n", index)
	}

	S = "babaabbababbbbabaabbababbbbabaabbababbb"
	index = LazyIndex(S, T, 10)
	if index != -1 {
		t.Errorf("index should be -1, but get %v.\n", index)
	}

}
