package leetcode;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * leetcode 954题目
 * https://leetcode.com/problems/array-of-doubled-pairs/submissions/
 * @author remark
 * @version 2020年05月24日14:21:04
 */
public class ArrayofDoubledPairs {


    /*
     整体思路就是排序后寻找每个值的目标值，主要需要关注的是负数的问题，因为负数*2后会更小
     */
    public boolean canReorderDoubled(int[] A) {

        Map<Integer, Integer> datas = new HashMap<>();

        for(int item : A){
            datas.put(item, datas.getOrDefault(item, 0) + 1);
        }

        Arrays.sort(A);

        for(int i = 0 ; i < A.length; i ++){
            if(datas.get(A[i]) == 0){
                continue;
            }

            if(datas.get(A[i]) < 0){
                return false;
            }

            int target = A[i] * 2;
            if(A[i] < 0){
                if(A[i]%2 != 0){
                    return false;
                }
                target = A[i]/2;
            }
            if(datas.getOrDefault(target, 0) <= 0){
                return false;
            }
            datas.put(target, datas.get(target) - 1);
            datas.put(A[i], datas.get(A[i]) - 1);

        }

        return true;
    }

    public static void main(String[] args){
        int[] datas= new int[]{4,-2,2,-4};
        ArrayofDoubledPairs arrayofDoubledPairs = new ArrayofDoubledPairs();

        System.out.println(arrayofDoubledPairs.canReorderDoubled(datas));
    }
}
