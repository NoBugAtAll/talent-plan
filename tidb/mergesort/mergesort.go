package main

import (
	"sync"
)

const threshold = 1600
// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	temp := make([]int64, len(src))		// auxiliary array
	copy(temp, src)
	//fmt.Println(temp)
	Sort(0, int64(len(src) - 1), temp, src)
}

func Sort(left, right int64, src, temp []int64) {
	mid := (left + right) / 2
	if right <= left {
		return
	} else if right - left <= threshold {
		// mid := (left + right) / 2
		Sort(left, mid, temp, src)
		Sort(mid + 1, right, temp, src)
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			Sort(left, mid, temp, src)
		}()
		go func() {
			defer wg.Done()
			Sort(mid + 1, right, temp, src)
		}()
		wg.Wait()
	}

	Merge(left, mid, right, src, temp)
	return
}

func Merge(left, mid, right int64, src, temp []int64) {
	if src[mid] <= src[mid + 1] {
		copy(temp[left:right + 1], src[left:right + 1])
		return
	}
	i := left
	j := mid + 1
	k := left
	// temp := make([]int64, right - left + 1)
	for i <= mid && j <= right {
		if src[i] < src[j] {
			temp[k] = src[i]
			i++
			k++
		} else {
			temp[k] = src[j]
			j++
			k++
		}
	}
	for i <= mid {
		temp[k] = src[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = src[j]
		j++
		k++
	}
	return
}