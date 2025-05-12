package main

import "fmt"

func quickSort (arr []int, low, high int){
		if low < high{
			pi := partition(arr, low, high)
			quickSort(arr, low, pi-1)
			quickSort(arr, pi + 1, high)
		}
}

func partition(arr []int, low, high int) int{
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++{
		if arr[j] < pivot{
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++{
		for j := 0; j < n-i-1; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++{
		minIndex := i
		for j := i+1; j < n; j++{
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func mergeSort(arr []int) []int{
	if len(arr) <= 1{
		return arr
	}
	mid := len(arr) / 2

	left := mergeSort(arr[:mid]) 
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int{
	final := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0{
		if left[0] <= right[0]{
			final = append(final, left[0])
			left = left[1:]
		} else {
			final = append(final, right[0])
			right = right[1:]
		}
	}

	final = append(final, left...)
	final = append(final, right...)
	return final
}

func main(){
	arr := []int{1, 4, 5, 3, 2, 90, 67, 23}
	// bubbleSort(arr)
	// selectSort(arr)
	// quickSort(arr, 0, len(arr)-1)
	arr = mergeSort(arr)
	fmt.Println(arr)
}