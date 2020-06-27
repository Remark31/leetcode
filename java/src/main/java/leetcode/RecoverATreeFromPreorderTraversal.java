package leetcode;

import leetcode.model.TreeNode;

import java.util.Stack;

public class RecoverATreeFromPreorderTraversal {





    public TreeNode recoverFromPreorder(String S) {

        if(S.length() == 0){
            return null;
        }

        String[] m = S.split("-");

        Stack<TreeNode> stack = new Stack<>();
        Stack<Integer> depth = new Stack<>();

        TreeNode root = new TreeNode(Integer.valueOf(m[0]));
        depth.push(0);
        stack.push(root);

        int index = 1;
        while (index < m.length){
            int nowD = 1;
            while (index < m.length && m[index].equals("")){
                nowD ++;
                index ++;
            }

            if(index == m.length){
                break;
            }
            int preD = depth.peek();

            while (nowD != (preD + 1)){
                stack.pop();
                depth.pop();
                preD = depth.peek();
            }

            if (nowD == preD + 1){
                TreeNode pre = stack.peek();
                TreeNode now = new TreeNode(Integer.valueOf(m[index]));

                if(pre.left == null) {
                    pre.left = now;
                    stack.push(now);
                    depth.push(nowD);
                } else {
                    pre.right = now;
                    stack.push(now);
                    depth.push(nowD);
                }

                index ++;

            }
        }



        return root;
    }

    public static void main(String[] args){
        String data = "1-2--3---4-5--6---7";

        RecoverATreeFromPreorderTraversal recoverATreeFromPreorderTraversal = new RecoverATreeFromPreorderTraversal();

        recoverATreeFromPreorderTraversal.recoverFromPreorder(data);
    }
}
