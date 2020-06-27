package demo;

import com.alibaba.fastjson.JSON;

import java.util.HashMap;
import java.util.Map;

public class MapDemo {

    public static void main(String[] args){
        Map<String,Object> data = new HashMap<>();

        Long a = 6849898923242l;

        data.put("ff", a);

        System.out.println(JSON.toJSONString(data));

    }
}
