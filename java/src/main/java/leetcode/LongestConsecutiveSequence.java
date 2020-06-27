package leetcode;


import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

/**
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/
 *
 * @author remark
 * @version 2020年06月06日01:04:50
 */
public class LongestConsecutiveSequence {

    public int longestConsecutive(int[] nums) {
        if(nums == null || nums.length == 0){
            return 0;
        }



        Set<Integer> data = new HashSet<>();

        for(int i = 0 ; i < nums.length; i ++){
            data.add(nums[i]);
        }


        int maxLen = 0;

        for(int i = 0 ; i < nums.length ; i ++){
            int now = nums[i];
            if(!data.contains(now - 1)){
                int curNum = now;
                int curLen = 1;
                while (data.contains(curNum + 1)){
                    curNum += 1;
                    curLen += 1;
                }

               maxLen = Math.max(maxLen, curLen);
            }
        }

        return maxLen;
    }
}
