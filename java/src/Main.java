//// http://codingbat.com/prob/p117334
public class Main {
    public static void main(String[] args) {
        Main.stringSplosion("Code");
        Main.stringSplosion("abc");
        Main.stringSplosion("There");
        Main.stringSplosion("Kitten");
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
