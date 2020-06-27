package leetcode;


/**
 * https://leetcode-cn.com/problems/first-missing-positive/solution/
 *
 * @author remark
 * @version 2020年06月27日15:39:05
 */
public class FirstMissingPositive {

    public int firstMissingPositive(int[] nums) {
        if(nums == null || nums.length == 0){
            return 0;
        }

        int n = nums.length;

        for(int i = 0 ; i < n ; i ++){
            if(nums[i] <= 0){
                nums[i] = n + 1;
            }
        }

        for(int i = 0 ; i < n ; i ++){
            int tmp = Math.abs(nums[i]);
            if(tmp < n && nums[tmp - 1] > 0){
                nums[tmp - 1] = 0 - nums[tmp - 1];
            }
        }



        for(int i = 0 ; i < n ; i ++){
            if(nums[i] > 0){
                return i + 1;
            }
        }
        return n;

    }

    public static void  main(String[] args){
        FirstMissingPositive firstMissingPositive = new FirstMissingPositive();

        int[] data = new int[]{0,1,2};

        System.out.println(firstMissingPositive.firstMissingPositive(data));
    }
}
