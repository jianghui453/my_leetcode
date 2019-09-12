package linked_list

import "testing"

func TestPartition(t *testing.T) {
	var head, ret *ListNode
	var x int
	var hope, retSlice []int

	//head = &ListNode{
	//    1,
	//    &ListNode{
	//        4,
	//        &ListNode{
	//            3,
	//            &ListNode{
	//                2,
	//                &ListNode{
	//                    5,
	//                    &ListNode{
	//                        2,
	//                        nil,
	//                    },
	//                },
	//            },
	//        },
	//    },
	//}
	//x = 3
	//hope = []int{1, 2, 2, 4, 3, 5}
	//retSlice = make([]int, 0)
	//ret = partition(head, x)
	//for ret != nil {
	//    retSlice = append(retSlice, ret.Val)
	//    ret = ret.Next
	//}
	//t.Logf("\n    hope=%v \nretSlice=%v", hope, retSlice)

	head = &ListNode{
		1,
		nil,
	}
	x = 2
	hope = []int{1}
	retSlice = make([]int, 0)
	ret = partition(head, x)
	for ret != nil {
		retSlice = append(retSlice, ret.Val)
		ret = ret.Next
	}
	t.Logf("\n    hope=%v \nretSlice=%v", hope, retSlice)
}
