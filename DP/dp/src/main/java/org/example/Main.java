package org.example;

import java.util.Arrays;

public class Main {

    static int[] dp = new int[100000];
    public static int solve(int n)
    {
        // base case
        if (n < 0)
            return 0;
        if (n == 0)
            return 1;

        if (dp[n] != -1)
            return dp[n];
        return dp[n] = solve(n - 1) + solve(n - 3) + solve(n - 5);
    }

    public static void main(String[] args) {
        Arrays.fill(dp, -1);
        System.out.println(solve(6));
    }
}