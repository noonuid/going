package nc078

import (
	"github.com/anchorportal/go/nowcoder/structure"
)

type ListNode = structure.ListNode

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param pHead ListNodeš▒╗
 * @return ListNodeš▒╗
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var prev *ListNode
	for pHead != nil {
		next := pHead.Next
		pHead.Next = prev
		prev = pHead
		pHead = next
	}
	return prev
}
