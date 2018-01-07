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
}
