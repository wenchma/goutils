package algorithm

func QuickSort(arr []int, low int, high int) {
	var v, i, j int = arr[low], low, high
	if (low >= high) {
		return
	}
	for i < j {
		for i < j && arr[j] >= v {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= v {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = v
	QuickSort(arr, low, i-1)
	QuickSort(arr, i+1, high)
}
