#include <gtest/gtest.h>
#include "../solution.h"

TEST(SolutionTest, isIsomorphic)
{
    Solution* s = new Solution();
    EXPECT_TRUE(s->isIsomorphic("egg", "add"));
    EXPECT_FALSE(s->isIsomorphic("foo", "bar"));
    EXPECT_TRUE(s->isIsomorphic("paper", "title"));
    delete s;
}