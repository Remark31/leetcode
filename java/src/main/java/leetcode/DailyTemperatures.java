package leetcode;


import java.util.Stack;

/**
 * https://leetcode-cn.com/problems/daily-temperatures/submissions/
 *
 * @author remark
 * @version 2020年06月11日00:43:54
 */
public class DailyTemperatures {

    public int[] dailyTemperatures(int[] T) {

        int length = T.length;
        int[] ans = new int[length];
        Stack<Integer> stack = new Stack<>();
        for (int i = 0; i < length; i++) {
            int temperature = T[i];
            while (!stack.isEmpty() && temperature > T[stack.peek()]) {
                int prevIndex = stack.pop();
                ans[prevIndex] = i - prevIndex;
            }
            stack.push(i);
        }
        return ans;

    }
}
