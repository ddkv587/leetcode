/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
package addTwoNumbers

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func init() {
    fmt.Println( "addTwoNumbers init" )
}

func AddTwoNumbers( l1 *ListNode, l2 *ListNode ) *ListNode {
    var head *ListNode  = nil
    ptr := head
    
    bCarry  := 0
    add1    := 0
    add2    := 0
    for ; l1 != nil || l2 != nil ; {
        if l1 == nil {
            add1 = 0
        } else {
            add1 = l1.Val
            l1   = l1.Next
        }

        if l2 == nil {
            add2 = 0
        } else {
            add2 = l2.Val
            l2   = l2.Next
        }

        add1    = add1 + add2 + bCarry
        bCarry  = 0
        
        if add1 >= 10 {
            bCarry = 1
            add1 = add1 - 10
        }
        
        tmp := ListNode{ Val: add1, Next: nil }
        if ptr != nil {
            ptr.Next = &tmp
            ptr = ptr.Next
        } else {
            ptr     = &tmp
            head    = ptr
        }
    }

    if bCarry == 1 {
        tmp := ListNode{ Val: bCarry, Next: nil }
        ptr.Next = &tmp
    }

    return head
}