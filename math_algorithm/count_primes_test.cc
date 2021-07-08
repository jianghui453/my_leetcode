#include <gtest/gtest.h>
#include "../solution.h"

TEST(SolutionTest, countPrimes)
{
    Solution* s = new Solution();
    EXPECT_EQ(s->countPrimes(10), 4);
    delete s;
}