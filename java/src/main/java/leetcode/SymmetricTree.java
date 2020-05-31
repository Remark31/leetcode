package leetcode;


import leetcode.model.TreeNode;

/**
 * https://leetcode-cn.com/problems/symmetric-tree/
 *
 * @author remark
 * @version 2020年05月31日15:42:36
 */
public class SymmetricTree {

    public boolean isSymmetric(TreeNode root) {

        if(root == null){
            return true;
        }

        return symmetric(root.left, root.right);
    }

    private boolean symmetric(TreeNode a , TreeNode b){
        if(a == null && b == null){
            return true;
        }
        if(a == null || b == null){
            return false;
        }

        if(a.val == b.val){
            return symmetric(a.left, b.right) && symmetric(a.right , b.left);
        } else {
            return false;
        }
    }



    public static void main(String[] args){

    }
}
