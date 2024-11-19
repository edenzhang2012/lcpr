/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30204
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (40.41%)
 * Likes:    10425
 * Dislikes: 0
 * Total Accepted:    3.1M
 * Total Submissions: 7.8M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	/*
		常规思路【也叫滑动窗口】：
		在主串上进行遍历，并记录下子串的起始与结束位置，子串的尾index在主串上递增，
		一直到尾index指定的值在子串中存在，记录下此时子串的长度。将子串的头index累加后继续，
		一直到子串的尾index等于主串的尾index
		注意：
		初始条件判断
		边界条件-index如何增加，何时退出循环等等
		空字符串处理
	*/
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	subStringHeadIndex := 0
	subStringTailIndex := subStringHeadIndex + 1
	maxLen := 1
	for ; subStringTailIndex < len(s); subStringTailIndex++ {
		if index := strings.IndexByte(s[subStringHeadIndex:subStringTailIndex],
			s[subStringTailIndex]); index >= 0 {
			if maxLen < subStringTailIndex-subStringHeadIndex {
				maxLen = subStringTailIndex - subStringHeadIndex
			}
			subStringHeadIndex++
			subStringTailIndex = subStringHeadIndex
		}
	}

	if maxLen < subStringTailIndex-subStringHeadIndex {
		maxLen = subStringTailIndex - subStringHeadIndex
	}

	return maxLen

	/*
		如何进一步提升时间与空间效率？
	*/
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
