class Solution {
public:
    int pivotIndex(vector<int>& nums) {
        int n = nums.size();
        vector<int> pref(n + 1, 0);
        for (int i = 0; i < n; ++i)
            pref[i + 1] = pref[i] + nums[i];
        for (int i = 0; i < n; ++i)
            if (pref[i] == pref[n] - pref[i + 1]) return i;
        return -1;
    }
};
class Solution {
public:
    int pivotIndex(vector<int> &nums) {
        int total_sum = 0;
        for (int num : nums) total_sum += num;
        int left_sum=0;
        for (int i=0; i < nums.size(); ++i) {
            if (left_sum==total_sum-left_sum-nums[i]) return i;
            left_sum+=nums[i];
        }
        return -1;
    }
};