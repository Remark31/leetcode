package leetcode;

import java.util.Arrays;


/**
 * https://leetcode-cn.com/problems/3sum-closest/submissions/
 *
 * @author remark
 * @version 2020年06月25日01:34:22
 */
public class _3sumClosest {

    private static int closet(int a , int b){
        if(Math.abs(a) > Math.abs(b)){
            return b;
        }
        return a;
    }

    public static int threeSumClosest(int[] nums, int target) {
        if(nums == null || nums.length == 0){
            return 0;
        }

        Arrays.sort(nums);


        for(int i = 0 ; i < nums.length; i ++){
            System.out.print(nums[i] + ",");
        }
        System.out.println();

        int minGap = 9999999;

        for(int i = 0 ; i < nums.length - 2 ; i ++){


            for(int j = i + 1; j < nums.length - 1; j ++){

                for(int k = j + 1; k < nums.length ; k ++){

                    int gap = target - (nums[i] + nums[j] + nums[k]);
                    System.out.println(gap);
                    minGap = closet(gap, minGap);

                    if(nums[i] + nums[j] + nums[k] > target){
                        break;
                    }
                }
                if(nums[j + 1] > 0 && nums[i] + nums[j] > target){
                    break;
                }
            }

            if(nums[i + 1] > 0 && nums[i] > target){
                break;
            }
        }
        return target - minGap;
    }

    public static void main(String[] args){
        int[] nums = new int[]{-100,-98,-2,-1};
        int target = -101;
        System.out.println(threeSumClosest(nums ,target));
    }
}
