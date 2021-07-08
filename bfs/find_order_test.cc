#include "../solution.h"
#include <gtest/gtest.h>
#include <vector>

TEST(SolutionTest, findOrder)
{
    Solution* s = new Solution();

    int numCourses;
    vector< vector<int> > prerequisites(1);
    vector<int> ret;

    prerequisites[0].push_back(1);
    prerequisites[0].push_back(0);
    ret.push_back(0);
    ret.push_back(1);
    EXPECT_EQ(s->findOrder(2, prerequisites), ret);

    // prerequisites.resize(4);
    // prerequisites[1].push_back(2);
    // prerequisites[1].push_back(0);
    // prerequisites[2].push_back(3);
    // prerequisites[2].push_back(1);
    // prerequisites[3].push_back(3);
    // prerequisites[3].push_back(2);
    // ret.push_back(2);
    // ret.push_back(3);
    // EXPECT_EQ(s->findOrder(4, prerequisites), ret);

    delete(s);
}