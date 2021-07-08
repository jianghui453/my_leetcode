#include <gtest/gtest.h>
#include "../solution.h"
#include <vector>

using namespace std;

TEST(SolutionTest, canFinish)
{
    Solution* s = new Solution();

    vector< vector<int> > prerequisites(1);
    prerequisites[0].push_back(1);
    prerequisites[0].push_back(0);
    EXPECT_TRUE(s->canFinish(2, prerequisites));
    prerequisites.resize(2);
    prerequisites[1].push_back(0);
    prerequisites[1].push_back(1);
    EXPECT_FALSE(s->canFinish(2, prerequisites));
    prerequisites.resize(4);
    prerequisites[0].clear();
    prerequisites[0].push_back(0);
    prerequisites[0].push_back(1);
    prerequisites[1].clear();
    prerequisites[1].push_back(3);
    prerequisites[1].push_back(1);
    prerequisites[2].push_back(1);
    prerequisites[2].push_back(3);
    prerequisites[3].push_back(3);
    prerequisites[3].push_back(2);
    EXPECT_FALSE(s->canFinish(4, prerequisites));
    delete(s);
}