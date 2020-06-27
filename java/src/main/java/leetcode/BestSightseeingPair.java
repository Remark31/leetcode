package leetcode;

import java.util.Stack;

public class BestSightseeingPair {
    public int maxScoreSightseeingPair(int[] A) {

        if(A == null || A.length == 0){
            return 0;
        }

        int[] dp = new int[A.length + 1];

        dp[0] = 0;

        int max = 0;

        for(int i = 1 ; i < A.length ; i ++){
            dp[i] = Math.max(dp[i - 1], A[i - 1] + i - 1);
            max = Math.max(max, dp[i] + A[i] - i);
        }


        return max;

    }
}
