package leetcode;

public class LongestCommonPrefix {

    public String longestCommonPrefix(String[] strs) {

        if(strs == null || strs.length == 0){
            return "";
        }

        String res = "";

        int min = Integer.MAX_VALUE;
        for(int i = 0 ; i < strs.length ; i ++){
            if(min > strs[i].length()){
                min = strs[i].length();
            }
        }



        for(int i = 0 ; i < min ; i ++){
            char now = strs[0].charAt(i);
            boolean flag = true;
            for(int j = 0; j < strs.length ; j ++){
                if(now != strs[j].charAt(i)){
                    flag = false;
                    break;
                }
            }
            if(flag){
                res += now;
            } else {
                break;
            }
        }

        return res;
    }
}
