package leetcode;


/**
 * https://leetcode.com/problems/sum-of-square-numbers/
 */
public class SumofSquareNumbers {

    public boolean judgeSquareSum(int c) {
        if(c < 0){
            return false;
        }

        int num =  Double.valueOf(Math.sqrt(c)).intValue() + 1;
        for(int i = 0 ; i < num ; i ++){
            int d = c - i * i;
            int k =  Double.valueOf(Math.sqrt(d)).intValue();
            if(k * k == d){
                return true;
            }
        }
        return false;
    }

    public static void main(String[] args){

    }
}
