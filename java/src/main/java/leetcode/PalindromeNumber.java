package leetcode;


import java.util.List;

/**
 * https://leetcode-cn.com/problems/palindrome-number/
 *
 * @author remark
 * @version 2020年06月10日00:35:18
 */
public class PalindromeNumber {
    public boolean isPalindrome(int x) {
        if(x < 0){
            return false;
        }

        if( x == 0){
            return true;
        }

        String tn = String.valueOf(x);
        int size = tn.length()/2;
        for(int i = 0 ; i <= size ; i ++){
            if(tn.charAt(i) != tn.charAt(tn.length() - i - 1)){
                return false;
            }
        }
        return true;
    }
}
