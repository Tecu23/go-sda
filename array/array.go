// Common algorithms implemented on the standard Array Structure
package array

func LinearSearch(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func BinarySearch(arr []int, value int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == value {
			return true
		} else if arr[mid] > value {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return false
}

func SimpleSort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[i] > (*arr)[j] {
				tmp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
	}
}

func QuickSort(arr *[]int) {}

// TODO: Implement QuickSort
