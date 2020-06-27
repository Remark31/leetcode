package leetcode.model;

public class TreeNode {
    public int val;
    public TreeNode left;
    public TreeNode right;
    public TreeNode(int x){
        val = x;
    }


    public void print(TreeNode node){
        if(node == null){
            return ;
        }
        System.out.println(node.val);
        if(node.left != null){
            print(node.left);
        }

        if(node.right != null){
            print(node.right);
        }

    }


}
