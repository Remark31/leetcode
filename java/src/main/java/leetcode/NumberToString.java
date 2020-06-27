package leetcode;


import java.util.ArrayList;
import java.util.List;

/**
 * https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
 *
 * @author remark
 * @version 2020年06月09日00:04:13
 */
public class NumberToString {


    public int transNum(List<String> numlist , int nowIndex){
        if(nowIndex >= numlist.size()){
            return 1;
        }
        if(nowIndex == numlist.size() - 1){
            return 1;
        }


        int nowInt = Integer.valueOf(numlist.get(nowIndex));
        int preInt = Integer.valueOf(numlist.get(nowIndex + 1));

        if(nowInt + preInt * 10 < 26 && preInt != 0){
            return transNum(numlist, nowIndex + 1) + transNum(numlist, nowIndex + 2);
        } else {
            return transNum(numlist, nowIndex + 1);
        }

    }

    public int translateNum(int num) {
        int res = 0;

        if(num == 0){
            return 1;
        }

        List<String> numList = new ArrayList<>();

        while (num > 0) {
            int now = num % 10;
            num = num / 10;
            numList.add(String.valueOf(now));
        }


//        res = transNum(numList, 0);

        int[] dp = new int[numList.size() + 2];
        dp[0] = 1;
        dp[1] = 1;

        for(int i = 0; i < numList.size() - 1; i ++){
            int nowInt = Integer.valueOf(numList.get(i));
            int preInt = Integer.valueOf(numList.get(i + 1));

            if(nowInt + preInt * 10 < 26 && preInt != 0){
                dp[i + 2] = dp[i] + dp[i + 1];
            } else {
                dp[i + 2] = dp[i + 1];
            }


        }


        res = dp[numList.size()];

        return res;
    }

    public static void main(String[] args){
        int data = 12258;
        NumberToString numberToString = new NumberToString();
        System.out.println(numberToString.translateNum(data));
    }
}
