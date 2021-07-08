#include "../solution.h"
#include <vector>

int Solution::countPrimes(int n)
{
    std::vector<int> isPrime(n, 1);
    int ans = 0;
    for (int i = 2; i < n; ++i)
    {
        if (isPrime[i])
        {
            ++ans;
            // 避免溢出
            if ((long long) i * i < n) {
                // 质数的所有倍数都是非质数
                // 从 i * i 开始，因为上次的已经把小于 i * i 的部分执行过了
                for (int j = i * i; j < n; j += i)
                {
                    isPrime[j] = 0;
                }
            }
        }
    }
    return ans;
}