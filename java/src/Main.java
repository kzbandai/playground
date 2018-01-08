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
}
