package main

import (
	"fmt"
	"math"
	"time"
)

func main123() {
	//s := "abcabcbb"
	//fmt.Println(lengthOfLongestSubstring(s))
	//
	//s = "bbbbb"
	//fmt.Println(lengthOfLongestSubstring(s))
	//
	//s = "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))

	//num1 := []int{1, 3}
	//num2 := []int{2}
	//fmt.Println(findMedianSortedArrays(num1, num2))
	//x := 1534236469
	//fmt.Println(reverse(x))
	//x := "-91283472332"
	//fmt.Println(myAtoi(x))

	x := "LVIII"
	fmt.Println(romanToInt(x))
}

func main() {
	fmt.Println("hello go world")

	arr := []int{
		43, 34, 656, 767, 323, 3, 545, 76, 11, 434, 776, 4343, 54, 432,
	}

	fmt.Printf("sort before:%v\n", arr)
	//arr = insertSort(arr)
	//arr = insertSortNew(arr)
	//arr = bubbleSort(arr)
	//arr = selectSort(arr)
	//arr = shellSort(arr)
	//arr = shellSortNew(arr)
	//arr = mergeSort(arr)
	//arr = quickSort(arr)
	//arr = quickSortNew(arr)
	//arr = heapSort(arr)
	//arr = countSort(arr)
	arr = bucketSort(arr, 5)
	//arr = radixSort(arr)

	fmt.Printf("sort after:%v\n", arr)

	fmt.Printf("%s success\n", time.Now().Format("2006-01-02 15:04:05"))
}

/*
*
基数排序
*/
func radixSort(arr []int) []int {
	size := len(arr)
	if size < 2 {
		return arr
	}
	maxBitNum := _maxBitNum(arr)
	dev := 1
	mod := 10
	for i := 0; i < maxBitNum; i++ {
		buckets := make([][]int, 10)
		result := make([]int, 0)
		for _, v := range arr {
			n := v / dev % mod
			buckets[n] = append(buckets[n], v)
		}
		dev *= 10
		for j := 0; j < 10; j++ {
			result = append(result, buckets[j]...)
		}
		for k := range arr {
			arr[k] = result[k]
		}
	}
	return arr
}

// 获取待排序数据的最大位数
func _maxBitNum(arr []int) int {
	ret := 1
	count := 10
	for i := 0; i < len(arr); i++ {
		for arr[i] > count {
			count *= 10
			ret++
		}
	}
	return ret
}

/*
*
桶排序
bucketSize:桶的个数
*/
func bucketSort(arr []int, bucketSize int) []int {
	size := len(arr)
	if size < 2 {
		return arr
	}
	//获取待排序数据的最大值与最小值
	minValue := arr[0]
	maxValue := arr[1]
	for i := 1; i < size; i++ {
		if arr[i] < minValue {
			minValue = arr[i]
		} else if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	//桶的个数，并初始化桶，bucketCount桶的大小
	bucketCount := (maxValue-minValue)/bucketSize + 1
	buckets := make([][]int, bucketSize)
	for i := 0; i < bucketSize; i++ {
		buckets[i] = make([]int, 0)
	}
	//把各个元素映射到各个桶中
	for i := 0; i < size; i++ {
		id := (arr[i] - minValue) / bucketCount
		buckets[id] = append(buckets[id], arr[i])
	}
	//对每个桶进行排序，然后按顺序取出桶中的数据放入arr返回结果中
	arrIndex := 0
	for i := 0; i < bucketSize; i++ {
		if 0 == len(buckets[i]) {
			continue
		}
		//直接使用现成的排序算法
		buckets[i] = insertSortNew(buckets[i])
		for j := 0; j < len(buckets[i]); j++ {
			arr[arrIndex] = buckets[i][j]
			arrIndex++
		}
	}
	return arr
}

/*
*
计数排序
*/
func countSort(arr []int) []int {
	size := len(arr)
	if size < 2 {
		return arr
	}
	//获取待排序数列中的最大值
	maxValue := _getMaxValue(arr)
	bucketLen := maxValue + 1
	//构建一个统计数组，索引为待排序的元素，value为元素出现的个数
	bucket := make([]int, bucketLen)
	sortIndex := 0
	for i := 0; i < size; i++ {
		bucket[arr[i]] += 1
	}
	//重新遍历统计数组，索引为已经排好序的元素，value为该元素出现的个数
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 { //判断是否已经会写完所有的元素
			arr[sortIndex] = j
			sortIndex++
			bucket[j]--
		}
	}
	return arr
}
func _getMaxValue(arr []int) int {
	largest := math.MinInt32
	smallest := math.MaxInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}
	maxValue := largest // - smallest
	return maxValue
}

/*
*
堆排序
*/
func heapSort(arr []int) []int {
	size := len(arr)
	if size < 2 {
		return arr
	}
	//构建最大堆，最大堆构建完成后，第0个元素就是待排序数组中的最大元素
	_buildMaxHeap(arr, size)
	for i := size - 1; i >= 0; i-- {
		//每次都进行第0号（最大），第n-1号（元素）元素进行交换
		_swap(arr, 0, i)
		//交换后，把剩下的堆0~（n-1-i）进行调整，（由于最后一个元素已经有序，所以size总数需要减1个）
		size -= 1
		_heapify(arr, 0, size)
	}
	return arr
}
func _buildMaxHeap(arr []int, size int) {
	for i := size / 2; i >= 0; i-- {
		_heapify(arr, i, size)
	}
}

