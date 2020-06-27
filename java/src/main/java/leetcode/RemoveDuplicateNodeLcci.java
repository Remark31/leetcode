package leetcode;


import leetcode.model.ListNode;

import java.util.HashSet;
import java.util.Set;

/**
 *
 * https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
 *
 * @author remark
 * @version 2020年06月26日01:48:14
 */
public class RemoveDuplicateNodeLcci {

    public ListNode removeDuplicateNodes(ListNode head) {
        Set<Integer> data = new HashSet<>();

        if(head == null){
            return head;
        }

        ListNode now = head.next;
        ListNode pre = head;
        data.add(head.val);


        while (now != null){
            if(data.contains(now.val)){
                pre.next = now.next;
                now = now.next;
            } else {
                data.add(now.val);
                now = now.next;
                pre = pre.next;
            }

        }

        return head;
    }
}
