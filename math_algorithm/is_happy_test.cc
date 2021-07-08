#include <gtest/gtest.h>
#include "../solution.h"

TEST(SolutionTest, isHappy)
{
    Solution* s = new Solution();
    EXPECT_TRUE(s->isHappy(19));
    EXPECT_FALSE(s->isHappy(2));
    EXPECT_TRUE(s->isHappy(7));
    delete s;
}