// 堆化函数
func _heapify(arr []int, i, size int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < size && arr[left] > arr[largest] {
		largest = left
	}
	if right < size && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		_swap(arr, i, largest)
		_heapify(arr, largest, size)
	}
}

/*
*
快速排序（另一种实现）
*/
func quickSortNew(arr []int) []int {
	size := len(arr)
	if size < 2 {
		return arr
	}
	//根据基准数据，把数据分成左右两个子序列
	pivot := arr[0]
	var left, right []int
	var result []int
	for i := 1; i < size; i++ {
		if pivot > arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	//对左边子序列进行排序
	left = quickSortNew(left)
	//对右边子序列进行排序
	right = quickSortNew(right)
	//组合left、pivot、right
	result = append(result, left...)
	result = append(result, pivot)
	result = append(result, right...)
	return result
}

/*
*
快速排序
*/
func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}
func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := _partition1Way(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

// 快速排序--单路
func _partition1Way(arr []int, left, right int) int {
	//先分区，最后把基准换到边界上
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			//当前值小于基准就交换
			_swap(arr, i, index)
			index += 1 //交换后的下一个
		}
	}
	_swap(arr, pivot, index-1)
	return index - 1
}
func _partition2Way(arr []int, low int, high int) int {
	pivot := arr[low]
	for low < high {
		//当队尾的元素大于等于基准数据时，向前移动high指针
		for low < high && arr[high] >= pivot {
			high--
		}
		//如果队尾元素小于pivot了，需要将其赋值给low
		arr[low] = arr[high]
		//当队首元素小于等于pivot时，向后移动low指针
		for low < high && arr[low] <= pivot {
			low++
		}
		//当队首元素大于pivot时，需要将其赋值给high
		arr[high] = arr[low]
	}
	//跳出循环时low和high相等，此时的low或high就是pivot的正确索引位置
	arr[low] = pivot
	return low
}
func _swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
*
归并排序
*/
func mergeSort(arr []int) []int {
	size := len(arr)
	if size < 2 {
		return arr
	}
	//首先数组进行按中线位置分成2份
	mid := size / 2
	left := arr[0:mid]
	right := arr[mid:]
	//对左右两份数据进行归并排序，最后把结果进行有序合并
	return _merge(mergeSort(left), mergeSort(right))
}

// 按有序合并两个排序数组
func _merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if 0 == len(left) {
		result = append(result, right...)
	}
	//for len(left) != 0 {
	//	result = append(result, left[0])
	//	left = left[1:]
	//}
	if 0 == len(right) {
		result = append(result, left...)
	}
	//for len(right) != 0 {
	//	result = append(result, right[0])
	//	right = right[1:]
	//}
	return result
}

/*
*
希尔排序
*/
func shellSort(arr []int) []int {
	size := len(arr)
	if size <= 0 {
		return arr
	}
	//初始设置一个分组大小gap（一般为原数组长度的一半），然后不断缩小gap值，直至gap=1
	for gap := size / 2; gap >= 1; gap = gap / 2 {
		//以下就是一个插入排序，当gap=1时，就是原始的插入排序
		for i := gap; i < size; i++ {
			tmp := arr[i]
			j := i - gap
			for ; j >= 0 && tmp < arr[j]; j = j - gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = tmp
		}
	}
	return arr
}
func shellSortNew(arr []int) []int {
	size := len(arr)
	if size <= 0 {
		return arr
	}
	gap := 1
	for gap < size/3 {
		gap = gap*3 + 1
	}
	for ; gap > 0; gap = gap / 3 {
		for i := gap; i < size; i++ {
			tmp := arr[i]
			j := i - gap
			for ; j >= 0 && arr[j] > tmp; j = j - gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = tmp
		}
	}
	return arr
}

/*
*
选择排序
*/
func selectSort(arr []int) []int {
	size := len(arr)
	if size <= 0 {
		return arr
	}
	for i := 0; i < size; i++ {
		minIndex := i                   //默认以循环的索引位置为最小的值
		for j := i + 1; j < size; j++ { //寻找后续数据中的最小值，不需要每次从0开始，需要根据当前循环的位置来判断需要在哪些数据列表中找最小值，已经循环的不需要再找了
			//找最小的值索引
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

/*
*
冒泡排序
*/
func bubbleSort(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}
	size := len(arr)
	for i := 0; i < size; i++ {
		//每次经过一次排序，最大的会落到末尾去，所以下一次相邻两个数比较的时候，最后的几位不需要在进行比较了
		for j := 0; j < size-1-i; j++ {
			if arr[j] > arr[j+1] {
				//如果前一个数比后一个数大，交换位置（大的数放到后面）
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

/*
*
插入排序
*/
func insertSort(list []int) []int {
	if len(list) <= 0 {
		return list
	}
	size := len(list)
	for i := 0; i < size-1; i++ {
		end := i
		tmp := list[end+1]
		for end >= 0 {
			if tmp < list[end] {
				list[end+1] = list[end]
				end--
			} else {
				break
			}
		}
		list[end+1] = tmp
	}
	return list
}

/*
*
插入排序
*/
func insertSortNew(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}
	for i, v := range arr {
		sortIndex := i - 1
		//插入排序，取出的某个位置的元素，去已经排好序的列表（0~sortIndex）里，倒序比较，如果比当前位置的值小，已经排序好的列表往后面挪一个位置（空一个位置出来，插入新的值）
		for ; 0 <= sortIndex && arr[sortIndex] > v; sortIndex-- {
			arr[sortIndex+1] = arr[sortIndex]
		}
		//v插入到比v小的sortIndex的下一个位置sortIndex+1
		arr[sortIndex+1] = v
	}
	return arr
}
