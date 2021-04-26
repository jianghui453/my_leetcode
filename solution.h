#include <string>
#include <vector>

using namespace std;

class Solution {
    public:
    int rangeBitwiseAnd(int, int);
    bool isHappy(int);
    int countPrimes(int);
    bool isIsomorphic(string, string);
    bool canFinish(int, vector< vector<int> >&);

    private:
    int bitSquareSum(int n);
    void topoDFS(int, vector<int>&, vector< vector<int> >&, bool&);
};