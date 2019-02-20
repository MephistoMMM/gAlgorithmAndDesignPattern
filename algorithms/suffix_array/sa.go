package suffixarray

import (
	"fmt"
	"math"
	"sort"
)

const MAXLENGTH = 1048576

// ShowSufarr print suffix array
func ShowSufarr(str string, sufarr []int) {
	for _, i := range sufarr {
		fmt.Printf("%s\n", str[i:])
	}
	fmt.Println("")
}

// What we need ?
// 1. sufarr - i: index in partial ordered suffix array, v: suffix in string
// 2. rank - i: suffix in string, v: rank in partial ordered suffiex array
// 3. rankAdj - i: suffix in string, v: adjancy rank in partial ordered suffiex array
// 4. rankAdjnrr - i: index in ordered adjancy rank, v: suffix in string

// the final sufarr is the ordered suffix array

var rank, rankAdj *[MAXLENGTH]int
var rankAdjnrr *[MAXLENGTH]int
var tempSufarr [MAXLENGTH]int
var radixSortArr [MAXLENGTH]int

func init() {
	rank = &[MAXLENGTH]int{}
	rankAdj = &[MAXLENGTH]int{}
	rankAdjnrr = &[MAXLENGTH]int{}
}

type rawsuff struct {
	sufarr []int
	str    string
}

// Len ...
func (rs *rawsuff) Len() int {
	return len(rs.sufarr)
}

// Less ...
func (rs *rawsuff) Less(i, j int) bool {

	return string(rs.str[rs.sufarr[i]:]) < string(rs.str[rs.sufarr[j]:])
}

// Swap ...
func (rs *rawsuff) Swap(i, j int) {
	rs.sufarr[i], rs.sufarr[j] = rs.sufarr[j], rs.sufarr[i]
}

func RawSort(str string) []int {
	strLen := len(str)
	sufarr := make([]int, strLen)

	for i := range sufarr {
		sufarr[i] = i
	}

	sort.Sort(&rawsuff{
		sufarr: sufarr,
		str:    str,
	})

	return sufarr

}

func isEqual(i, j, l int) bool {
	strLen := len(rank)
	var (
		iAdjRank int
		jAdjRank int
	)

	if i+l >= strLen {
		iAdjRank = -1
	} else {
		iAdjRank = rank[i+l]
	}
	if j+l >= strLen {
		jAdjRank = -1
	} else {
		jAdjRank = rank[j+l]
	}

	return rank[i] == rank[j] && iAdjRank == jAdjRank
}

// BuildByRadixSort ...
func BuildByRadixSort(str string) []int {
	strLen := len(str)
	radixLen := int(math.Max(float64(strLen), 256))

	sufarr := make([]int, strLen)

	// Init Index And Rank
	for i := range str {
		sufarr[i] = i
		rank[i] = int(str[i])
	}

	var radixTemp int
	for i, l, k := 0, 0, 1; k < strLen; k <<= 1 {
		l = k - 1

		// Generate RankAdjIndex According To Index And Rank
		p := 0
		for i = strLen - l; i < strLen; i++ {
			rankAdjnrr[p] = i
			p++
		}
		for i = 0; i < strLen; i++ {
			if sufarr[i] >= l {
				rankAdjnrr[p] = sufarr[i] - l
				p++
			}
		}

		// Radix Sort Suffixes According To Rank And RankAdj
		// RadixSortSufarr(sufarr, rank, rankAdj)
		for i = 0; i < radixLen; i++ {
			radixSortArr[i] = 0
		}
		for i = 0; i < strLen; i++ {
			radixSortArr[rank[i]]++
		}
		radixTemp = radixSortArr[0]
		for i = 1; i < radixLen; i++ {
			radixSortArr[i] += radixTemp
			radixTemp = radixSortArr[i]
		}
		for i = strLen - 1; i > -1; i-- {
			radixTemp = rankAdjnrr[i]
			radixSortArr[rank[radixTemp]]--
			sufarr[radixSortArr[rank[radixTemp]]] = radixTemp
		}

		// Assigning new rank to suffixes
		// After Radix Sort, rankAdjnrr could be used as the new rank
		newRank := rankAdjnrr
		r := 1
		newRank[sufarr[0]] = 0
		for i = 1; i < strLen; i++ {
			if isEqual(sufarr[i], sufarr[i-1], l) {
				newRank[sufarr[i]] = r - 1
			} else {
				newRank[sufarr[i]] = r
				r++
			}
		}
		// Let Origin Rank Be The New RankAdjRank
		rank, rankAdjnrr = newRank, rank
	}

	return sufarr
}

