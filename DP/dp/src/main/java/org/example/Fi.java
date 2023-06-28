package org.example;

import java.util.Arrays;

public class Fi {

    static int[] dp = new int[100000];

    public static int fib(int n) {
        if (n <= 1) {
            return 1;
        }

        if (dp[n] != -1) {
            return dp[n];
        }

        return  dp[n] = fib(n - 1) + fib(n - 2);
    }

    public static void main(String[] args) {
        Arrays.fill(dp, -1);
        System.out.println(fib(16));
    }
}
