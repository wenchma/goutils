package algorithm

func merge(arr []int, low, mid, high int) {
	i, j, t := low, mid+1, 0
	temp := make([]int, high-low+1)
	for i <= mid && j <=high {
		if arr[i] < arr[j] {
			temp[t] = arr[i]
			i++
			t++
		} else {
			temp[t] = arr[j]
			j++
			t++
		}
	}
	for i <= mid {
		temp[t] = arr[i]
		i++
		t++
	}
	for j <= high {
		temp[t] = arr[j]
		j++
		t++
	}
	t = 0
	for low <= high {
		arr[low] = temp[t]
		low++
		t++
	}
}

func MergeSort(arr []int, low, high int) {
	if low < high {
		var mid int = (low + high)/2
		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
}