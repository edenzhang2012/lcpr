/*
 * @lc app=leetcode.cn id=6 lang=golang
 * @lcpr version=30204
 *
 * [6] Z 字形变换
 *
 * https://leetcode.cn/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (53.44%)
 * Likes:    2427
 * Dislikes: 0
 * Total Accepted:    747.4K
 * Total Submissions: 1.4M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
 * 输入：s = "PAYPALISHIRING", numRows = 4
 * 输出："PINALSIGYAHRPI"
 * 解释：
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * 示例 3：
 *
 * 输入：s = "A", numRows = 1
 * 输出："A"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1 <= numRows <= 1000
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end

package main

// @lc code=start
func convert(s string, numRows int) string {
	n := len(s)
	if n <= numRows || numRows == 1 {
		return s
	}

	/*
		方案1：模拟
		建立二维数组，将字符串的值按照题设方法依次填进数组，再从左到右扫描，跳过为空的部分
		两个难点：
		1. 如何确定二维数组的列数（go语言里利用slice可以规避这个问题）
			按照规律将数图切成 V 字形，从第1个顶点到第2个顶点之间（不含第二个顶点）的字母个数为
			r+r-2。N 字形可以拆成多个这样的 V 字形。每个 V 字形的的列数为 r-1。len(s)/(2r-2)
			可得到V字形个数，len(s)%(2r-2)得到余数，余数<=r 则总列数+1，否则总列数+(余数-r+1)
		2. 填充时如何确定数据对应在二维数组中的位置
			遍历数组，在对应位置填值

		1157/1157 cases passed (7 ms)
		Your runtime beats 57.58 % of golang submissions
		Your memory usage beats 41.64 % of golang submissions (8.6 MB)

		空间复杂度：二维数组+返回数组，返回数组不算空间占用。二维数组O(行数(numRows)*列数(x<n))
		时间复杂度：遍历一次字符串O(n),遍历一次二维数组O(n*r)
	*/

	// a := n / (2*numRows - 2)
	// b := n % (2*numRows - 2)
	// c := 0
	// if b > 0 && b <= numRows {
	// 	c = 1
	// } else if b > numRows {
	// 	c = b - numRows + 1
	// }

	// x := a*(numRows-1) + c

	// tmp := make([][]byte, numRows)
	// for i := 0; i < numRows; i++ { //初始化二维数组
	// 	tmp[i] = make([]byte, x)
	// }

	// //填充二维数组
	// for j, index := 0, 0; j < x; {
	// 	if j%(numRows-1) == 0 { //需要填充整列
	// 		for i := 0; i < numRows; i++ {
	// 			if index < n {
	// 				tmp[i][j] = s[index]
	// 				index++
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 		j++
	// 	} else {
	// 		for i := numRows - 2; i > 0; i-- { //台阶式填充
	// 			if index < n {
	// 				tmp[i][j] = s[index]
	// 				index++
	// 				j++
	// 			} else {
	// 				j++
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	// ret := make([]byte, n)
	// for i, index := 0, 0; i < numRows; i++ {
	// 	for j := 0; j < x; j++ {
	// 		if tmp[i][j] != 0 {
	// 			ret[index] = tmp[i][j]
	// 			index++
	// 		}
	// 	}
	// }

	// return string(ret)

	/*
		方案2：直接构建
		字符串在Z字形上的排列是有规则的，找到规律可以直接构建。跟解法1一样，可以发现周期
		T=2*numRows-2,每个周期内的元素下标都是有规律的。按行遍历（r）,每个周期（t）内
		第一个数据下标为 t*T+r；两个数据下标之和为 (2*t+1)*T,所以第二个数据下标为
		（t+1）*T-r。注意第二个数据的边界条件

		空间复杂度: 使用了一个辅助数组长度为n，O(n)
		时间复杂度：O(r*t)

		1157/1157 cases passed (0 ms)
		Your runtime beats 100 % of golang submissions
		Your memory usage beats 87.46 % of golang submissions (5.6 MB)
	*/

	T := 2*numRows - 2 //周期
	tMax := 0
	if n%T == 0 {
		tMax = n / T
	} else {
		tMax = n/T + 1
	}

	ret := make([]byte, n)
	index := 0
	for r := 0; r < numRows; r++ { //遍历行
		for t := 0; t < tMax; t++ { //遍历周期
			if t*T+r < n {
				ret[index] = s[t*T+r] //第一个数字
				index++
			}
			if r > 0 && r < numRows-1 && (t+1)*T-r < n {
				ret[index] = s[(t+1)*T-r] //第二个数字
				index++
			}
		}
	}

	return string(ret)
}

// @lc code=end

/*
// @lcpr case=start
// "PAYPALISHIRING"\n3\n
// @lcpr case=end

// @lcpr case=start
// "PAYPALISHIRING"\n4\n
// @lcpr case=end

// @lcpr case=start
// "A"\n1\n
// @lcpr case=end

*/
