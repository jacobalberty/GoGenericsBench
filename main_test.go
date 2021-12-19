package main

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

var numAr *[]int

func init() {
	size := 500
	if ssize := os.Getenv("ARRAY_SIZE"); ssize != "" {
		if tmp, err := strconv.Atoi(ssize); err == nil {
			size = tmp
		}
	}
	r := rand.Perm(size)
	numAr = &r
}

func getNumAr() []int {
	return append([]int{}, *numAr...)
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		bsort(numar)
	}
}

func BenchmarkCountingSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		csort(numar)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		isort(numar)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		msort(numar)
	}
}

func BenchmarkParallelMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		pmsort(numar)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numar := getNumAr()
		ssort(numar)
	}
}

func bsort(arr []int) []int {
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

func csort(arr []int) []int {
	res := make([]int, len(arr))
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	ca := make([]int, max+1)
	for _, v := range arr {
		ca[v]++
	}

	for i := 1; i < len(ca); i++ {
		ca[i] += ca[i-1]
	}

	for _, v := range arr {
		res[ca[v]-1] = v
		ca[v]--
	}

	return res
}

func isort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1; j = j - 1 {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}

	}
	return arr
}

func merge(first, second []int) []int {
	na := make([]int, len(first)+len(second))
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

func msort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mp := len(arr) / 2
	var first []int
	var second []int
	first = msort(arr[:mp])
	second = msort(arr[mp:])

	return merge(first, second)
}

func pmsort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mp := len(arr) / 2
	var first []int
	var second []int
	wait := make(chan struct{})
	go func() {
		first = msort(arr[:mp])
		close(wait)
	}()
	second = msort(arr[mp:])
	<-wait

	return merge(first, second)
}

func ssort(arr []int) []int {
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
