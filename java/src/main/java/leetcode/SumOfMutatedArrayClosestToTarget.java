package leetcode;


import java.util.Arrays;

/**
 * https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/
 *
 * @author remark
 * @version 2020年06月14日01:40:39
 */
public class SumOfMutatedArrayClosestToTarget {

    private int getCloseOne(int ans , int target, int len){
        int m = target - ans * len;
        int n = (ans + 1) * len - target;

        if(m <= n){
            return ans;
        } else {
            return ans + 1;
        }
    }

    public int findBestValue(int[] arr, int target) {
        if(arr == null || arr.length == 0){
            return 0;
        }

        Arrays.sort(arr);

        int sum = 0 ;
        int[] sumArrs = new int[arr.length];

        int mid = target / arr.length;
        int midIndex = 0;

        for(int i = 0 ; i < arr.length ; i ++) {
            sum += arr[i];
            sumArrs[i] = sum;
            if(mid < arr[i]){
                midIndex = Math.max(0, i - 1);
                mid = Integer.MAX_VALUE;
            }
        }

        if(midIndex == arr.length - 1){
            return arr[midIndex];
        }

        if(midIndex == 0) {
            int gapValue = target / (arr.length);
            if(gapValue <= arr[0]){
                return getCloseOne(gapValue, target, arr.length);
            }
        }


        for(int i = midIndex ; i < arr.length - 1 ; i ++) {
            int now = arr[i];
            int next = arr[i + 1];
            int gapValue = (target - sumArrs[i]) / (arr.length - i - 1);



            if(gapValue >= now && gapValue <= next){
                return getCloseOne(gapValue, target - sumArrs[i], arr.length - midIndex - 1);
            }
        }
        return arr[arr.length - 1];
    }

    public static void main(String[] args){


        int[] data = new int[]{1,2,23,24,34,36};
        int target = 110;

        SumOfMutatedArrayClosestToTarget sumOfMutatedArrayClosestToTarget = new SumOfMutatedArrayClosestToTarget();
        System.out.println(sumOfMutatedArrayClosestToTarget.findBestValue(data, target));
    }
}
