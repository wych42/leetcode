/*
 * @lc app=leetcode.cn id=121 lang=java
 *
 * [121] Best Time to Buy and Sell Stock
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (57.98%)
 * Likes:    3017
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.9M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * You are given an array prices where prices[i] is the price of a given stock
 * on the i^th day.
 * 
 * You want to maximize your profit by choosing a single day to buy one stock
 * and choosing a different day in the future to sell that stock.
 * 
 * Return the maximum profit you can achieve from this transaction. If you
 * cannot achieve any profit, return 0.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: prices = [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
 * = 6-1 = 5.
 * Note that buying on day 2 and selling on day 1 is not allowed because you
 * must buy before you sell.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: prices = [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transactions are done and the max profit =
 * 0.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= prices.length <= 10^5
 * 0 <= prices[i] <= 10^4
 * 
 * 
 */

// @lc code=start
class Solution {
    // brute force
    public int maxProfitBruteForce(int[] prices) {
        int maxProfit = 0;
        for (int i=0;i<prices.length;i++) {
            int buyPrice = prices[i];
            // can'n sell before buy
            for (int j=i+1;j<prices.length;j++) {
                int profit = prices[j] - buyPrice;
                if (profit>maxProfit) {
                    maxProfit = profit; 
                }
            }
        }
        return maxProfit;
    }

    // we want to buy at lowest price, sell at highest price
    // we must buy before sell
    // go throught the prices:
    // record the lowest price so far
    // record max profile since the lowest price
    // if an lower price than current lowest, replace the lowes price, continue the process
    public int maxProfit(int[] prices) {
        int lowestPrice = 1<<31 - 1;
        int maxProfit = 0;
        for (int i=0;i<prices.length;i++) {
            int currentPrice = prices[i];
            if (currentPrice < lowestPrice) {
                lowestPrice =  currentPrice;
            }
            maxProfit = max(maxProfit, currentPrice-lowestPrice);
        }
        return maxProfit;
    }

    private int max(int a, int b) {
        return a>b ? a:b;
    }


}
// @lc code=end

