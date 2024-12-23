/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30204
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (38.88%)
 * Likes:    7461
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 4.8M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的 回文 子串。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// @lcpr-template-start

package main

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	/*
		解法1：
		思路：遍历字符串，中心点下标为i，此时子串长度为1，在这种情况下，向右查询与中心点相同的字符
		一直找到不同的字符为止，更新子串的前后下标，记录长度；然后再检查子串左右的元素是否相同，
		更新子串，并记录长度，直到遍历完毕
		TODO 这个思路应该没啥问题，但是实现上问题比较多，思路清晰，但是边界条件处理有问题
		这个解法就是官方解法中的中心扩展算法

		Your runtime beats 77.14 % of golang submissions
		Your memory usage beats 58.35 % of golang submissions (4.3 MB)
		耗时 0:21:9

		时间复杂度: O(n^2)
		空间复杂度: O(1)
	*/
	start := 0
	end := 0
	n := len(s)
	for i := 0; i < n; i++ {
		left1, right1 := expendProcess(s, i, i)
		left2, right2 := expendProcess(s, i, i+1)

		if right1-left1 > end-start {
			start = left1
			end = right1
		}

		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}

	return s[start : end+1]

	/*
		解法2：动态规划
		所谓动态规划，就是将复杂问题分解成重复的简单问题来解决。针对这道算法题。回文子串去除掉
		最左与最右的元素一定还是个回文子串，依次逐层向下分离最后就剩两种情况，最后剩下一个字符，
		那肯定是回文串，最后剩下2个字符，若这两个字符相等，则是回文串，否则不是。
		解决动态规划思想之后，还需要考虑如何进行遍历--将所有的长度都遍历一遍，每一种长度都需要
		从头遍历到尾
		还有另外一个问题，在golang里面如何方便快速的记录回文串并便于查找

		本地测试没问题，但是提交的时候每展示结果，可能当字符串太长的时候内存占用太大

		淦！golang可以用二维数组拼成vector

		时间复杂度：O(n^2)
		空间复杂度：O(n^2)

		142/142 cases passed (82 ms)
		Your runtime beats 18.22 % of golang submissions
		Your memory usage beats 10.14 % of golang submissions (8.7 MB)
		耗时 0:0:35
	*/
	// maxLen := 1
	// maxLenstart := 0
	// // records := make(map[string]bool, len(s))
	// records := make([][]bool, len(s))

	// //所有的单个字符都是回文串
	// for i := 0; i < len(s); i++ {
	// 	// records[fmt.Sprintf("%d,%d", i, i)] = true
	// 	records[i] = make([]bool, len(s))
	// 	records[i][i] = true
	// }

	// for l := 2; l <= len(s); l++ { //这里是长度，可以等于
	// 	for i := 0; i < len(s); i++ {
	// 		j := i + l - 1
	// 		if j >= len(s) {
	// 			break
	// 		}

	// 		if s[i] != s[j] {
	// 			// records[fmt.Sprintf("%d,%d", i, j)] = false
	// 			records[i][j] = false
	// 		} else {
	// 			if j-i < 3 { //优化，这种情况下一定是回文串，可以避免一次查找
	// 				// records[fmt.Sprintf("%d,%d", i, j)] = true
	// 				records[i][j] = true
	// 			} else {
	// 				records[i][j] = records[i+1][j-1]
	// 				// record = records[fmt.Sprintf("%d,%d", i+1, j-1)]
	// 				// records[fmt.Sprintf("%d,%d", i, j)] = record
	// 			}
	// 		}

	// 		// if record && j-i+1 > maxLen {
	// 		if records[i][j] && j-i+1 > maxLen {
	// 			maxLen = j - i + 1
	// 			maxLenstart = i
	// 		}
	// 	}
	// }

	// return s[maxLenstart : maxLenstart+maxLen]
}

func expendProcess(s string, i, j int) (int, int) {
	n := len(s)
	for ; i >= 0 && j < n && s[i] == s[j]; i, j = i-1, j+1 {
	}
	return i + 1, j - 1
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

// @lcpr case=start
// "vvgogaewginhopuxzwyryobjjpieyhwfopiowxmyylvcgsnhfcnogpqpukzmnpewavoutbloyrrhatimmxfqmcwgfebuoqbbgvubbkjfvxivjfijlpvtsnhagzfptahhyojwzamayoiegkenycnkxzkambimhdykdcxyyfjnvztzypmfczdhhnkmfuvgkhzbwmjznykkagqdrueohgcmeidjqsvfugcioeduohprjtfdmtzosnhoiganffarokxjifzzxhixdzycwfheqqegduzwtiacmdhqfmxhazcxsqyrvrihfqzjxvawdeandnwgjlquvzadruiqmcsgibglhicsvzqknztqpkiihdoisxipkourentfvrltieihiktgzswtgcmmlfrjifqinhrbplbsgswqlbjkyxjwoshsvxlhlpgzwbmxzwaemtowcxwourjwmmwjruowxcwotmeawzxmbwzgplhlxvshsowjxykjblqwsgsblpbrhniqfijrflmmcgtwszgtkihieitlrvftneruokpixsiodhiikpqtznkqzvscihlgbigscmqiurdazvuqljgwndnaedwavxjzqfhirvryqsxczahxmfqhdmcaitwzudgeqqehfwcyzdxihxzzfijxkoraffnagiohnsoztmdftjrphoudeoicgufvsqjdiemcghoeurdqgakkynzjmwbzhkgvufmknhhdzcfmpyztzvnjfyyxcdkydhmibmakzxkncynekgeioyamazwjoyhhatpfzgahnstvpljifjvixvfjkbbuvgbbqoubefgwcmqfxmmitahrryolbtuovawepnmzkupqpgoncfhnsgcvlyymxwoipofwhyeipjjboyrywzxupohnigweagogvv"\n
// @lcpr case=end

*/
