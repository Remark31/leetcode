package leetcode;

public class AddBinary {

    public String addBinary(String a, String b) {

        if(a.length() < b.length()){
            return addBinary(b , a);
        }


        StringBuilder sb = new StringBuilder();


        boolean flag = false;


        int end = 0;
        while (b.length() - end > 0){

            char ta = a.charAt(a.length() - end - 1);
            char tb = b.charAt(b.length() - end - 1);
            if(ta == '0' && tb == '0' && !flag){
                sb.insert(0 , "0");
                flag = false;
            } else if(ta == '0' && tb == '0' && flag){
                sb.insert(0 , "1");
                flag = false;
            }
            else if (ta != tb && !flag) {
                sb.insert(0 , "1");
                flag = false;
            } else if (ta != tb && flag){
                sb.insert(0, "0");
                flag = true;
            } else if (ta == tb && !flag){
                sb.insert(0, "0");
                flag = true;
            } else {
                sb.insert(0, "1");
                flag = true;
            }
            end ++;
        }




        while (a.length() - end > 0){
            char ta = a.charAt(a.length() - end - 1);
            if(ta == '0' && flag){
                sb.insert(0, "1");
                flag = false;
            } else if(ta == '0' && !flag){
                sb.insert(0, "0");
                flag = false;
            } else if(ta == '1' && flag){
                sb.insert(0, "0");
                flag = true;
            } else {
                sb.insert(0, "1");
                flag = false;
            }
            end ++;
        }


        if(flag){
            sb.insert(0, "1");
        }

        return sb.toString();
    }
}
