package leetcode;

public class MinimumSizeSubarraySum {

    public int minSubArrayLen(int s, int[] nums) {

        if(nums == null || nums.length == 0){
            return 0;
        }

        int min = 9999999;

        int left = 0;
        int sum = 0;
        for(int i = 0 ; i < nums.length ; i ++){
            sum += nums[i];
            while (sum >= s){
                min = Math.min(min , i - left + 1);
                sum -= nums[left];
                left ++;
            }
        }

        if(min == 9999999){
            return 0;
        }
        return min;
    }
}
