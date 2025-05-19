package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int {1,1,3,2,2,4,4}
	//result := singleNumber(nums)
	//fmt.Println(result)

	// num := 12345
	// numstr :=strconv.Itoa(num)
	// numstr2 :=""
	// for _, v := range numstr {
	// 	numstr2 = string(v) + numstr2
	// 	fmt.Println(numstr2)
	// }
	// fmt.Println(numstr2)

	// result := isPalindrome(0)
	// fmt.Println(result)
	//result := isValid("({{[]}})")
	//fmt.Println(result)
	// result := longestCommonPrefix([]string{"dog","racecar","car"})
	// fmt.Println(result)

	// result := removeDuplicatesNew([]int{1, 1, 2})
	// fmt.Println(result)
	//result := plusOne([]int{1,2,9})
	//result := merge([][]int{{2,6},{1,3},{8,10},{15,18}})
	result := twoSum2([]int{2,7,11,15},26)
	fmt.Println(result)
	
}
// 使用map 实现两数之和 时间复杂度O(n)
func twoSum2(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, num := range nums {
		// 判断map中是否存在目标值=当前值,如果存在 返回两个数的下标
        if idx, exists  := hashMap[num]; exists  {
            return []int{idx, i}
        }
        hashMap[target-num] = i
    }
    return nil
}

//两数之和 时间复杂度O(n^2)
func twoSum(nums []int, target int) []int {
	// 循环数组
	for i := 0; i < len(nums); i++ {
		// 循环数组
		for j := i + 1; j < len(nums); j++ {
			// 如果两个数相加等于目标值
			if nums[i]+nums[j] == target {
				// 返回两个数的下标
				return []int{i, j}
			}
		}
	}
	// 如果没有找到 返回空数组
	return []int{}
}
//合并区间
func merge(intervals [][]int) [][]int {
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	// 合并区间
	var result [][]int
	for i := 0; i < len(intervals); i++ {
		// 如果结果集为空 或者 结果集的最后一个区间的右边界小于当前区间的左边界
		if len(result) == 0 || result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals[i])
		} else {
			// 否则 合并区间
			result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
		}
	}
	return result
}
//加一
func plusOne(digits []int) []int {
	//从切片末尾开始处理进位，遇到9则置零并继续前移，若全部为9则在数组头部插入1
	for i := len(digits)-1; i >-1 ; i-- {
		//不为9时,直接加1 且退出循环
		if digits[i]!=9 {
			digits[i]++
			return digits
		}
		//为9时 置零 
		digits[i] = 0

	}
	//如果全部为9 则新建切片 在切片头部插入1
	return append([]int{1}, digits...)

    //切片最后一位如果是9 怎新建一个切片,如果小于9则+1
	// if digits[len(digits)-1]==9{
	// 	digits[len(digits)-1]=1
	// 	digits = append(digits, 0)
	// }else{
	// 	digits[len(digits)-1]++
	// }
	//return  digits
}

//双指针
func removeDuplicatesNew(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
	//慢指针
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
		// 如果快指针不等于慢指针 慢指针加1 将快指针的值赋值给慢指针
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
	fmt.Println(nums)
    return slow + 1
}


// 删除有序数组中的重复项 新数组接收
func removeDuplicates(nums []int) int {
	// 如果数组长度为0 或者 数组只有一个元素 直接放回数组长度
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	//目标切片
	target := []int{}
	// 循环数组
	for i := 0; i < len(nums); {
		//取当前元素
		baseNum := nums[i]
		//判断第一个元素是否等于第二个元素，或者第三个元素，如果相等一直i++ 到i>=len(nums) 退出循环
		for i < len(nums)-1 && baseNum == nums[i+1] {
			i++
		}
		//未碰到相等的元素 下标自动加1
		i++
		// 将数组首元素元素添加到目标切片
		target = append(target, baseNum)
	}
	fmt.Println(target)
	return len(target)
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 判断数组是否为空
	if len(strs) == 0 {
		return ""
	}
	result := ""
	// 数据集合
	var runes [][]rune
	//循环字符串集合 获取字符串 塞入 runes
	for _, v := range strs {
		if len(v) == 0 {
			return ""
		}
		runes = append(runes, []rune(v))
	}
	//以第一个字符串为基准
	for i := 0; i < len(runes[0]); i++ {
		// 获取第一个字符串的第i个字符
		temp := runes[0][i]
		// 循环字符串集合
		for j := 0; j < len(runes); j++ {
			// 判断字符串集合的第j个字符串的第i个字符是否等于第一个字符串的第i个字符
			if i >= len(runes[j]) || runes[j][i] != temp {
				return result
			}
		}
		// 将第一个字符串的第i个字符添加到结果中
		result += string(temp)
	}

	return result
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 回文数
func isPalindrome(x int) bool {
	//排出负数  和 末尾数是0的
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	a := 0
	for x > a {
		a = a*10 + x%10
		x /= 10
	}
	// 前判断偶数位 后判断计数位
	return x == a || x == a/10
}

// 有效括号
func isValid(s string) bool {
	//判断括号是偶数
	if len(s)%2 != 0 {
		return false
	}
	//定一个类似栈的数组 append(stack,s[i]) 压栈 stack[:len(stack)-1] 出栈
	stack := make([]byte, 0)
	// 设置括号的对应关系 因为是右括号 所以是右括号为key
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		// 判断是否是左括号
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			//左括号压栈
			stack = append(stack, s[i])
			//右括号验证
		} else if len(stack) == 0 || m[s[i]] != stack[len(stack)-1] {
			return false
		} else {
			// 右括号匹配成功 出栈
			stack = stack[:len(stack)-1]
		}
	}
	//栈内无元素 则说明全部匹配
	return len(stack) == 0
}