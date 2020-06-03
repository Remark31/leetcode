package leetcode;



public class ProductOfArrayExceptSelf {

    public int[] productExceptSelf(int[] nums) {
        if(nums.length == 1 || nums == null){
            return nums;
        }


        int before = 1;
        int sum = nums[nums.length - 1];

        int[] res = new int[nums.length];

        for(int i = nums.length - 1 ; i > 0; i --){
            res[i] = sum;
            sum = sum * nums[i - 1];
        }

        for(int i = 0 ; i < nums.length - 1 ; i ++){
            res[i] = res[i+1] * before;
            before = before * nums[i];
        }

        res[nums.length - 1] = before;

        return res;


    }

}
