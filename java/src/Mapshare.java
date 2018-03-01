import java.util.Map;

// http://codingbat.com/prob/p148813
// https://techdevguide.withgoogle.com/paths/foundational/mapshare-problem-return-given-map/
public class Mapshare {
    public Map<String, String> mapShare(Map<String, String> map) {
        if (map.containsKey("a")) {
            map.put("b", map.get("a"));
        }

        if (map.containsKey("c")) {
            map.remove("c");
        }

        return map;
    }
}
