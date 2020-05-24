package leetcode;


import java.util.ArrayList;
import java.util.List;

/**
 * leetcode:https://leetcode.com/problems/find-positive-integer-solution-for-a-given-equation/
 *
 * @author remark
 * @version 2020年05月24日15:30:29
 */
public class FindPositiveIntegerSolutionforaGivenEquation {

    class CustomFunction {
        // Returns f(x, y) for any given positive integers x and y.
        // Note that f(x, y) is increasing with respect to both x and y.
        // i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)


        public int f(int x, int y){
            return 0;
        };
    }

    public List<List<Integer>> findSolution(CustomFunction customfunction, int z) {
        if (z <= 0) {
            return new ArrayList<>();
        }
        List<List<Integer>> ans = new ArrayList<>();

        for(int i = 1; i <= z ; i ++){
            for(int j = 1; j <= z ; j ++){
                if(customfunction.f(i,j) == z){
                    List<Integer> tmp = new ArrayList<>();
                    tmp.add(i);
                    tmp.add(j);
                    ans.add(tmp);
                }
                if(customfunction.f(i,j) > z){
                    break;
                }
            }
        }
        return ans;
    }
}
