package leetcode;


import java.util.ArrayList;
import java.util.List;

/**
 * https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
 *
 * @author remark
 * @version 2020年06月05日00:09:08
 */
public class SpiralMatrix {

    private String LEFT = "left";
    private String RIGHT = "right";
    private String DOWN = "down";
    private String UP = "up";

    public int[] spiralOrder(int[][] matrix) {
        if(matrix == null || matrix.length == 0){
            return new int[]{};
        }


        int len = matrix.length * matrix[0].length;


        List<Integer> resList = new ArrayList();


        int column = 0;
        int row = 0;
        int endcolumn = matrix.length;
        int endrow = matrix[0].length;
        while (true) {

            // 右
            if(step(resList, matrix, column, row, endcolumn, endrow, RIGHT)){
                break;
            }
            column ++;
            // 下
            if(step(resList, matrix, column, row, endcolumn, endrow, DOWN)){
                break;
            }
            endrow --;
            // 左
            if(step(resList, matrix, column, row, endcolumn, endrow, LEFT)){
                break;
            }
            endcolumn --;

            // 上
            if(step(resList,matrix, column, row, endcolumn, endrow, UP)){
                break;
            }
            row ++;

        }


        int[] res = new int[resList.size()];

        for(int i = 0 ; i < resList.size(); i ++){
            res[i] = resList.get(i);
        }

        return res;
    }

    public boolean step(List<Integer> resList, int[][] matrix, int column, int row, int endcolumn, int endrow, String arrow){
        if(column == endcolumn || row == endrow){
            return false;
        }


        if(arrow.equals(RIGHT)) {
            for(int j = row ; j < endrow; j ++ ){
                resList.add(matrix[column][j]);
            }
            return true;
        }

        if(arrow.equals(DOWN)) {
            for(int i = column; i < endcolumn ; i ++){
                resList.add(matrix[i][endrow -1]);
            }
            return true;
        }

        if(arrow.equals(LEFT)){
            for(int j = endrow - 1; j >= row ; j --){
                resList.add(matrix[endcolumn - 1][j]);
            }
            return true;
        }

        if(arrow.equals(UP)){
            for(int i = endcolumn -1 ; i >= column ; i --){
                resList.add(matrix[i][row]);
            }
            return true;
        }

        return true;

    }



}
