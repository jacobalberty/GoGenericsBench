//go:build go1.18
// +build go1.18

package main

import (
	"constraints"
	"testing"
)

func BenchmarkBubbleSortG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		bsortg(numar)
	}
}

func BenchmarkCountingSortG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		csortg(numar)
	}
}

func BenchmarkInsertSortG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		isortg(numar)
	}
}

func BenchmarkMergeSortG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		msortg(numar)
	}
}

func BenchmarkParallelMergeSortG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		pmsortg(numar)
	}
}

func BenchmarkSelectionSortG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		ssortg(numar)
	}
}

func bsortg[P constraints.Ordered](arr []P) []P {
	done := false
	for !done {
		done = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				done = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

func csortg[P constraints.Ordered](arr []P) []P {
	var cak []P
	ca := make(map[P]int)
	for _, v := range arr {
		if ca[v] == 0 {
			cak = append(cak, v)
		}
		ca[v]++
	}
	cak = bsortg(cak)
	var last *P
	last = nil
	for _, v := range cak {
		tmp := v
		if last != nil {
			ca[v] += ca[*last]
		}
		last = &tmp
	}
	res := make([]P, len(arr))

	for _, v := range arr {
		res[ca[v]-1] = v
		ca[v] -= 1
	}

	return res
}

func isortg[P constraints.Ordered](arr []P) []P {
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1; j = j - 1 {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}

	}
	return arr
}

func mergeg[P constraints.Ordered](first, second []P) []P {
	na := make([]P, len(first)+len(second))
	i := 0
	j := 0
	k := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			na[k] = first[i]
			i++
		} else {
			na[k] = second[j]
			j++
		}
		k++
	}

	for i < len(first) {
		na[k] = first[i]
		i++
		k++
	}
	for j < len(second) {
		na[k] = second[j]
		j++
		k++
	}

	return na
}

func pmsortg[P constraints.Ordered](arr []P) []P {
	if len(arr) == 1 {
		return arr
	}

	mp := len(arr) / 2
	var first []P
	var second []P
	wait := make(chan struct{})
	go func() {
		first = msortg(arr[:mp])
		close(wait)
	}()
	second = msortg(arr[mp:])
	<-wait

	return mergeg(first, second)
}

func msortg[P constraints.Ordered](arr []P) []P {
	if len(arr) == 1 {
		return arr
	}

	mp := len(arr) / 2
	var first []P
	var second []P
	first = msortg(arr[:mp])
	second = msortg(arr[mp:])

	return mergeg(first, second)
}

func ssortg[P constraints.Ordered](arr []P) []P {
	for i := 0; i < len(arr); i++ {
		k := i
		if i < len(arr) {
			for j := i; j < len(arr); j++ {
				if arr[k] > arr[j] {
					k = j
				}
			}
		}
		if k != i {
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
	return arr
}
