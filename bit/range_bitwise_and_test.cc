#include <gtest/gtest.h>
#include "../solution.h"

TEST(SolutionTest, rangeBitwiseAnd)
{
    Solution* s = new Solution();

    EXPECT_EQ(s->rangeBitwiseAnd(5, 7), 4);

    delete s;
}