package kmp

import (
	"testing"
)

func TestMultipleTNodeMatch(t *testing.T) {
	T := NewTNode("ababaaababaa")
	S := "ababababababaaababaabbababbb"

	index := T.Match(S, 0, len(S))

	if index != 8 {
		t.Errorf("index should be 8, but get %v.\n", index)
	}

	S = "ababaaababaaababaaababaabbababbb"
	index = T.Match(S, 10, len(S))
	if index != 12 {
		t.Errorf("index should be 12, but get %v.\n", index)
	}

	S = "babaabbababbbbabaabbababbbbabaabbababbb"
	index = T.Match(S, 10, len(S))
	if index != -1 {
		t.Errorf("index should be -1, but get %v.\n", index)
	}
}

func TestMultipleKMPMatchMchine(t *testing.T) {
	kmm := NewKMPMatchMachine("not", "like", "insert", "update", "delete")
	S := `insert into test_table values(123, insert 321, 1123)
insert into test_table values(123, insert 321, 1123)
update into test_table values(123, insert 321, 1123)
insert into test_table values(123, insert 321, 1123)
`

	matched := kmm.MatchAll(S)
	if matched != 2 {
		t.Errorf("matched should be 2, but get %v.\n", matched)
	}
	if matchedOfNode := kmm.patterns4m[2].matched; matchedOfNode != 0 {
		t.Errorf("matched of node should be 0, but get %v.\n", matchedOfNode)
	}
	t.Log(kmm.patterns4m[3].matched)
	t.Log(S[kmm.patterns4m[3].matched : kmm.patterns4m[3].matched+6])

	S = `update into test_table values(123, update 321, 1123)
update into test_table values(123, update 321, 1123)
insert into test_table values(123, update 321, 1123)
update into test_table values(123, update 321, 1123)
`
	matched = kmm.MatchAll(S)
	if matched != 2 {
		t.Errorf("matched should be 0, but get %v.\n", matched)
	}
	if matchedOfNode := kmm.patterns4m[2].matched; matchedOfNode != 106 {
		t.Errorf("matched of node should be 106, but get %v.\n", matchedOfNode)
	}
}
