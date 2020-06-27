package leetcode;



public class ProductOfArrayExceptSelf {


    public int divid(int a , int b){

        boolean flag = true;
        if(a > 0 && b < 0 || a < 0 && b > 0){
            flag = false;
        }

        int t = 1;
        if(!flag){
            t = -1;
        }

        int m = Math.abs(a);
        int n = Math.abs(b);

        if(m < n){
            return  0;
        }

        if(n == 1) {
            return m*t;
        }

        for(int i = 1 ; i < m ; i ++){
            if(n * i >= m){
                return i*t;
            }
        }
        return m*t;
    }

    public int[] productExceptSelf(int[] nums) {

        if(nums == null || nums.length <= 1){
            return nums;
        }


        int zeroNum = 0;

        int sum = 1;
        int nozeroSum = 1;
        for(int i = 0 ; i < nums.length ; i ++){
            sum *= nums[i];

            if(nums[i] == 0){
                zeroNum ++;
                continue;
            }
            nozeroSum *= nums[i];
        }

        System.out.println("sum:" + sum);

        if (zeroNum > 1){
            for(int i = 0 ; i < nums.length; i ++){
                nums[i] = 0;
            }
        } else if (zeroNum == 0){
            for(int i = 0 ; i < nums.length; i ++){
                nums[i] = divid(sum, nums[i]);
            }
        } else {
            for(int i =0 ; i < nums.length ; i ++){
                if(nums[i] != 0){
                    nums[i] = divid(sum, nums[i]);
                } else{
                    nums[i] = nozeroSum;
                }

            }

        }
        return nums;


    }


    public static void main(String[] args){
        int[] data = new int[]{5,9,2,-9,-9,-7,-8,7,-9,10};

        ProductOfArrayExceptSelf productOfArrayExceptSelf = new ProductOfArrayExceptSelf();
        productOfArrayExceptSelf.productExceptSelf(data);

        System.out.println(productOfArrayExceptSelf.divid(-257191200, 9));
    }

}
