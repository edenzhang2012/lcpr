/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30204
 *
 * [2] 两数相加
 *
 * https://leetcode.cn/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (44.57%)
 * Likes:    10896
 * Dislikes: 0
 * Total Accepted:    2.2M
 * Total Submissions: 5M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 *
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 *
 *
 * 示例 1：
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[7,0,8]
 * 解释：342 + 465 = 807.
 *
 *
 * 示例 2：
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 * 示例 3：
 *
 * 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * 输出：[8,9,9,9,0,0,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每个链表中的节点数在范围 [1, 100] 内
 * 0 <= Node.val <= 9
 * 题目数据保证列表表示的数字不含前导零
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		方案1：常规方案
		遍历链表，进行累加，注意进位，注意各种边界条件，注意两个链表一长一短
	*/

	ret := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := ret
	carry := 0
	for a, b := l1, l2; ; {
		if a == nil && b == nil {
			break
		}

		aVal, bVal := 0, 0
		if a != nil {
			aVal = a.Val
		}
		if b != nil {
			bVal = b.Val
		}

		if a != l1 {
			tmp := &ListNode{
				Val:  0,
				Next: nil,
			}

			current.Next = tmp
			current = current.Next
		}

		current.Val = (aVal+bVal)%10 + carry
		carry = (aVal + bVal) / 10
	}

	if carry > 0 {
		tmp := &ListNode{
			Val:  carry,
			Next: nil,
		}
		current.Next = tmp
		current = current.Next
	}

	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/

