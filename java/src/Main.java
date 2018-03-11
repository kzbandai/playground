//// http://codingbat.com/prob/p117334
public class Main {
    public static void main(String[] args) {
//        Main.stringSplosion("Code");
//        Main.stringSplosion("abc");
//        Main.stringSplosion("There");
//        Main.stringSplosion("Kitten");
//
//        int u;
//        u = Main.sumNumbers("aa11b33");
//        System.out.println(u);
//        u = Main.sumNumbers("aa11b2");
//        System.out.println(u);

        HangmanLexicon hang = new HangmanLexicon();
        System.out.println(hang.getWordCount());
    }

    // http://codingbat.com/prob/p158767
    // https://techdevguide.withgoogle.com/paths/foundational/canbalance-problem-arrays-non-empty
    private static boolean canBalance(int[] nums) {
        int total = 0;
        for (int i = 0; i < nums.length; i++) {
            total += nums[i];
        }

        int sub = 0;
        for (int k = 0; k < nums.length; k++) {
            sub += nums[k];
            if (total - sub == sub) {
                return true;
            }
        }

        return false;
    }

    // https://techdevguide.withgoogle.com/paths/foundational/subnumbers-problem-string-return-sum
    private static int sumNumbers(String str) {
        char[] chars = str.toCharArray();

        int res = 0;
        String s = "";
        for (char ch : chars) {
            if (Character.isDigit(ch)) {
                s = s + String.valueOf(ch);
                System.out.println(s);
            } else if (!Character.isDigit(ch) && s.toCharArray().length > 0) {
                res += Integer.parseInt(s);
                s = "";
            }
        }

        if (s.toCharArray().length > 0) {
            res += Integer.parseInt(s);
        }

        return res;
    }

    private static void stringSplosion(String str) {
        String[] chars = str.split("");

        String res = "";
        for (int i = 0; i < chars.length; i++) {
            for (int k = 0; k <= i; k++) {
                res += chars[k];
            }
        }

        System.out.println(res);
    }

    // https://techdevguide.withgoogle.com/paths/foundational/maxspan-problem-return-largest-span-array
    // https://github.com/mirandaio/codingbat/blob/master/java/array-3/maxSpan.java
    public int maxSpan(int[] nums) {
        int max = 0;

        for (int i = 0; i < nums.length; i++) {
            int j = nums.length - 1;

            while (nums[i] != nums[j])
                j--;

            int span = j - i + 1;

            if (span > max)
                max = span;
        }

        return max;
    }

    // https://techdevguide.withgoogle.com/paths/foundational/withoutstring-problem-strings-base-remove-return
    public String withoutString(String base, String remove) {
        return base.replace(remove.toUpperCase(), "").replace(remove.toLowerCase(), "").replace(remove, "");
    }

    // http://codingbat.com/prob/p262890
    // https://techdevguide.withgoogle.com/paths/foundational/array-sort-problem-sorted-values
    int[] sort(int[] a) {
        int size = a.length;
        if (size == 0 || size == 1) {
            return a;
        }

        int[] tmp = new int[size];
        int tmp_v;
        int res_size = 0;

        for (int i = 0; i < size; i++) {
            tmp_v = a[i];
            if (i == 0) {
                tmp[res_size] = tmp_v;
                res_size++;
                continue;
            }

            if (tmp[res_size - 1] == tmp_v) {
                continue;
            }

            tmp[res_size] = tmp_v;
            res_size++;
        }

        int[] res = new int[res_size];
        for (int k = 0; k < res_size; k++) {
            res[k] = tmp[k];
        }

        return res;
    }

    // http://codingbat.com/prob/p294185
    // https://techdevguide.withgoogle.com/paths/foundational/interpreter-problem-for-java
    public int interpret(int value, String[] commands, int[] args) {
        for (int i = 0; i < commands.length; i++) {
            switch (commands[i]) {
                case "+":
                    value = value + args[i];
                    break;
                case "-":
                    value = value - args[i];
                    break;
                case "*":
                    value = value * args[i];
                    break;
                case "/":
                    value = value / args[i];
                    break;
                default:
                    return -1;
            }
        }

        return value;
    }

    // http://codingbat.com/prob/p152303
    // https://techdevguide.withgoogle.com/paths/foundational/word-zero-problem-warmup/#!
    public Map<String, Integer> word0(String[] strings) {
        ArrayList<String> tmp = new ArrayList<>();
        for (int i = 0; i < strings.length; i++) {
            tmp.add(strings[i]);
        }

        Map<String, Integer> result = new HashMap<>();
        for (int s = 0; s < tmp.size(); s++) {
            result.put(tmp.get(s), 0);
        }

        Object[] mapKey = result.keySet().toArray();
        Arrays.sort(mapKey);

        return result;
    }

    // http://codingbat.com/prob/p125327
    // https://techdevguide.withgoogle.com/paths/foundational/wordlen-problems-array-strings-medium/#!
    public Map<String, Integer> wordLen(String[] strings) {
        ArrayList<String> tmp = new ArrayList<>();
        for (int i = 0; i < strings.length; i++) {
            tmp.add(strings[i]);
        }

        Map<String, Integer> result = new HashMap<>();
        for (int s = 0; s < tmp.size(); s++) {
            result.put(tmp.get(s), tmp.get(s).length());
        }

        return result;
    }

    // http://codingbat.com/prob/p126332
    // https://techdevguide.withgoogle.com/paths/foundational/pairs-problem-classic-algorithm-hard/#!
    public Map<String, String> pairs(String[] strings) {
        ArrayList<String> tmp = new ArrayList<>();
        for (int i = 0; i < strings.length; i++) {
            tmp.add(strings[i]);
        }

        Map<String, String> result = new HashMap<>();
        for (int s = 0; s < tmp.size(); s++) {
            result.put(String.valueOf(tmp.get(s).charAt(0)), String.valueOf(tmp.get(s).charAt(tmp.get(s).length() - 1)));
        }

        return result;
    }

    // http://codingbat.com/prob/p117630
    // https://techdevguide.withgoogle.com/paths/foundational/wordcount-problem-classic-algorithm-hard/#!
    public Map<String, Integer> wordCount(String[] strings) {
        ArrayList<String> tmp = new ArrayList<>();
        for (int i = 0; i < strings.length; i++) {
            tmp.add(strings[i]);
        }

        Map<String, Integer> result = new HashMap<>();
        for (int s = 0; s < tmp.size(); s++) {
            String key = tmp.get(s);
            if (result.containsKey(key)) {
                int v = result.get(key);
                v++;
                result.put(key, v);
            } else {
                result.put(tmp.get(s), 1);
            }
        }

        return result;
    }
}
