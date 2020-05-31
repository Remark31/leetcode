package leetcode;

/**
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
 *
 * @author remark
 * @version 2020年05月30日23:55:59
 */
public class LargestRectangleInHistogram {

    public int largestRectangleArea(int[] heights) {
        int ans = 0;

        for(int i = 0 ; i < heights.length ; i ++){
            int minHigh = heights[i];

            for(int j = i ; j < heights.length ; j ++){
                minHigh = Math.min(minHigh, heights[j]);
                ans = Math.max(ans , minHigh * (j - i + 1));
            }
        }

        return ans;
    }

    public static void main(String[] args){

    }
}
