package leetcode;

import leetcode.model.TreeNode;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class Codec {


    private int calDeepth(TreeNode root, Integer now, Integer max){
        if(root == null){
            return Math.max(now ,max);
        }


        return Math.max(calDeepth(root.left, now + 1, max),calDeepth(root.right, now + 1, max));
    }

    private void serializeTree(TreeNode node, Integer[] res , int index){

        if(res == null || res.length == 0){
            return ;
        }

        if(node == null) {
            return ;
        }

        res[index] = node.val;

        serializeTree(node.left, res, index * 2 + 1);
        serializeTree(node.right, res, index * 2 + 2);

    }


    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {

        if(root == null){
            return null;
        }

        // 计算树的深度
        int length = calDeepth(root, 0, 0);

        int sum = 1;
        for(int i = 0; i < length ; i ++ ){
            sum = sum * 2;
        }

        Integer[] res = new Integer[sum];

        serializeTree(root, res , 0);

        StringBuilder resSB = new StringBuilder();

        for(int i = 0 ; i < res.length - 1; i ++){

            if(res[i] != null){
                resSB.append(res[i].toString());
            } else {
                resSB.append("null");
            }


            resSB.append(",");
        }


        resSB.append(res[res.length - 1]);

        return resSB.toString();
    }

    private void MakeTree(Integer[] res , TreeNode node, int index){
        int leftIndex = index * 2 + 1;
        int rightIndex = index * 2 + 2;

        if(leftIndex >= res.length){
            node.left = null;
        } else{
            if(res[leftIndex] != null){
                node.left = new TreeNode(res[leftIndex]);
                MakeTree(res, node.left, leftIndex);
            }


        }

        if(rightIndex >= res.length){
            node.right = null;
        } else {
            if(res[rightIndex] != null){
                node.right = new TreeNode(res[rightIndex]);
                MakeTree(res, node.right, rightIndex);
            }
        }






    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if(data == null || data.length() == 0 || data.equals("")){
            return null;
        }

        String[] resArrs = data.split(",");



        Integer[] res = new Integer[resArrs.length];

        for(int i = 0 ; i < resArrs.length; i ++){

            if(resArrs[i].equalsIgnoreCase("null")){
                continue;
            }
            res[i] = Integer.valueOf(resArrs[i]);
        }

        TreeNode root = new TreeNode(res[0]);


        MakeTree(res, root, 0);


        return root;
    }





    public static void main(String[] args){
        TreeNode root = new TreeNode(1);



        root.right = new TreeNode(2);



//        TreeNode root = null;

        root.right.right = new TreeNode(3);
        root.right.right.right = new TreeNode(5);

        Codec codec = new Codec();
        System.out.println(codec.serialize(root));


       TreeNode newRT =  codec.deserialize(codec.serialize(root));
       if(newRT != null){
           newRT.print(newRT);
       }


    }

}
