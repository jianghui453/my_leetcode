#include "../solution.h"
#include <unordered_map>

bool Solution::isIsomorphic(string t, string s)
{
    if (t.size() != s.size())
    {
        return false;
    }
    unordered_map<char, char> t2s;
    unordered_map<char, char> s2t;
    int len = s.length();
    for (int i = 0; i < len; ++i)
    {
        if ((t2s.count(t[i]) && t2s[t[i]] != s[i]) || (s2t.count(s[i]) && s2t[s[i]] != t[i]))
        {
            return false;
        }
        t2s[t[i]] = s[i];
        s2t[s[i]] = t[i];
    }
    return true;
}