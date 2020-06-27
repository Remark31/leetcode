package leetcode;

import java.util.*;

public class _3SUM {
    private boolean checkRepeat(List<List<Integer>> res, List<Integer> tmp){
        boolean flag = false;
        for(List<Integer> re : res){
            for(int i = 0 ; i < 3 ; i ++){
                if(re.get(i) == tmp.get(i)){
                    flag = true;
                } else {
                    flag = false;
                    break;
                }
            }
            if(flag){
                return true;
            }
        }
        return false;
    }

    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();

        Arrays.sort(nums);
        Map<Integer,Integer> numMap = new HashMap<>();

        for(int i = 0 ; i < nums.length; i ++){
            int time = numMap.getOrDefault(nums[i], 0);
            numMap.put(nums[i], time + 1);
        }


        for(int i = 0 ; i < nums.length; i ++){

            for(int j = i + 1 ; j < nums.length ; j ++){
                int iTime = numMap.get(nums[i]);
                numMap.put(nums[i], iTime - 1);
                int jTime = numMap.get(nums[j]);
                numMap.put(nums[j], jTime - 1);

                int n = 0 - (nums[i] + nums[j]);

                if(n < nums[j]) {
                    numMap.put(nums[j], jTime);
                    numMap.put(nums[i], iTime);
                    continue;
                }

                if(numMap.get(n) != null && numMap.get(n) > 0){
                    List<Integer> tmp = new ArrayList<>();
                    tmp.add(nums[i]);
                    tmp.add(nums[j]);
                    tmp.add(n);
                    if(!checkRepeat(res, tmp)){
                        res.add(tmp);
                    }

                }

                numMap.put(nums[j], jTime);
                numMap.put(nums[i], iTime);


            }

        }





        return res;
    }

    public static void main(String[] args){
        int[] data = new int[]{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6};

        _3SUM sum = new _3SUM();
        System.out.println(sum.threeSum(data));
    }
}