// RadixSortVerbose ...
func RadixSortVerbose(sufarr []int) {
	strLen := len(sufarr)
	radixLen := int(math.Max(float64(strLen), 256)) + 1

	// count sort according to adjancy rank
	var i, t int
	for i = 0; i < radixLen; i++ {
		radixSortArr[i] = 0
	}
	for i = 0; i < strLen; i++ {
		radixSortArr[rankAdj[i]+1]++
	}
	t = radixSortArr[0]
	for i = 1; i < radixLen; i++ {
		radixSortArr[i] += t
		t = radixSortArr[i]
	}
	for i = strLen - 1; i > -1; i-- {
		radixSortArr[rankAdj[i]+1]--
		tempSufarr[radixSortArr[rankAdj[i]+1]] = i
	}

	// count sort according to rank
	for i = 0; i < radixLen; i++ {
		radixSortArr[i] = 0
	}
	for i = 0; i < strLen; i++ {
		radixSortArr[rank[i]]++
	}
	t = radixSortArr[0]
	for i = 1; i < radixLen; i++ {
		radixSortArr[i] += t
		t = radixSortArr[i]
	}
	for i = strLen - 1; i > -1; i-- {
		t = tempSufarr[i]
		radixSortArr[rank[t]]--
		sufarr[radixSortArr[rank[t]]] = t
	}
}

// BuildByRadixSortVerbose ...
func BuildByRadixSortVerbose(str string) []int {
	strLen := len(str)
	sufarr := make([]int, strLen)

	// Init Index And Rank
	for i := range str {
		sufarr[i] = i
		rank[i] = int(str[i])
	}

	for i, l, k := 0, 0, 1; k < strLen; k <<= 1 {
		l = k - 1

		// Generate RankAdj According To Index And Rank
		for i = 0; i < strLen; i++ {
			if i+l >= strLen {
				rankAdj[i] = -1
			} else {
				rankAdj[i] = rank[i+l]
			}
		}

		// Radix Sort Suffixes According To Rank And RankAdj
		RadixSortVerbose(sufarr)

		// Assigning new rank to suffixes
		// After Radix Sort, rankAdjSufarr could be used as the new rank
		newRank := rankAdj
		r := 1
		newRank[sufarr[0]] = 0
		for i := 1; i < strLen; i++ {
			curSuf, prevSuf := sufarr[i], sufarr[i-1]
			if isEqual(curSuf, prevSuf, l) {
				newRank[curSuf] = r - 1
			} else {
				newRank[curSuf] = r
				r++
			}
		}
		// Let Origin Rank Be The New RankAdjRank
		rank, rankAdj = newRank, rank
	}

	return sufarr
}

type suffix struct {
	i    int
	rank [2]int
}

type suffixes []suffix

// Len ...
func (ss suffixes) Len() int {
	return len(ss)
}

// Less ...
// rank[0] is the first order, rank[1] is the second order
func (ss suffixes) Less(i, j int) bool {
	a, b := &ss[i], &ss[j]
	if a.rank[0] == b.rank[0] {
		return a.rank[1] < b.rank[1]
	}

	return a.rank[0] < b.rank[0]
}

// funcname ...
func (ss suffixes) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

// A structure to store suffixes and their sufarr
// k: index, v: suffix
var sufarr [MAXLENGTH]suffix

// Equal ...
func Equal(ss *[MAXLENGTH]suffix, i, j int) bool {
	a, b := &ss[i], &ss[j]
	if a.rank[0] == b.rank[0] && a.rank[1] == b.rank[1] {
		return true
	}

	return false
}

// k: suffix, v: index
var sufnrr [MAXLENGTH]int

// BuildByCmpSort create suffix array for str
// str should only include acsii code
func BuildByCmpSort(str string) []int {
	strLen := len(str)

	// Store suffixes and their sufarr in an array of structures.
	// The structure is needed to sort the suffixes alphabatically
	// and maintain their old sufarr while sorting
	for i := 0; i < strLen; i++ {
		sufarr[i].i = i
		sufarr[i].rank[0] = int(str[i])
		sufnrr[i] = i
	}

	// At his point, all suffixes are sorted according to first
	// 2 characters.  Let us sort suffixes according to first 4
	// characters, then first 8 and so on
	for i, k, l := 0, 1, 0; k < strLen; k <<= 1 {
		l = k - 1

		// Assign next rank to every suffix
		for i = 0; i < strLen; i++ {
			adjSuf := sufarr[i].i + l
			if adjSuf < strLen {
				sufarr[i].rank[1] = sufarr[sufnrr[adjSuf]].rank[0]
			} else {
				sufarr[i].rank[1] = -1
			}
		}

		// Sort the suffixes according to first k characters
		sort.Sort(suffixes(sufarr[:strLen]))

		// Assigning rank and index values to first suffix
		// Assigning rank to suffixes
		r := 0
		for i = 1; i < strLen; i++ {
			sufnrr[sufarr[i-1].i] = i - 1
			if Equal(&sufarr, i, i-1) {
				sufarr[i-1].rank[0] = r
			} else {
				sufarr[i-1].rank[0] = r
				r++
			}
		}
		sufnrr[sufarr[i-1].i] = i - 1
		sufarr[i-1].rank[0] = r
	}

	suffixArray := make([]int, strLen)
	// Store sufarr of all sorted suffixes in the suffix array
	for i := 0; i < strLen; i++ {
		suffixArray[i] = sufarr[i].i
	}

	return suffixArray
}
