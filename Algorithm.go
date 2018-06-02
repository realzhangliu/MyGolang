package main

var a chan string

func MergeSort(a []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSort(a, low, mid)
		MergeSort(a, mid+1, high)
		merge(a, low, mid, high)
	}
}
func merge(a []int, low int, mid int, high int) {
	N := high - low + 1
	var b = make([]int, N)
	left := low
	right := mid + 1
	bIndex := 0
	for left <= mid && right <= high {
		if a[left] <= a[right] {
			b[bIndex] = a[left]
			bIndex++
			left++
		} else {
			b[bIndex] = a[right]
			bIndex++
			right++
		}
	}
	for left <= mid {
		b[bIndex] = a[left]
		bIndex++
		left++
	}
	for right <= high {
		b[bIndex] = a[right]
		bIndex++
		right++
	}
	for i := 0; i < N; i++ {
		a[i+low] = b[i]
	}
}

func QuickSort(a []int, low, high int) {
	if low < high {
		pivot := partition(a, low, high)
		QuickSort(a, low, pivot)
		QuickSort(a, pivot+1, high)
	}

}

func partition(a []int, low int, high int) int {
	m := low
	p := a[low]
	for k := low + 1; k <= high; k++ {
		if a[k] < p {
			m++
			a[k], a[m] = a[m], a[k]
		}
	}
	a[low], a[m] = a[m], a[low]
	return m
}
