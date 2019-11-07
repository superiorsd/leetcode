package main

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left + right) / 2
	l := mergeLists(lists, left, mid)
	r := mergeLists(lists, mid+1, right)

	return mergeTwoLists(l, r)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	sentinel := &ListNode{}
	cur := sentinel

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return sentinel.Next
}
