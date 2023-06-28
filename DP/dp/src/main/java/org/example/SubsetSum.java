package org.example;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class SubsetSum {

    static Map<String, Boolean> dp = new HashMap<>();

    static boolean isSubsetSum(int set[], int n, int sum){
        if (sum == 0) {
            return true;
        }
        if (n == 0) {
            return false;
        }

        String key1 = (n -1) + "_" + sum;
        if (dp.containsKey(key1)) {
            return dp.get(key1);
        }

        String key2 = (n -1) + "_" + (sum - set[n - 1]);
        if (dp.containsKey(key2)) {
            return dp.get(key2);
        }

        boolean b1 = isSubsetSum(set, n - 1, sum);
        dp.put(key1, b1);

        boolean b2 = isSubsetSum(set, n - 1, sum - set[n - 1]);
        dp.put(key2, b2);
        return b1 || b2;
    }

    static boolean isSubsetSum_DP(int set[], int n, int sum){
        boolean db[][] = new boolean[n + 1][sum + 1];
        for(int i = 0; i <= n; i++) {
            db[i][0] = true;
        }

        // bottom - up
        for(int j = 1; j <= sum; j++) {
            for (int i = 1; i <= n; i++) {
                db[i][j] = db[i - 1][j];
                if (j >= set[i - 1]) {
                    db[i][j] = db[i][j] || db[i - 1][j - set[i - 1]];
                }
            }
        }
        return db[n][sum];
    }

    public static void main(String[] args) {
        int set[] = { 3, 4, 5, 2};
        int sum = 13;
        int n = set.length;
        System.out.println(isSubsetSum_DP(set, n, sum));
    }
}
