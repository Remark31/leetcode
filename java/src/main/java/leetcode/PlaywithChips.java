package leetcode;


import java.util.List;

/**
 *
 * https://leetcode.com/problems/play-with-chips/
 *
 * @author remark
 * @version 2020年05月30日23:34:05
 */
public class PlaywithChips {

    public int minCostToMoveChips(int[] chips) {
        int[] ans = new int[]{0,0};
        for(int i = 0 ; i < chips.length ; i ++){
            ans[chips[i]%2] ++;
        }

        return Math.min(ans[0], ans[1]);
    }

    public static void main(String[] args) {

    }
}
