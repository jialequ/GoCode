package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    // 思路，因为数组有序，所以可以从最后一个往前遍历，将大的数放在num1的最后一位
    // 定义两个下标来存储现在的位置
    i, i1, i2:= m + n - 1, m - 1, n - 1
    for i2 >= 0 {
        if i1 >= 0 && nums1[i1] > nums2[i2] {
            // 如果nums1>nums2，nums1和最后一个交换
            // nums1[i], nums1[i1] = nums1[i1], nums1[i]
			nums1[i] = nums1[i1]
            i1--
		} else {
            // 如果nums1<nums2，nums2写入最后一个
            nums1[i] = nums2[i2]
            i2--
        }
		i--
    }
}

func main() {
	nums1 := []int{4,5,6,0,0,0}
	nums2 := []int{1,2,3}
	merge(nums1, len(nums1) - len(nums2), nums2, len(nums2))
	fmt.Println(nums1)
}