package _7_sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_mergeSort(t *testing.T) {
	const n = 10000
	var arr []int
	arr = make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	MergeSort(arr)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}