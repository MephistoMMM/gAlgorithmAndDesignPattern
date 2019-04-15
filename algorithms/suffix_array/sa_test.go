package suffixarray

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	str1     = "//A can can can a can."
	patlen1  = 4
	pattern1 = " can"
	count1   = 4

	str2     = "int a=~~~~~~~~~~~~~~~~~~~~~0;"
	patlen2  = 3
	pattern2 = "~~~"
	count2   = 19

	str3 = "banana"
)

const letterBytes = " 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

var str = randStringBytesMaskImprSrc(10000)

func verify(str string, sufarr []int) bool {
	for i := 1; i < len(str); i++ {
		prev := str[sufarr[i-1]:]
		cur := str[sufarr[i]:]
		if cur < prev {
			fmt.Printf("prev(%s) bigger than cur(%s)\n", prev, cur)
			return false
		}
	}

	return true
}

func TestBuildByCmpSort(t *testing.T) {
	sufarr1 := BuildByRadixSort(str1)
	ShowSufarr(str1, sufarr1)
	if ret := verify(str1, sufarr1); !ret {
		t.Error(">> RadixSort Error! <<")
	}

	sufarr1 = RawSort(str1)
	ShowSufarr(str1, sufarr1)
	if ret := verify(str1, sufarr1); !ret {
		t.Error(">> RawSort Error! <<")
	}

	sufarr1 = BuildByCmpSort(str1)
	ShowSufarr(str1, sufarr1)
	if ret := verify(str1, sufarr1); !ret {
		t.Error(">> CmpSort Error! <<")
	}

	sufarr1 = BuildByRadixSortVerbose(str1)
	ShowSufarr(str1, sufarr1)
	if ret := verify(str1, sufarr1); !ret {
		t.Error(">> RadixSortVerbose Error! <<")
	}
}

// func BenchmarkRawSort(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		RawSort(str)
// 	}
// }

// func BenchmarkCmpSort(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		BuildByCmpSort(str)
// 	}
// }

// func BenchmarkRadixSort1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		BuildByRadixSortVerbose(str)
// 	}
// }

func BenchmarkRadixSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildByRadixSort(str)
	}
}
