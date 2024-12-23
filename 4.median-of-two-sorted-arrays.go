/*
 * @lc app=leetcode.cn id=4 lang=golang
 * @lcpr version=30204
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (42.80%)
 * Likes:    7302
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.8M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*
		   解法1：最通俗的解法，先合并，再找中位数
		   空间复杂度:O(m+n)
		   时间复杂度:O(m+n)
		   BUG：
		   1. 嵌套遍历不对，因为不知道哪个数组先遍历完，改为单层遍历
		   2. golang数字精度问题，整数除浮点数会得到整数，需要在除之前就进行转换
		   结果：
		   Your runtime beats 100 % of golang submissions
			Your memory usage beats 6.39 % of golang submissions (6.6 MB)
	*/
	// total := make([]int, 0, len(nums1)+len(nums2))
	// i, j := 0, 0
	// for i < len(nums1) && j < len(nums2) {
	// 	if nums1[i] < nums2[j] {
	// 		total = append(total, nums1[i])
	// 		i++
	// 	} else {
	// 		total = append(total, nums2[j])
	// 		j++
	// 	}
	// }

	// if i != len(nums1) {
	// 	for ; i < len(nums1); i++ {
	// 		total = append(total, nums1[i])
	// 	}
	// }

	// if j != len(nums2) {
	// 	for ; j < len(nums2); j++ {
	// 		total = append(total, nums2[j])
	// 	}
	// }

	// if len(total)%2 == 0 {
	// 	return float64((total[len(total)/2] + total[len(total)/2-1])) / 2.0
	// } else {
	// 	return float64(total[len(total)/2])
	// }

	/*
		解法2：中位数，两个数组是有序的，两个数组长度都知道，中位数的位置也可以知道，不需要合并，
		直接找到对应位置的数，计算即可
		时间复杂度：O(M+N)
		空间复杂度：O(1)
		BUG：
		1. 没有考虑当其中一个数据被先遍历完的边界场景
		2. 粗心!!!
		结果：
		Your runtime beats 100 % of golang submissions
		Your memory usage beats 23.37 % of golang submissions (6.1 MB)

	*/
	// total := len(nums1) + len(nums2)
	// middle1 := total/2 - 1
	// middle2 := total / 2
	// middle1value, middle2Value, index := 0, 0, 0
	// i, j := 0, 0

	// for index <= middle2 && i < len(nums1) && j < len(nums2) {
	// 	if nums1[i] < nums2[j] {
	// 		if index == middle2 {
	// 			middle2Value = nums1[i]
	// 			break
	// 		}

	// 		if total%2 == 0 && index == middle1 {
	// 			middle1value = nums1[i]
	// 		}
	// 		i++
	// 	} else {
	// 		if index == middle2 {
	// 			middle2Value = nums2[j]
	// 			break
	// 		}

	// 		if total%2 == 0 && index == middle1 {
	// 			middle1value = nums2[j]
	// 		}
	// 		j++
	// 	}
	// 	index++
	// }

	// if i == len(nums1) {
	// 	for ; j < len(nums2); j++ {
	// 		if index == middle2 {
	// 			middle2Value = nums2[j]
	// 			break
	// 		}

	// 		if total%2 == 0 && index == middle1 {
	// 			middle1value = nums2[j]
	// 		}
	// 		index++
	// 	}
	// }

	// if j == len(nums2) {
	// 	for ; i < len(nums1); i++ {
	// 		if index == middle2 {
	// 			middle2Value = nums1[i]
	// 			break
	// 		}

	// 		if total%2 == 0 && index == middle1 {
	// 			middle1value = nums1[i]
	// 		}
	// 		index++
	// 	}
	// }

	// if total%2 == 0 {
	// 	return float64(middle1value+middle2Value) / 2.0
	// } else {
	// 	return float64(middle2Value)
	// }

	/*
		解法3：二分查找
		两个数组的中位数，总数为奇数时，中位数为第(m+n)/2+1个数；若总数为偶数，中位数为第
		(m+n)/2和第(m+n)/2+1个数的平均值，所以问题转化为寻找两个有序数组中的第 k 小的数，
		k=(m+n)/2+1或者k=(m+n)/2
		由于两个数组都有序，并且要求时间复杂度为O(log(n))，明显需要用到二分查找
		由上分析如何采取二分查找法。对于数组A，B，为了找到其第k大的数，可以比较A[k/2-1]和
		B[k/2-1],其前面分别有k/2-1个数，所以这两个数之前一共有k-2个数。数组是从小到大有序的，
		所以，若：
		A[k/2-1] <= B[k/2-1],则比A[k/2-1]小的数最多只有k-2个，不满足条件。同理在A[k/2-1]
		左边的数也全部可以排除
		A[k/2-1] > B[k/2-1],则比B[k/2-1]小的数最多只有k-2个，不满足条件。同理在B[k/2-1]
		左边的数也全部可以排除
		综上，通过不断地二分下去，即可找到对应的第k小的数
		时间复杂度：O(log(n))
		空间复杂度：O(1)
		结果：
		Your runtime beats 17.88 % of golang submissions
		Your memory usage beats 31.42 % of golang submissions (6 MB)
	*/
	total := len(nums1) + len(nums2)
	if total%2 != 0 {
		return float64(findKthSortedArrays(nums1, nums2, total/2+1))
	} else {
		a := findKthSortedArrays(nums1, nums2, total/2+1)
		b := findKthSortedArrays(nums1, nums2, total/2)
		return float64(a+b) / 2.0
	}

}

// 循环式写法
func findKthSortedArrays(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		//nums1已经遍历完，nums2还未遍历完，返回nums2中从index开始的第k个数
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		//nums2已经遍历完，nums1还未遍历完，返回nums1中从index开始的第k个数
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		//k==1是截止条件，表示要找第一个数，返回两个数组中的较小值即可
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			//排除了一部分值之后，k的值需要更新，减去排除掉数据的个数
			k -= (newIndex1 - index1 + 1)
			//排除了一部分值之后，index1的值也需要更新，从上一次index的下一位数字开始
			index1 = newIndex1 + 1
		} else {
			//同上
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

//TODO 递归式写法
// func findKthSortedArrays(nums1 []int, nums2 []int, k int) int {
// 	index1, index2 := 0, 0
// 	for {
// 		if index1 == len(nums1) {
// 			return nums2[index2+k-1]
// 		}
// 		if index2 == len(nums2) {
// 			return nums1[index1+k-1]
// 		}
// 		if k == 1 {
// 			return min(nums1[index1], nums2[index2])
// 		}
// 		half := k / 2
// 		newIndex1 := min(index1+half, len(nums1)) - 1
// 		newIndex2 := min(index2+half, len(nums2)) - 1
// 		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
// 		if pivot1 <= pivot2 {
// 			k -= (newIndex1 - index1 + 1)
// 			index1 = newIndex1 + 1
// 		} else {
// 			k -= (newIndex2 - index2 + 1)
// 			index2 = newIndex2 + 1
// 		}
// 	}
// }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3,4]\n
// @lcpr case=end

*/
