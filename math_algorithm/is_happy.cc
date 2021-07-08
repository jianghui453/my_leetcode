#include "../solution.h"
#include <iostream>
#include <math.h>

bool Solution::isHappy(int n)
{
    int fast = this->bitSquareSum(n);
    int slow = n;
    while (slow != fast)
    {
        fast = this->bitSquareSum(fast);
        fast = this->bitSquareSum(fast);
        slow = this->bitSquareSum(slow);
    }
    return slow == 1;
}

int Solution::bitSquareSum(int n)
{
    int sum = 0;
    while (n > 0)
    {
        sum += pow((n % 10), 2);
        n = n / 10;
    }
    return sum;
}