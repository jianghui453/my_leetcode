#include "../solution.h"
#include <vector>
#include <iostream>

using namespace std;

vector<int> Solution::findOrder(int numCourses, vector< vector<int> >& prerequisites)
{
    vector< vector<int> > edges(numCourses);
    vector<int> inputs(numCourses, 0);
    vector<int> courses(numCourses, 1);

    int size = prerequisites.size();
    for (int i = 0; i < size; ++i)
    {
        vector<int> prerequisite = prerequisites[i];
        edges[prerequisite[1]].push_back(prerequisite[0]);

        ++inputs[prerequisite[0]];
    }

    vector<int> ret;

    while (ret.size() < numCourses)
    {
        cout << ret.size() << "\n";
        for (int i = 0; i < numCourses; ++i)
        {
            if (courses[i] == 0)
            {
                continue;
            }
            if (inputs[i] == 0)
            {
                ret.push_back(i);
                courses[i] = 0;
                for (int j = 0; j < edges[i].size(); ++j)
                {
                    --inputs[j];
                }
            }
        }
    }
    return ret;
}