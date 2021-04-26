#include "../solution.h"
#include <vector>

bool Solution::canFinish(int numCourses, vector< vector<int> >& prerequisites) 
{
    vector<int> visited(numCourses, 0);
    vector< vector<int> > edges(numCourses);
    bool valid = true;
    for (int i = 0; i < prerequisites.size(); ++i)
    {
        edges[prerequisites[i][1]].push_back(prerequisites[i][0]);
    }

    for (int i = 0; i < numCourses && valid; ++i)
    {
        if (!visited[i])
        {
            this->topoDFS(i, visited, edges, valid);
        }
    }

    return valid;
}

void Solution::topoDFS(int u, vector<int>& visited, vector< vector<int> >& edges, bool& valid) {
    visited[u] = 1;
    for (int i = 0; i < edges[u].size(); ++i) 
    {
        int v = edges[u][i];
        if (visited[v] == 0)
        {
            this->topoDFS(v, visited, edges, valid);
            if (!valid)
            {
                return;
            }
        }
        if (visited[v] == 1)
        {
            valid = false;
            return;
        }
    }
    visited[u] = 2;
}
