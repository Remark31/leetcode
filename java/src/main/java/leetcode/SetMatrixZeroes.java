package leetcode;


import javafx.util.Pair;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * https://leetcode-cn.com/problems/set-matrix-zeroes/
 *
 * @author remark
 * @version 2020年06月01日23:06:42
 */
public class SetMatrixZeroes {

    public void setZeroes(int[][] matrix) {



        for(int i = 1 ; i < matrix.length ; i ++) {
            for(int j = 1 ; j < matrix[i].length ; j ++){
                if(matrix[i][j] == 0){
                    matrix[i][0] = Integer.MIN_VALUE;
                    matrix[0][i] = Integer.MIN_VALUE;
                }
            }
        }

        for(int i = 1; i < matrix.length ; i ++){
            for(int j = 1; j < matrix[i].length ; j ++){
                if(matrix[i][0] == Integer.MIN_VALUE || matrix[0][i] == Integer.MIN_VALUE){
                    matrix[i][j] = 0;
                }
            }
        }

        boolean rowsFlag = false;
        for(int i = 0 ; i < matrix.length; i ++){
            if(matrix[i][0] == 0) {
                rowsFlag = true;
            }
        }

        boolean columnFlag = false;
        for(int j = 0 ; j < matrix[0].length ; j ++){
            if(matrix[0][j] == 0){
                columnFlag = true;
            }
        }

        if(rowsFlag){
            for(int i = 0 ; i < matrix.length; i ++){
               matrix[i][0] = 0;
            }
        } else{
            for(int i = 0 ; i < matrix.length; i ++){
                if(matrix[i][0] == Integer.MIN_VALUE){
                    matrix[i][0] = 0;
                }
            }
        }

        if(columnFlag){
            for(int j = 0 ; j < matrix[0].length; j ++){
                matrix[0][j] = 0;
            }
        } else{
            for(int j = 0 ; j < matrix[0].length; j ++){
                if(matrix[0][j] == Integer.MIN_VALUE){
                    matrix[0][j] = 0;
                }
            }
        }


    }
}
