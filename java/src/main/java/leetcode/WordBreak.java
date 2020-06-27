package leetcode;


import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/**
 *
 * https://leetcode-cn.com/problems/word-break/
 *
 * @author remark
 * @version 2020年06月25日01:09:13
 */
public class WordBreak {


    public static boolean wordBreak(String s, List<String> wordDict) {
        if(s == null || s.length() == 0){
            return true;
        }

        Set<String> wordSet = new HashSet<>(wordDict);

        boolean[] dp = new boolean[s.length() + 1];
        dp[0] = true;



        for(int i = 1 ; i <= s.length() ; i ++){
            for(int j = 0 ; j < i ; j ++){
                String now = s.substring(j , i);
                if(dp[j] && wordSet.contains(now)) {
                    dp[i] = true;
                    break;
                }
            }
        }

        for(int i = 0 ; i < dp.length ; i ++){
            System.out.println(dp[i]);
        }

        return dp[s.length()];
    }


    public static void main(String[] args){
        String data = "leetcode";
        List<String> b = new ArrayList<>();
        b.add("leet");
        b.add("code");

        System.out.println(wordBreak(data, b));
    }
}
