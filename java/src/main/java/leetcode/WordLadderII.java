package leetcode;

import java.util.*;

public class WordLadderII {

    private boolean diff(String a, String b){
        int n = 0;

        if(a.length() != b.length()){
            return false;
        }

        for(int i = 0; i < a.length() ; i ++){
            if(a.charAt(i) != b.charAt(i)){
                n ++;
            }
        }

        return n == 1;
    }

    private List<String> getNextPath(Map<String,Set<String>> path, String nowWord, Set<String> gone){
        List<String> res = new ArrayList<>();

        Set<String> nowSet = path.get(nowWord);

        for(String key: nowSet){
           if(!gone.contains(key)){
               res.add(key);
           }
        }

        return res;
    }

    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordLists) {


        if(wordLists == null || wordLists.size() == 0){
            return new ArrayList<>();
        }

        List<String> wordList = new ArrayList<>(wordLists);
        wordList.add(beginWord);

        // 开始词和结尾词都无
        int beginIndex = -1;
        int endIndex = -1;
        for(int i = 0 ; i < wordList.size() ; i ++){
            if(wordList.get(i).equals(beginWord)){
                beginIndex = i;
            }
            if(wordList.get(i).equals(endWord)){
                endIndex = i;
            }
        }

        if(beginIndex == -1 || endIndex == -1){
            return new ArrayList<>();
        }


        // 构建图
        Map<String,Set<String>> path = new HashMap<>();

        for(int i = 0; i < wordList.size() ; i ++){
            Set<String> tmp = new HashSet<>();
            for(int j = 0 ; j < wordList.size() ; j ++){
                if(i == j) {
                   continue;
                } else if(diff(wordList.get(i), wordList.get(j))){
                    tmp.add(wordList.get(j));
                }
            }
            path.put(wordList.get(i), tmp);
        }



        System.out.println(path);

        List<List<String>> res = new ArrayList<>();

        Set<String> gone = new HashSet<>();

        List<String> nowPath = new ArrayList<>();
        // 深度优先搜索
//        DFSearch(resInt, beginIndex, endIndex, path, gone, nowPath);


        // 广度优先搜索
        Queue<String> queue = new LinkedList<>();
        Queue<List<String>> nowPathQueue = new LinkedList<>();
        Queue<Set<String>> nowGoneQueue = new LinkedList<>();
        queue.add(beginWord);
        nowPath.add(beginWord);
        gone.add(beginWord);

        nowPathQueue.add(nowPath);
        nowGoneQueue.add(gone);

        int minStep = wordList.size();

        while (!queue.isEmpty()){
            String now = queue.remove();
            List<String> nowList = new ArrayList<>(nowPathQueue.remove());
            Set<String> nowGone = new HashSet<>(nowGoneQueue.remove());

            if(nowList.size() > minStep){
                break;
            }

            if(now.equals(endWord)){
                res.add(nowList);
                if(nowList.size() < minStep){
                    minStep = nowList.size();
                }
                break;
            }

            List<String> nearBys = getNextPath(path, now, nowGone);
            for(String nearBy: nearBys){
                queue.add(nearBy);
                Set<String> tmpGone = new HashSet<>(nowGone);
                List<String> tmpList = new ArrayList<>(nowList);
                tmpGone.add(nearBy);
                tmpList.add(nearBy);
                nowPathQueue.add(tmpList);
                nowGoneQueue.add(tmpGone);
            }
        }

        return res;
    }

    public static void main(String[] args){
        String begin = "cet";
        String end = "ism";


        List<String> wordList = Arrays.asList("kid","tag","pup","ail","tun","woo","erg","luz","brr","gay","sip","kay","per","val","mes","ohs","now","boa","cet","pal","bar","die","war","hay","eco","pub","lob","rue","fry","lit","rex","jan","cot","bid","ali","pay","col","gum","ger","row","won","dan","rum","fad","tut","sag","yip","sui","ark","has","zip","fez","own","ump","dis","ads","max","jaw","out","btu","ana","gap","cry","led","abe","box","ore","pig","fie","toy","fat","cal","lie","noh","sew","ono","tam","flu","mgm","ply","awe","pry","tit","tie","yet","too","tax","jim","san","pan","map","ski","ova","wed","non","wac","nut","why","bye","lye","oct","old","fin","feb","chi","sap","owl","log","tod","dot","bow","fob","for","joe","ivy","fan","age","fax","hip","jib","mel","hus","sob","ifs","tab","ara","dab","jag","jar","arm","lot","tom","sax","tex","yum","pei","wen","wry","ire","irk","far","mew","wit","doe","gas","rte","ian","pot","ask","wag","hag","amy","nag","ron","soy","gin","don","tug","fay","vic","boo","nam","ave","buy","sop","but","orb","fen","paw","his","sub","bob","yea","oft","inn","rod","yam","pew","web","hod","hun","gyp","wei","wis","rob","gad","pie","mon","dog","bib","rub","ere","dig","era","cat","fox","bee","mod","day","apr","vie","nev","jam","pam","new","aye","ani","and","ibm","yap","can","pyx","tar","kin","fog","hum","pip","cup","dye","lyx","jog","nun","par","wan","fey","bus","oak","bad","ats","set","qom","vat","eat","pus","rev","axe","ion","six","ila","lao","mom","mas","pro","few","opt","poe","art","ash","oar","cap","lop","may","shy","rid","bat","sum","rim","fee","bmw","sky","maj","hue","thy","ava","rap","den","fla","auk","cox","ibo","hey","saw","vim","sec","ltd","you","its","tat","dew","eva","tog","ram","let","see","zit","maw","nix","ate","gig","rep","owe","ind","hog","eve","sam","zoo","any","dow","cod","bed","vet","ham","sis","hex","via","fir","nod","mao","aug","mum","hoe","bah","hal","keg","hew","zed","tow","gog","ass","dem","who","bet","gos","son","ear","spy","kit","boy","due","sen","oaf","mix","hep","fur","ada","bin","nil","mia","ewe","hit","fix","sad","rib","eye","hop","haw","wax","mid","tad","ken","wad","rye","pap","bog","gut","ito","woe","our","ado","sin","mad","ray","hon","roy","dip","hen","iva","lug","asp","hui","yak","bay","poi","yep","bun","try","lad","elm","nat","wyo","gym","dug","toe","dee","wig","sly","rip","geo","cog","pas","zen","odd","nan","lay","pod","fit","hem","joy","bum","rio","yon","dec","leg","put","sue","dim","pet","yaw","nub","bit","bur","sid","sun","oil","red","doc","moe","caw","eel","dix","cub","end","gem","off","yew","hug","pop","tub","sgt","lid","pun","ton","sol","din","yup","jab","pea","bug","gag","mil","jig","hub","low","did","tin","get","gte","sox","lei","mig","fig","lon","use","ban","flo","nov","jut","bag","mir","sty","lap","two","ins","con","ant","net","tux","ode","stu","mug","cad","nap","gun","fop","tot","sow","sal","sic","ted","wot","del","imp","cob","way","ann","tan","mci","job","wet","ism","err","him","all","pad","hah","hie","aim","ike","jed","ego","mac","baa","min","com","ill","was","cab","ago","ina","big","ilk","gal","tap","duh","ola","ran","lab","top","gob","hot","ora","tia","kip","han","met","hut","she","sac","fed","goo","tee","ell","not","act","gil","rut","ala","ape","rig","cid","god","duo","lin","aid","gel","awl","lag","elf","liz","ref","aha","fib","oho","tho","her","nor","ace","adz","fun","ned","coo","win","tao","coy","van","man","pit","guy","foe","hid","mai","sup","jay","hob","mow","jot","are","pol","arc","lax","aft","alb","len","air","pug","pox","vow","got","meg","zoe","amp","ale","bud","gee","pin","dun","pat","ten","mob");
        WordLadderII wordLadderII = new WordLadderII();

        System.out.println(wordLadderII.findLadders(begin, end, wordList));
    }
}
