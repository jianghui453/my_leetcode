// 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

// 现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

// 求所能获得硬币的最大数量。

// 说明:

// 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
// 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
// 示例:

// 输入: [3,1,5,8]
// 输出: 167 
// 解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//      coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

package array

// 【分析】区间型dp，从最后一步出发，最后一步必定扎破一个气球，编号为i，这一步获得金币1* nums[i] * 1，此时i前面的气球1～i-1以及i后面的气球i+1～n都被扎破了，需要知道两边最多能获得多少个金币，再加上最后一步，就是结果。
// 【状态转移方程】由于最后一步是1 * nums[i] * 1，我们可以认为两端有两个不能扎破的气球，值为1，dp[i] [j]代表扎破i+1号气球～j-1号气球能获得的金币数，i和j是不能被扎破的，因为是两端，并且当前气球k不能被扎破，要分别考虑k的左侧（i～k-1）和右侧（k+1～j），状态转移方程为：
// dp[i][j] = max{dp[i][k] + dp[k][j] + a[i] * a[k] * a[j]},k∈(i,j)
// dp[i] [k]代表扎破i+1～k-1号气球，dp[k] [j]代表扎破k+1～j-1号气球，再加上扎破这个气球获得的金币数
// 【初始条件】没有气球要扎破就获得0个金币
// dp[0][1] = dp[1][2] = ... = dp[n-2][n-1] = 0

import (
	// "fmt"
)

func maxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)

	var (
		l = len(nums)
		dp = make([][]int, l)
	)

	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	for r := 2; r < l; r++ {
		for i := 0; i < l-r; i++ {
			j := i+r
			for k := i+1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
			}
		}
	}
	
	return dp[0][l-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// class Solution {
// public:
// 	int maxCoins(vector<int>& nums) {
// 		nums.insert(nums.begin(),1);
// 		nums.push_back(1);
// 		int n=nums.size();
// 		int dp[n][n];   //dp[i][j]表示第i至第j个元素这个区间能获得的最大硬币数
// 		for(int i=0;i<n;i++)
// 			for(int j=0;j<n;j++)
// 				dp[i][j]=0;

// 		for(int r=2;r<n;r++)            //r为区间长度
// 			for(int i=0;i<n-r;i++){    //i为左区间
// 				int j=i+r;            //j为右区间
// 				for(int k=i+1;k<j;k++)
// 					dp[i][j]=max(dp[i][j],dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j]);
// 			}

// 		return dp[0][n-1];
// 	}
// };
