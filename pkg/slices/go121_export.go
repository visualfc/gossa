// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package slices

import (
	_ "slices"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "slices",
		Path: "slices",
		Deps: map[string]string{
			"cmp":       "cmp",
			"math/bits": "bits",
			"unsafe":    "unsafe",
		},
		Source: source,
	})
}

var source = "package slices\n\nimport (\n\t\"cmp\"\n\t\"unsafe\"\n\t\"math/bits\"\n)\n\nfunc Equal[S ~[]E, E comparable](s1, s2 S) bool {\n\tif len(s1) != len(s2) {\n\t\treturn false\n\t}\n\tfor i := range s1 {\n\t\tif s1[i] != s2[i] {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\n\nfunc EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {\n\tif len(s1) != len(s2) {\n\t\treturn false\n\t}\n\tfor i, v1 := range s1 {\n\t\tv2 := s2[i]\n\t\tif !eq(v1, v2) {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\n\nfunc Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int {\n\tfor i, v1 := range s1 {\n\t\tif i >= len(s2) {\n\t\t\treturn +1\n\t\t}\n\t\tv2 := s2[i]\n\t\tif c := cmp.Compare(v1, v2); c != 0 {\n\t\t\treturn c\n\t\t}\n\t}\n\tif len(s1) < len(s2) {\n\t\treturn -1\n\t}\n\treturn 0\n}\n\nfunc CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int {\n\tfor i, v1 := range s1 {\n\t\tif i >= len(s2) {\n\t\t\treturn +1\n\t\t}\n\t\tv2 := s2[i]\n\t\tif c := cmp(v1, v2); c != 0 {\n\t\t\treturn c\n\t\t}\n\t}\n\tif len(s1) < len(s2) {\n\t\treturn -1\n\t}\n\treturn 0\n}\n\nfunc Index[S ~[]E, E comparable](s S, v E) int {\n\tfor i := range s {\n\t\tif v == s[i] {\n\t\t\treturn i\n\t\t}\n\t}\n\treturn -1\n}\n\nfunc IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {\n\tfor i := range s {\n\t\tif f(s[i]) {\n\t\t\treturn i\n\t\t}\n\t}\n\treturn -1\n}\n\nfunc Contains[S ~[]E, E comparable](s S, v E) bool {\n\treturn Index(s, v) >= 0\n}\n\nfunc ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {\n\treturn IndexFunc(s, f) >= 0\n}\n\nfunc Insert[S ~[]E, E any](s S, i int, v ...E) S {\n\tm := len(v)\n\tif m == 0 {\n\t\treturn s\n\t}\n\tn := len(s)\n\tif i == n {\n\t\treturn append(s, v...)\n\t}\n\tif n+m > cap(s) {\n\n\t\ts2 := append(s[:i], make(S, n+m-i)...)\n\t\tcopy(s2[i:], v)\n\t\tcopy(s2[i+m:], s[i:])\n\t\treturn s2\n\t}\n\ts = s[:n+m]\n\n\tif !overlaps(v, s[i+m:]) {\n\n\t\tcopy(s[i+m:], s[i:])\n\n\t\tcopy(s[i:], v)\n\n\t\treturn s\n\t}\n\n\tcopy(s[n:], v)\n\n\trotateRight(s[i:], m)\n\n\treturn s\n}\n\nfunc Delete[S ~[]E, E any](s S, i, j int) S {\n\t_ = s[i:j]\n\n\treturn append(s[:i], s[j:]...)\n}\n\nfunc DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {\n\n\tfor i, v := range s {\n\t\tif del(v) {\n\t\t\tj := i\n\t\t\tfor i++; i < len(s); i++ {\n\t\t\t\tv = s[i]\n\t\t\t\tif !del(v) {\n\t\t\t\t\ts[j] = v\n\t\t\t\t\tj++\n\t\t\t\t}\n\t\t\t}\n\t\t\treturn s[:j]\n\t\t}\n\t}\n\treturn s\n}\n\nfunc Replace[S ~[]E, E any](s S, i, j int, v ...E) S {\n\t_ = s[i:j]\n\n\tif i == j {\n\t\treturn Insert(s, i, v...)\n\t}\n\tif j == len(s) {\n\t\treturn append(s[:i], v...)\n\t}\n\n\ttot := len(s[:i]) + len(v) + len(s[j:])\n\tif tot > cap(s) {\n\n\t\ts2 := append(s[:i], make(S, tot-i)...)\n\t\tcopy(s2[i:], v)\n\t\tcopy(s2[i+len(v):], s[j:])\n\t\treturn s2\n\t}\n\n\tr := s[:tot]\n\n\tif i+len(v) <= j {\n\n\t\tcopy(r[i:], v)\n\t\tif i+len(v) != j {\n\t\t\tcopy(r[i+len(v):], s[j:])\n\t\t}\n\t\treturn r\n\t}\n\n\tif !overlaps(r[i+len(v):], v) {\n\n\t\tcopy(r[i+len(v):], s[j:])\n\t\tcopy(r[i:], v)\n\t\treturn r\n\t}\n\n\ty := len(v) - (j - i)\n\n\tif !overlaps(r[i:j], v) {\n\t\tcopy(r[i:j], v[y:])\n\t\tcopy(r[len(s):], v[:y])\n\t\trotateRight(r[i:], y)\n\t\treturn r\n\t}\n\tif !overlaps(r[len(s):], v) {\n\t\tcopy(r[len(s):], v[:y])\n\t\tcopy(r[i:j], v[y:])\n\t\trotateRight(r[i:], y)\n\t\treturn r\n\t}\n\n\tk := startIdx(v, s[j:])\n\tcopy(r[i:], v)\n\tcopy(r[i+len(v):], r[i+k:])\n\treturn r\n}\n\nfunc Clone[S ~[]E, E any](s S) S {\n\n\tif s == nil {\n\t\treturn nil\n\t}\n\treturn append(S([]E{}), s...)\n}\n\nfunc Compact[S ~[]E, E comparable](s S) S {\n\tif len(s) < 2 {\n\t\treturn s\n\t}\n\ti := 1\n\tfor k := 1; k < len(s); k++ {\n\t\tif s[k] != s[k-1] {\n\t\t\tif i != k {\n\t\t\t\ts[i] = s[k]\n\t\t\t}\n\t\t\ti++\n\t\t}\n\t}\n\treturn s[:i]\n}\n\nfunc CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {\n\tif len(s) < 2 {\n\t\treturn s\n\t}\n\ti := 1\n\tfor k := 1; k < len(s); k++ {\n\t\tif !eq(s[k], s[k-1]) {\n\t\t\tif i != k {\n\t\t\t\ts[i] = s[k]\n\t\t\t}\n\t\t\ti++\n\t\t}\n\t}\n\treturn s[:i]\n}\n\nfunc Grow[S ~[]E, E any](s S, n int) S {\n\tif n < 0 {\n\t\tpanic(\"cannot be negative\")\n\t}\n\tif n -= cap(s) - len(s); n > 0 {\n\t\ts = append(s[:cap(s)], make([]E, n)...)[:len(s)]\n\t}\n\treturn s\n}\n\nfunc Clip[S ~[]E, E any](s S) S {\n\treturn s[:len(s):len(s)]\n}\n\nfunc rotateLeft[E any](s []E, r int) {\n\tfor r != 0 && r != len(s) {\n\t\tif r*2 <= len(s) {\n\t\t\tswap(s[:r], s[len(s)-r:])\n\t\t\ts = s[:len(s)-r]\n\t\t} else {\n\t\t\tswap(s[:len(s)-r], s[r:])\n\t\t\ts, r = s[len(s)-r:], r*2-len(s)\n\t\t}\n\t}\n}\nfunc rotateRight[E any](s []E, r int) {\n\trotateLeft(s, len(s)-r)\n}\n\nfunc swap[E any](x, y []E) {\n\tfor i := 0; i < len(x); i++ {\n\t\tx[i], y[i] = y[i], x[i]\n\t}\n}\n\nfunc overlaps[E any](a, b []E) bool {\n\tif len(a) == 0 || len(b) == 0 {\n\t\treturn false\n\t}\n\telemSize := unsafe.Sizeof(a[0])\n\tif elemSize == 0 {\n\t\treturn false\n\t}\n\n\treturn uintptr(unsafe.Pointer(&a[0])) <= uintptr(unsafe.Pointer(&b[len(b)-1]))+(elemSize-1) &&\n\t\tuintptr(unsafe.Pointer(&b[0])) <= uintptr(unsafe.Pointer(&a[len(a)-1]))+(elemSize-1)\n}\n\nfunc startIdx[E any](haystack, needle []E) int {\n\tp := &needle[0]\n\tfor i := range haystack {\n\t\tif p == &haystack[i] {\n\t\t\treturn i\n\t\t}\n\t}\n\n\tpanic(\"needle not found\")\n}\n\nfunc Reverse[S ~[]E, E any](s S) {\n\tfor i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {\n\t\ts[i], s[j] = s[j], s[i]\n\t}\n}\nfunc Sort[S ~[]E, E cmp.Ordered](x S) {\n\tn := len(x)\n\tpdqsortOrdered(x, 0, n, bits.Len(uint(n)))\n}\n\nfunc SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {\n\tn := len(x)\n\tpdqsortCmpFunc(x, 0, n, bits.Len(uint(n)), cmp)\n}\n\nfunc SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {\n\tstableCmpFunc(x, len(x), cmp)\n}\n\nfunc IsSorted[S ~[]E, E cmp.Ordered](x S) bool {\n\tfor i := len(x) - 1; i > 0; i-- {\n\t\tif cmp.Less(x[i], x[i-1]) {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\n\nfunc IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool {\n\tfor i := len(x) - 1; i > 0; i-- {\n\t\tif cmp(x[i], x[i-1]) < 0 {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\n\nfunc Min[S ~[]E, E cmp.Ordered](x S) E {\n\tif len(x) < 1 {\n\t\tpanic(\"slices.Min: empty list\")\n\t}\n\tm := x[0]\n\tfor i := 1; i < len(x); i++ {\n\t\tm = min(m, x[i])\n\t}\n\treturn m\n}\n\nfunc MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {\n\tif len(x) < 1 {\n\t\tpanic(\"slices.MinFunc: empty list\")\n\t}\n\tm := x[0]\n\tfor i := 1; i < len(x); i++ {\n\t\tif cmp(x[i], m) < 0 {\n\t\t\tm = x[i]\n\t\t}\n\t}\n\treturn m\n}\n\nfunc Max[S ~[]E, E cmp.Ordered](x S) E {\n\tif len(x) < 1 {\n\t\tpanic(\"slices.Max: empty list\")\n\t}\n\tm := x[0]\n\tfor i := 1; i < len(x); i++ {\n\t\tm = max(m, x[i])\n\t}\n\treturn m\n}\n\nfunc MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E {\n\tif len(x) < 1 {\n\t\tpanic(\"slices.MaxFunc: empty list\")\n\t}\n\tm := x[0]\n\tfor i := 1; i < len(x); i++ {\n\t\tif cmp(x[i], m) > 0 {\n\t\t\tm = x[i]\n\t\t}\n\t}\n\treturn m\n}\n\nfunc BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {\n\n\tn := len(x)\n\n\ti, j := 0, n\n\tfor i < j {\n\t\th := int(uint(i+j) >> 1)\n\n\t\tif cmp.Less(x[h], target) {\n\t\t\ti = h + 1\n\t\t} else {\n\t\t\tj = h\n\t\t}\n\t}\n\n\treturn i, i < n && (x[i] == target || (isNaN(x[i]) && isNaN(target)))\n}\n\nfunc BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {\n\tn := len(x)\n\n\ti, j := 0, n\n\tfor i < j {\n\t\th := int(uint(i+j) >> 1)\n\n\t\tif cmp(x[h], target) < 0 {\n\t\t\ti = h + 1\n\t\t} else {\n\t\t\tj = h\n\t\t}\n\t}\n\n\treturn i, i < n && cmp(x[i], target) == 0\n}\n\ntype sortedHint int\n\nconst (\n\tunknownHint\tsortedHint\t= iota\n\tincreasingHint\n\tdecreasingHint\n)\n\ntype xorshift uint64\n\nfunc (r *xorshift) Next() uint64 {\n\t*r ^= *r << 13\n\t*r ^= *r >> 17\n\t*r ^= *r << 5\n\treturn uint64(*r)\n}\n\nfunc nextPowerOfTwo(length int) uint {\n\treturn 1 << bits.Len(uint(length))\n}\n\nfunc isNaN[T cmp.Ordered](x T) bool {\n\treturn x != x\n}\nfunc insertionSortCmpFunc[E any](data []E, a, b int, cmp func(a, b E) int) {\n\tfor i := a + 1; i < b; i++ {\n\t\tfor j := i; j > a && (cmp(data[j], data[j-1]) < 0); j-- {\n\t\t\tdata[j], data[j-1] = data[j-1], data[j]\n\t\t}\n\t}\n}\n\nfunc siftDownCmpFunc[E any](data []E, lo, hi, first int, cmp func(a, b E) int) {\n\troot := lo\n\tfor {\n\t\tchild := 2*root + 1\n\t\tif child >= hi {\n\t\t\tbreak\n\t\t}\n\t\tif child+1 < hi && (cmp(data[first+child], data[first+child+1]) < 0) {\n\t\t\tchild++\n\t\t}\n\t\tif !(cmp(data[first+root], data[first+child]) < 0) {\n\t\t\treturn\n\t\t}\n\t\tdata[first+root], data[first+child] = data[first+child], data[first+root]\n\t\troot = child\n\t}\n}\n\nfunc heapSortCmpFunc[E any](data []E, a, b int, cmp func(a, b E) int) {\n\tfirst := a\n\tlo := 0\n\thi := b - a\n\n\tfor i := (hi - 1) / 2; i >= 0; i-- {\n\t\tsiftDownCmpFunc(data, i, hi, first, cmp)\n\t}\n\n\tfor i := hi - 1; i >= 0; i-- {\n\t\tdata[first], data[first+i] = data[first+i], data[first]\n\t\tsiftDownCmpFunc(data, lo, i, first, cmp)\n\t}\n}\n\nfunc pdqsortCmpFunc[E any](data []E, a, b, limit int, cmp func(a, b E) int) {\n\tconst maxInsertion = 12\n\n\tvar (\n\t\twasBalanced\t= true\n\t\twasPartitioned\t= true\n\t)\n\n\tfor {\n\t\tlength := b - a\n\n\t\tif length <= maxInsertion {\n\t\t\tinsertionSortCmpFunc(data, a, b, cmp)\n\t\t\treturn\n\t\t}\n\n\t\tif limit == 0 {\n\t\t\theapSortCmpFunc(data, a, b, cmp)\n\t\t\treturn\n\t\t}\n\n\t\tif !wasBalanced {\n\t\t\tbreakPatternsCmpFunc(data, a, b, cmp)\n\t\t\tlimit--\n\t\t}\n\n\t\tpivot, hint := choosePivotCmpFunc(data, a, b, cmp)\n\t\tif hint == decreasingHint {\n\t\t\treverseRangeCmpFunc(data, a, b, cmp)\n\n\t\t\tpivot = (b - 1) - (pivot - a)\n\t\t\thint = increasingHint\n\t\t}\n\n\t\tif wasBalanced && wasPartitioned && hint == increasingHint {\n\t\t\tif partialInsertionSortCmpFunc(data, a, b, cmp) {\n\t\t\t\treturn\n\t\t\t}\n\t\t}\n\n\t\tif a > 0 && !(cmp(data[a-1], data[pivot]) < 0) {\n\t\t\tmid := partitionEqualCmpFunc(data, a, b, pivot, cmp)\n\t\t\ta = mid\n\t\t\tcontinue\n\t\t}\n\n\t\tmid, alreadyPartitioned := partitionCmpFunc(data, a, b, pivot, cmp)\n\t\twasPartitioned = alreadyPartitioned\n\n\t\tleftLen, rightLen := mid-a, b-mid\n\t\tbalanceThreshold := length / 8\n\t\tif leftLen < rightLen {\n\t\t\twasBalanced = leftLen >= balanceThreshold\n\t\t\tpdqsortCmpFunc(data, a, mid, limit, cmp)\n\t\t\ta = mid + 1\n\t\t} else {\n\t\t\twasBalanced = rightLen >= balanceThreshold\n\t\t\tpdqsortCmpFunc(data, mid+1, b, limit, cmp)\n\t\t\tb = mid\n\t\t}\n\t}\n}\n\nfunc partitionCmpFunc[E any](data []E, a, b, pivot int, cmp func(a, b E) int) (newpivot int, alreadyPartitioned bool) {\n\tdata[a], data[pivot] = data[pivot], data[a]\n\ti, j := a+1, b-1\n\n\tfor i <= j && (cmp(data[i], data[a]) < 0) {\n\t\ti++\n\t}\n\tfor i <= j && !(cmp(data[j], data[a]) < 0) {\n\t\tj--\n\t}\n\tif i > j {\n\t\tdata[j], data[a] = data[a], data[j]\n\t\treturn j, true\n\t}\n\tdata[i], data[j] = data[j], data[i]\n\ti++\n\tj--\n\n\tfor {\n\t\tfor i <= j && (cmp(data[i], data[a]) < 0) {\n\t\t\ti++\n\t\t}\n\t\tfor i <= j && !(cmp(data[j], data[a]) < 0) {\n\t\t\tj--\n\t\t}\n\t\tif i > j {\n\t\t\tbreak\n\t\t}\n\t\tdata[i], data[j] = data[j], data[i]\n\t\ti++\n\t\tj--\n\t}\n\tdata[j], data[a] = data[a], data[j]\n\treturn j, false\n}\n\nfunc partitionEqualCmpFunc[E any](data []E, a, b, pivot int, cmp func(a, b E) int) (newpivot int) {\n\tdata[a], data[pivot] = data[pivot], data[a]\n\ti, j := a+1, b-1\n\n\tfor {\n\t\tfor i <= j && !(cmp(data[a], data[i]) < 0) {\n\t\t\ti++\n\t\t}\n\t\tfor i <= j && (cmp(data[a], data[j]) < 0) {\n\t\t\tj--\n\t\t}\n\t\tif i > j {\n\t\t\tbreak\n\t\t}\n\t\tdata[i], data[j] = data[j], data[i]\n\t\ti++\n\t\tj--\n\t}\n\treturn i\n}\n\nfunc partialInsertionSortCmpFunc[E any](data []E, a, b int, cmp func(a, b E) int) bool {\n\tconst (\n\t\tmaxSteps\t\t= 5\n\t\tshortestShifting\t= 50\n\t)\n\ti := a + 1\n\tfor j := 0; j < maxSteps; j++ {\n\t\tfor i < b && !(cmp(data[i], data[i-1]) < 0) {\n\t\t\ti++\n\t\t}\n\n\t\tif i == b {\n\t\t\treturn true\n\t\t}\n\n\t\tif b-a < shortestShifting {\n\t\t\treturn false\n\t\t}\n\n\t\tdata[i], data[i-1] = data[i-1], data[i]\n\n\t\tif i-a >= 2 {\n\t\t\tfor j := i - 1; j >= 1; j-- {\n\t\t\t\tif !(cmp(data[j], data[j-1]) < 0) {\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t\tdata[j], data[j-1] = data[j-1], data[j]\n\t\t\t}\n\t\t}\n\n\t\tif b-i >= 2 {\n\t\t\tfor j := i + 1; j < b; j++ {\n\t\t\t\tif !(cmp(data[j], data[j-1]) < 0) {\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t\tdata[j], data[j-1] = data[j-1], data[j]\n\t\t\t}\n\t\t}\n\t}\n\treturn false\n}\n\nfunc breakPatternsCmpFunc[E any](data []E, a, b int, cmp func(a, b E) int) {\n\tlength := b - a\n\tif length >= 8 {\n\t\trandom := xorshift(length)\n\t\tmodulus := nextPowerOfTwo(length)\n\n\t\tfor idx := a + (length/4)*2 - 1; idx <= a+(length/4)*2+1; idx++ {\n\t\t\tother := int(uint(random.Next()) & (modulus - 1))\n\t\t\tif other >= length {\n\t\t\t\tother -= length\n\t\t\t}\n\t\t\tdata[idx], data[a+other] = data[a+other], data[idx]\n\t\t}\n\t}\n}\n\nfunc choosePivotCmpFunc[E any](data []E, a, b int, cmp func(a, b E) int) (pivot int, hint sortedHint) {\n\tconst (\n\t\tshortestNinther\t= 50\n\t\tmaxSwaps\t= 4 * 3\n\t)\n\n\tl := b - a\n\n\tvar (\n\t\tswaps\tint\n\t\ti\t= a + l/4*1\n\t\tj\t= a + l/4*2\n\t\tk\t= a + l/4*3\n\t)\n\n\tif l >= 8 {\n\t\tif l >= shortestNinther {\n\n\t\t\ti = medianAdjacentCmpFunc(data, i, &swaps, cmp)\n\t\t\tj = medianAdjacentCmpFunc(data, j, &swaps, cmp)\n\t\t\tk = medianAdjacentCmpFunc(data, k, &swaps, cmp)\n\t\t}\n\n\t\tj = medianCmpFunc(data, i, j, k, &swaps, cmp)\n\t}\n\n\tswitch swaps {\n\tcase 0:\n\t\treturn j, increasingHint\n\tcase maxSwaps:\n\t\treturn j, decreasingHint\n\tdefault:\n\t\treturn j, unknownHint\n\t}\n}\n\nfunc order2CmpFunc[E any](data []E, a, b int, swaps *int, cmp func(a, b E) int) (int, int) {\n\tif cmp(data[b], data[a]) < 0 {\n\t\t*swaps++\n\t\treturn b, a\n\t}\n\treturn a, b\n}\n\nfunc medianCmpFunc[E any](data []E, a, b, c int, swaps *int, cmp func(a, b E) int) int {\n\ta, b = order2CmpFunc(data, a, b, swaps, cmp)\n\tb, c = order2CmpFunc(data, b, c, swaps, cmp)\n\ta, b = order2CmpFunc(data, a, b, swaps, cmp)\n\treturn b\n}\n\nfunc medianAdjacentCmpFunc[E any](data []E, a int, swaps *int, cmp func(a, b E) int) int {\n\treturn medianCmpFunc(data, a-1, a, a+1, swaps, cmp)\n}\n\nfunc reverseRangeCmpFunc[E any](data []E, a, b int, cmp func(a, b E) int) {\n\ti := a\n\tj := b - 1\n\tfor i < j {\n\t\tdata[i], data[j] = data[j], data[i]\n\t\ti++\n\t\tj--\n\t}\n}\n\nfunc swapRangeCmpFunc[E any](data []E, a, b, n int, cmp func(a, b E) int) {\n\tfor i := 0; i < n; i++ {\n\t\tdata[a+i], data[b+i] = data[b+i], data[a+i]\n\t}\n}\n\nfunc stableCmpFunc[E any](data []E, n int, cmp func(a, b E) int) {\n\tblockSize := 20\n\ta, b := 0, blockSize\n\tfor b <= n {\n\t\tinsertionSortCmpFunc(data, a, b, cmp)\n\t\ta = b\n\t\tb += blockSize\n\t}\n\tinsertionSortCmpFunc(data, a, n, cmp)\n\n\tfor blockSize < n {\n\t\ta, b = 0, 2*blockSize\n\t\tfor b <= n {\n\t\t\tsymMergeCmpFunc(data, a, a+blockSize, b, cmp)\n\t\t\ta = b\n\t\t\tb += 2 * blockSize\n\t\t}\n\t\tif m := a + blockSize; m < n {\n\t\t\tsymMergeCmpFunc(data, a, m, n, cmp)\n\t\t}\n\t\tblockSize *= 2\n\t}\n}\n\nfunc symMergeCmpFunc[E any](data []E, a, m, b int, cmp func(a, b E) int) {\n\n\tif m-a == 1 {\n\n\t\ti := m\n\t\tj := b\n\t\tfor i < j {\n\t\t\th := int(uint(i+j) >> 1)\n\t\t\tif cmp(data[h], data[a]) < 0 {\n\t\t\t\ti = h + 1\n\t\t\t} else {\n\t\t\t\tj = h\n\t\t\t}\n\t\t}\n\n\t\tfor k := a; k < i-1; k++ {\n\t\t\tdata[k], data[k+1] = data[k+1], data[k]\n\t\t}\n\t\treturn\n\t}\n\n\tif b-m == 1 {\n\n\t\ti := a\n\t\tj := m\n\t\tfor i < j {\n\t\t\th := int(uint(i+j) >> 1)\n\t\t\tif !(cmp(data[m], data[h]) < 0) {\n\t\t\t\ti = h + 1\n\t\t\t} else {\n\t\t\t\tj = h\n\t\t\t}\n\t\t}\n\n\t\tfor k := m; k > i; k-- {\n\t\t\tdata[k], data[k-1] = data[k-1], data[k]\n\t\t}\n\t\treturn\n\t}\n\n\tmid := int(uint(a+b) >> 1)\n\tn := mid + m\n\tvar start, r int\n\tif m > mid {\n\t\tstart = n - b\n\t\tr = mid\n\t} else {\n\t\tstart = a\n\t\tr = m\n\t}\n\tp := n - 1\n\n\tfor start < r {\n\t\tc := int(uint(start+r) >> 1)\n\t\tif !(cmp(data[p-c], data[c]) < 0) {\n\t\t\tstart = c + 1\n\t\t} else {\n\t\t\tr = c\n\t\t}\n\t}\n\n\tend := n - start\n\tif start < m && m < end {\n\t\trotateCmpFunc(data, start, m, end, cmp)\n\t}\n\tif a < start && start < mid {\n\t\tsymMergeCmpFunc(data, a, start, mid, cmp)\n\t}\n\tif mid < end && end < b {\n\t\tsymMergeCmpFunc(data, mid, end, b, cmp)\n\t}\n}\n\nfunc rotateCmpFunc[E any](data []E, a, m, b int, cmp func(a, b E) int) {\n\ti := m - a\n\tj := b - m\n\n\tfor i != j {\n\t\tif i > j {\n\t\t\tswapRangeCmpFunc(data, m-i, m, j, cmp)\n\t\t\ti -= j\n\t\t} else {\n\t\t\tswapRangeCmpFunc(data, m-i, m+j-i, i, cmp)\n\t\t\tj -= i\n\t\t}\n\t}\n\n\tswapRangeCmpFunc(data, m-i, m, i, cmp)\n}\nfunc insertionSortOrdered[E cmp.Ordered](data []E, a, b int) {\n\tfor i := a + 1; i < b; i++ {\n\t\tfor j := i; j > a && cmp.Less(data[j], data[j-1]); j-- {\n\t\t\tdata[j], data[j-1] = data[j-1], data[j]\n\t\t}\n\t}\n}\n\nfunc siftDownOrdered[E cmp.Ordered](data []E, lo, hi, first int) {\n\troot := lo\n\tfor {\n\t\tchild := 2*root + 1\n\t\tif child >= hi {\n\t\t\tbreak\n\t\t}\n\t\tif child+1 < hi && cmp.Less(data[first+child], data[first+child+1]) {\n\t\t\tchild++\n\t\t}\n\t\tif !cmp.Less(data[first+root], data[first+child]) {\n\t\t\treturn\n\t\t}\n\t\tdata[first+root], data[first+child] = data[first+child], data[first+root]\n\t\troot = child\n\t}\n}\n\nfunc heapSortOrdered[E cmp.Ordered](data []E, a, b int) {\n\tfirst := a\n\tlo := 0\n\thi := b - a\n\n\tfor i := (hi - 1) / 2; i >= 0; i-- {\n\t\tsiftDownOrdered(data, i, hi, first)\n\t}\n\n\tfor i := hi - 1; i >= 0; i-- {\n\t\tdata[first], data[first+i] = data[first+i], data[first]\n\t\tsiftDownOrdered(data, lo, i, first)\n\t}\n}\n\nfunc pdqsortOrdered[E cmp.Ordered](data []E, a, b, limit int) {\n\tconst maxInsertion = 12\n\n\tvar (\n\t\twasBalanced\t= true\n\t\twasPartitioned\t= true\n\t)\n\n\tfor {\n\t\tlength := b - a\n\n\t\tif length <= maxInsertion {\n\t\t\tinsertionSortOrdered(data, a, b)\n\t\t\treturn\n\t\t}\n\n\t\tif limit == 0 {\n\t\t\theapSortOrdered(data, a, b)\n\t\t\treturn\n\t\t}\n\n\t\tif !wasBalanced {\n\t\t\tbreakPatternsOrdered(data, a, b)\n\t\t\tlimit--\n\t\t}\n\n\t\tpivot, hint := choosePivotOrdered(data, a, b)\n\t\tif hint == decreasingHint {\n\t\t\treverseRangeOrdered(data, a, b)\n\n\t\t\tpivot = (b - 1) - (pivot - a)\n\t\t\thint = increasingHint\n\t\t}\n\n\t\tif wasBalanced && wasPartitioned && hint == increasingHint {\n\t\t\tif partialInsertionSortOrdered(data, a, b) {\n\t\t\t\treturn\n\t\t\t}\n\t\t}\n\n\t\tif a > 0 && !cmp.Less(data[a-1], data[pivot]) {\n\t\t\tmid := partitionEqualOrdered(data, a, b, pivot)\n\t\t\ta = mid\n\t\t\tcontinue\n\t\t}\n\n\t\tmid, alreadyPartitioned := partitionOrdered(data, a, b, pivot)\n\t\twasPartitioned = alreadyPartitioned\n\n\t\tleftLen, rightLen := mid-a, b-mid\n\t\tbalanceThreshold := length / 8\n\t\tif leftLen < rightLen {\n\t\t\twasBalanced = leftLen >= balanceThreshold\n\t\t\tpdqsortOrdered(data, a, mid, limit)\n\t\t\ta = mid + 1\n\t\t} else {\n\t\t\twasBalanced = rightLen >= balanceThreshold\n\t\t\tpdqsortOrdered(data, mid+1, b, limit)\n\t\t\tb = mid\n\t\t}\n\t}\n}\n\nfunc partitionOrdered[E cmp.Ordered](data []E, a, b, pivot int) (newpivot int, alreadyPartitioned bool) {\n\tdata[a], data[pivot] = data[pivot], data[a]\n\ti, j := a+1, b-1\n\n\tfor i <= j && cmp.Less(data[i], data[a]) {\n\t\ti++\n\t}\n\tfor i <= j && !cmp.Less(data[j], data[a]) {\n\t\tj--\n\t}\n\tif i > j {\n\t\tdata[j], data[a] = data[a], data[j]\n\t\treturn j, true\n\t}\n\tdata[i], data[j] = data[j], data[i]\n\ti++\n\tj--\n\n\tfor {\n\t\tfor i <= j && cmp.Less(data[i], data[a]) {\n\t\t\ti++\n\t\t}\n\t\tfor i <= j && !cmp.Less(data[j], data[a]) {\n\t\t\tj--\n\t\t}\n\t\tif i > j {\n\t\t\tbreak\n\t\t}\n\t\tdata[i], data[j] = data[j], data[i]\n\t\ti++\n\t\tj--\n\t}\n\tdata[j], data[a] = data[a], data[j]\n\treturn j, false\n}\n\nfunc partitionEqualOrdered[E cmp.Ordered](data []E, a, b, pivot int) (newpivot int) {\n\tdata[a], data[pivot] = data[pivot], data[a]\n\ti, j := a+1, b-1\n\n\tfor {\n\t\tfor i <= j && !cmp.Less(data[a], data[i]) {\n\t\t\ti++\n\t\t}\n\t\tfor i <= j && cmp.Less(data[a], data[j]) {\n\t\t\tj--\n\t\t}\n\t\tif i > j {\n\t\t\tbreak\n\t\t}\n\t\tdata[i], data[j] = data[j], data[i]\n\t\ti++\n\t\tj--\n\t}\n\treturn i\n}\n\nfunc partialInsertionSortOrdered[E cmp.Ordered](data []E, a, b int) bool {\n\tconst (\n\t\tmaxSteps\t\t= 5\n\t\tshortestShifting\t= 50\n\t)\n\ti := a + 1\n\tfor j := 0; j < maxSteps; j++ {\n\t\tfor i < b && !cmp.Less(data[i], data[i-1]) {\n\t\t\ti++\n\t\t}\n\n\t\tif i == b {\n\t\t\treturn true\n\t\t}\n\n\t\tif b-a < shortestShifting {\n\t\t\treturn false\n\t\t}\n\n\t\tdata[i], data[i-1] = data[i-1], data[i]\n\n\t\tif i-a >= 2 {\n\t\t\tfor j := i - 1; j >= 1; j-- {\n\t\t\t\tif !cmp.Less(data[j], data[j-1]) {\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t\tdata[j], data[j-1] = data[j-1], data[j]\n\t\t\t}\n\t\t}\n\n\t\tif b-i >= 2 {\n\t\t\tfor j := i + 1; j < b; j++ {\n\t\t\t\tif !cmp.Less(data[j], data[j-1]) {\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t\tdata[j], data[j-1] = data[j-1], data[j]\n\t\t\t}\n\t\t}\n\t}\n\treturn false\n}\n\nfunc breakPatternsOrdered[E cmp.Ordered](data []E, a, b int) {\n\tlength := b - a\n\tif length >= 8 {\n\t\trandom := xorshift(length)\n\t\tmodulus := nextPowerOfTwo(length)\n\n\t\tfor idx := a + (length/4)*2 - 1; idx <= a+(length/4)*2+1; idx++ {\n\t\t\tother := int(uint(random.Next()) & (modulus - 1))\n\t\t\tif other >= length {\n\t\t\t\tother -= length\n\t\t\t}\n\t\t\tdata[idx], data[a+other] = data[a+other], data[idx]\n\t\t}\n\t}\n}\n\nfunc choosePivotOrdered[E cmp.Ordered](data []E, a, b int) (pivot int, hint sortedHint) {\n\tconst (\n\t\tshortestNinther\t= 50\n\t\tmaxSwaps\t= 4 * 3\n\t)\n\n\tl := b - a\n\n\tvar (\n\t\tswaps\tint\n\t\ti\t= a + l/4*1\n\t\tj\t= a + l/4*2\n\t\tk\t= a + l/4*3\n\t)\n\n\tif l >= 8 {\n\t\tif l >= shortestNinther {\n\n\t\t\ti = medianAdjacentOrdered(data, i, &swaps)\n\t\t\tj = medianAdjacentOrdered(data, j, &swaps)\n\t\t\tk = medianAdjacentOrdered(data, k, &swaps)\n\t\t}\n\n\t\tj = medianOrdered(data, i, j, k, &swaps)\n\t}\n\n\tswitch swaps {\n\tcase 0:\n\t\treturn j, increasingHint\n\tcase maxSwaps:\n\t\treturn j, decreasingHint\n\tdefault:\n\t\treturn j, unknownHint\n\t}\n}\n\nfunc order2Ordered[E cmp.Ordered](data []E, a, b int, swaps *int) (int, int) {\n\tif cmp.Less(data[b], data[a]) {\n\t\t*swaps++\n\t\treturn b, a\n\t}\n\treturn a, b\n}\n\nfunc medianOrdered[E cmp.Ordered](data []E, a, b, c int, swaps *int) int {\n\ta, b = order2Ordered(data, a, b, swaps)\n\tb, c = order2Ordered(data, b, c, swaps)\n\ta, b = order2Ordered(data, a, b, swaps)\n\treturn b\n}\n\nfunc medianAdjacentOrdered[E cmp.Ordered](data []E, a int, swaps *int) int {\n\treturn medianOrdered(data, a-1, a, a+1, swaps)\n}\n\nfunc reverseRangeOrdered[E cmp.Ordered](data []E, a, b int) {\n\ti := a\n\tj := b - 1\n\tfor i < j {\n\t\tdata[i], data[j] = data[j], data[i]\n\t\ti++\n\t\tj--\n\t}\n}\n\nfunc swapRangeOrdered[E cmp.Ordered](data []E, a, b, n int) {\n\tfor i := 0; i < n; i++ {\n\t\tdata[a+i], data[b+i] = data[b+i], data[a+i]\n\t}\n}\n\nfunc stableOrdered[E cmp.Ordered](data []E, n int) {\n\tblockSize := 20\n\ta, b := 0, blockSize\n\tfor b <= n {\n\t\tinsertionSortOrdered(data, a, b)\n\t\ta = b\n\t\tb += blockSize\n\t}\n\tinsertionSortOrdered(data, a, n)\n\n\tfor blockSize < n {\n\t\ta, b = 0, 2*blockSize\n\t\tfor b <= n {\n\t\t\tsymMergeOrdered(data, a, a+blockSize, b)\n\t\t\ta = b\n\t\t\tb += 2 * blockSize\n\t\t}\n\t\tif m := a + blockSize; m < n {\n\t\t\tsymMergeOrdered(data, a, m, n)\n\t\t}\n\t\tblockSize *= 2\n\t}\n}\n\nfunc symMergeOrdered[E cmp.Ordered](data []E, a, m, b int) {\n\n\tif m-a == 1 {\n\n\t\ti := m\n\t\tj := b\n\t\tfor i < j {\n\t\t\th := int(uint(i+j) >> 1)\n\t\t\tif cmp.Less(data[h], data[a]) {\n\t\t\t\ti = h + 1\n\t\t\t} else {\n\t\t\t\tj = h\n\t\t\t}\n\t\t}\n\n\t\tfor k := a; k < i-1; k++ {\n\t\t\tdata[k], data[k+1] = data[k+1], data[k]\n\t\t}\n\t\treturn\n\t}\n\n\tif b-m == 1 {\n\n\t\ti := a\n\t\tj := m\n\t\tfor i < j {\n\t\t\th := int(uint(i+j) >> 1)\n\t\t\tif !cmp.Less(data[m], data[h]) {\n\t\t\t\ti = h + 1\n\t\t\t} else {\n\t\t\t\tj = h\n\t\t\t}\n\t\t}\n\n\t\tfor k := m; k > i; k-- {\n\t\t\tdata[k], data[k-1] = data[k-1], data[k]\n\t\t}\n\t\treturn\n\t}\n\n\tmid := int(uint(a+b) >> 1)\n\tn := mid + m\n\tvar start, r int\n\tif m > mid {\n\t\tstart = n - b\n\t\tr = mid\n\t} else {\n\t\tstart = a\n\t\tr = m\n\t}\n\tp := n - 1\n\n\tfor start < r {\n\t\tc := int(uint(start+r) >> 1)\n\t\tif !cmp.Less(data[p-c], data[c]) {\n\t\t\tstart = c + 1\n\t\t} else {\n\t\t\tr = c\n\t\t}\n\t}\n\n\tend := n - start\n\tif start < m && m < end {\n\t\trotateOrdered(data, start, m, end)\n\t}\n\tif a < start && start < mid {\n\t\tsymMergeOrdered(data, a, start, mid)\n\t}\n\tif mid < end && end < b {\n\t\tsymMergeOrdered(data, mid, end, b)\n\t}\n}\n\nfunc rotateOrdered[E cmp.Ordered](data []E, a, m, b int) {\n\ti := m - a\n\tj := b - m\n\n\tfor i != j {\n\t\tif i > j {\n\t\t\tswapRangeOrdered(data, m-i, m, j)\n\t\t\ti -= j\n\t\t} else {\n\t\t\tswapRangeOrdered(data, m-i, m+j-i, i)\n\t\t\tj -= i\n\t\t}\n\t}\n\n\tswapRangeOrdered(data, m-i, m, i)\n}\n"