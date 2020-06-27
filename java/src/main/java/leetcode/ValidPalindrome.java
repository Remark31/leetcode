package leetcode;


/**
 *
 * https://leetcode-cn.com/problems/valid-palindrome/
 *
 * @author remark
 * @version 2020年06月19日00:30:27
 *
 */
public class ValidPalindrome {

    public boolean isPalindrome(String s) {
        if(s == null || s.length() == 0){
            return true;
        }
        String s1 = s.trim().toLowerCase();

        String m = "";
        for(int i = 0 ; i < s1.length() ; i ++){
            char t = s1.charAt(i);
            if(t <= 'z' && t >= 'a' || t <= '9' && t >= '0'){
                m += String.valueOf(t);
            }
        }



        for(int i = 0 ; i < m.length() / 2 ; i ++){
            if(m.charAt(i) != m.charAt(m.length() - 1- i)){
                return false;
            }
        }
        return true;
    }
}
