package main

import (
	"fmt"
	"math"
	"strconv"
)

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var ret []int
	size1 := len(nums1)
	size2 := len(nums2)
	i := 0
	for ; i < size1 && i < size2; i++ {
		if nums1[i] < nums2[i] {
			ret = append(ret, nums1[i])
		} else {
			ret = append(ret, nums2[i])
		}
	}
	if i < size1 {
		ret = append(ret, nums1[i:]...)
	}
	if i < size2 {
		ret = append(ret, nums2[i:]...)
	}
	retSize := len(ret)
	if 0 == retSize {
		return 0.0
	}
	fmt.Printf("retsize:%d\n", retSize)
	fmt.Println(ret)
	mod := (retSize - 1) % 2
	mid := (retSize - 1) / 2
	if 0 == mod {
		return float64(ret[mid])
	}
	fmt.Println("hello")
	return float64(ret[mid]+ret[mid+1]) / float64(2)

}

func longestPalindrome(s string) string {
	size := len(s)
	if size <= 0 {
		return ""
	}
	retStr := ""
	for i := 0; i < size; i++ {
		hasDiff := false
		j := size - 1
		for ; j > i; j-- {
			//判断子串是否一致
			p := i
			q := j
			for p < q {
				if s[p] != s[q] {
					hasDiff = true
					break
				}
			}
		}
		if !hasDiff {
			if len(retStr) < len(s[i:j]) {
				retStr = s[i:j]
			}
		}
	}
	return retStr
}

func reverse(x int) int {
	str := fmt.Sprintf("%d", x)
	isMinus := false
	if x < 0 {
		//负数
		str = str[1:]
		isMinus = true
	}
	size := len(str)
	ret := make([]byte, size)
	for i := 0; i < size; i++ {
		ret[i] = str[size-1-i]
	}
	v, _ := strconv.ParseInt(string(ret), 10, 64)
	if isMinus {
		v = v * -1
		if v <= math.MinInt32 {
			return 0
		}
	} else {
		if v >= math.MaxInt32 {
			return 0
		}
	}
	return int(v)
}

func myAtoi(s string) int {
	size := len(s)
	if size <= 0 {
		return 0
	}
	index := 0
	for ; index < size; index++ {
		if ' ' != s[index] {
			break
		}
	}
	s = s[index:]
	if len(s) <= 0 {
		return 0
	}
	//fmt.Println(s)
	isMinus := false
	if '-' == s[0] {
		isMinus = true
		s = s[1:]
	} else if '+' == s[0] {
		isMinus = false
		s = s[1:]
	}
	//fmt.Println(s)
	size = len(s)
	if size <= 0 {
		return 0
	}
	ret := int(0)
	for i := 0; i < size; i++ {
		//fmt.Println(s[i])
		if '0' <= s[i] && s[i] <= '9' {
			ret = ret*10 + int(s[i]-'0')
			tmp := ret
			if isMinus {
				tmp *= -1
			}
			if tmp < math.MinInt32 {
				return math.MinInt32
			} else if tmp > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}
	}
	if isMinus {
		ret *= -1
	}
	return ret
}

func convert(s string, numRows int) string {
	size := len(s)
	if size < 2 || numRows < 2 {
		return s
	}
	tmp := make([][]rune, numRows)
	index := 0
	flag := -1
	for _, v := range s {
		tmp[index] = append(tmp[index], v)
		if 0 == index || index == numRows-1 {
			flag = -flag
		}
		index += flag
	}
	ret := ""
	for _, v := range tmp {
		ret += string(v)
	}
	return ret
}

var rome2Int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	size := len(s)
	if size <= 0 {
		return 0
	}
	sum := 0
	for i := 0; i < size; i++ {
		cur := s[i]
		val := rome2Int[cur]
		if i < size-1 && val < rome2Int[s[i+1]] {
			val *= -1
		}
		sum += val
	}
	return sum
}
