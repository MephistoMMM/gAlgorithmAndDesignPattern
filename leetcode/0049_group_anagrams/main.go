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
package main

import "fmt"

// https://leetcode.com/problems/group-anagrams/

type Signature [27]int

// name ...
func (st *Signature) Equal(other *Signature) bool {
	for i := range st {
		if (*st)[i] != (*other)[i] {
			return false
		}
	}

	return true
}

type Record struct {
	ID *Signature
	V  string

	Next *Record
}

func NewRecord(v string) *Record {
	record := &Record{
		V: v,
	}
	record.sign()
	return record
}

// sign ...
func (rc *Record) sign() {
	rc.ID = &Signature{}

	// store length of string to the first position of ID
	(*rc.ID)[0] = len(rc.V)

	for _, v := range rc.V {
		(*rc.ID)[int(v-96)]++
	}
}

type RecordGroup struct {
	ID *Signature

	Elems  *Record
	Length int
}

// Contain ...
func (rg *RecordGroup) Contain(o *Record) bool {
	return rg.ID.Equal(o.ID)
}

func (rg *RecordGroup) Add(o *Record) {
	o.Next = rg.Elems
	rg.Elems = o
	rg.Length++
}

// 224 ms	203.7 MB
func groupAnagrams(strs []string) [][]string {
	// first, classify
	anagrams := make(map[Signature]*RecordGroup)
	for _, v := range strs {
		rc := NewRecord(v)
		if anagram, ok := anagrams[*rc.ID]; ok {
			anagram.Add(rc)
		}

		anagram := &RecordGroup{
			ID:     rc.ID,
			Elems:  rc,
			Length: 1,
		}
		anagrams[*rc.ID] = anagram
	}
	// second, construct result
	results := make([][]string, 0, len(anagrams))
	for _, anagram := range anagrams {
		result := make([]string, 0, anagram.Length)
		rc := anagram.Elems
		for rc != nil {
			result = append(result, rc.V)
			rc = rc.Next
		}

		results = append(results, result)
	}
	return results
}

// 712 ms	340.3 MB
func groupAnagrams2(strs []string) [][]string {
	// first, classify
	anagrams := make([]*RecordGroup, 0)
	for _, v := range strs {
		rc := NewRecord(v)
		found := false
		for _, anagram := range anagrams {
			if anagram.Contain(rc) {
				anagram.Add(rc)
				found = true
				break
			}
		}

		if !found {
			anagram := &RecordGroup{
				ID:     rc.ID,
				Elems:  rc,
				Length: 1,
			}
			anagrams = append(anagrams, anagram)
		}
	}
	// second, construct result
	results := make([][]string, 0, len(anagrams))
	for _, anagram := range anagrams {
		result := make([]string, 0, anagram.Length)
		rc := anagram.Elems
		for rc != nil {
			result = append(result, rc.V)
			rc = rc.Next
		}

		results = append(results, result)
	}
	return results
}

func showOutput(output [][]string) {
	fmt.Println("{")
	for _, v := range output {
		fmt.Printf("\t%v\n", v)
	}
	fmt.Println("}")
}

func main() {
	showOutput(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
