package leetcode;


import java.util.HashMap;
import java.util.Map;

/**
 * https://leetcode-cn.com/problems/maximum-equal-frequency/
 *
 * @author remark
 * @version 2020年05月31日22:31:44
 */
public class MaximumEqualFrequency {


    // 1. max, submax ,submax,submax
    // 2. max, max, 1 , max
    // 3. 1, 1, 1, 1, 1
    public static int maxEqualFreq(int[] nums) {
        Map<Integer,Integer> feq = new HashMap<>();

        if(nums.length == 0){
            return 0;
        }

        int res = 1;
        int max = 0;
        int maxCount = 0;
        int oneCount = 0;
        int maxSubCount = 0;

        for(int i = 0 ; i < nums.length; i ++) {
            int now = nums[i];

            Integer count = feq.get(now);
            if(count == null){
                count = 0;
            }
            feq.put(now, count + 1);

            int nowFeq = feq.get(now);

            // 1
            if(nowFeq == 1){
                oneCount ++;
            } else if(nowFeq == 2){
                oneCount --;
            }

            // submax
            if(max >= 2){
                if(nowFeq == max - 1){
                    maxSubCount ++;
                } else if(nowFeq == max){
                    maxSubCount --;
                }
            }

            // max
            if(nowFeq > max){
                max = nowFeq;
                maxSubCount = maxCount > 0 ? maxCount - 1: 0;
                maxCount = 1;

            } else if (nowFeq == max){
                maxCount ++;
            }




            // 满足条件1
            // 满足条件2
            // 满足条件3
            int size = feq.size();
            if(oneCount == size || maxCount == 1 && maxSubCount == size - 1
                    || oneCount == 1 && maxCount == size - 1 || oneCount == size - 1 && maxCount == 1){
                res = i + 1;
            }



        }

        return res;

    }


    public static void main(String[] args){
        int[] data = new int[]{10,2,8,9,3,8,1,5,2,3,7,6};

        System.out.println(maxEqualFreq(data));
    }
}
