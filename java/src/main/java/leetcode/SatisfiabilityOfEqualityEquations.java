package leetcode;


import com.alibaba.fastjson.JSON;

import java.util.*;

/**
 * https://leetcode-cn.com/problems/satisfiability-of-equality-equations/
 *
 * @author remark
 * @version 2020年06月08日00:05:05
 */
public class SatisfiabilityOfEqualityEquations {


    public int find(int[] parent , int index){
        while (parent[index] != index){
            parent[index] = parent[parent[index]];
            index = parent[index];
        }

        return index;
    }

    public void union(int[] parent, int index1, int index2) {
        parent[find(parent, index1)] = find(parent, index2);
    }



    public boolean equationsPossible(String[] equations) {

        if(equations == null || equations.length == 0){
            return true;
        }

        int[] parent = new int[26];

        for(int i = 0 ; i < parent.length ; i ++){
            parent[i] = i;
        }

        for(String equation : equations){

            int a = equation.charAt(0) - 'a';
            int b = equation.charAt(3) - 'a';
            char c = equation.charAt(1);

            if(c == '='){
                union(parent, a, b);
            }
        }

        for(String equation : equations){
            int a = equation.charAt(0) - 'a';
            int b = equation.charAt(3) - 'a';
            char c = equation.charAt(1);

            if(c == '!'){
               if(find(parent, a) == find(parent , b)){
                   return false;
               }

            }
        }

        return true;

    }

    public static void main(String[] args) {
        String[] eq = new String[]{"a==b", "e==c", "b==c", "a!=e"};

        String[] eq1 = new String[]{"a==b","b!=a"};

        String[] eq2 = new String[]{"a==b","b!=c","c==a"};

        SatisfiabilityOfEqualityEquations satisfiabilityOfEqualityEquations = new SatisfiabilityOfEqualityEquations();
        System.out.println(satisfiabilityOfEqualityEquations.equationsPossible(eq2));
    }
}
