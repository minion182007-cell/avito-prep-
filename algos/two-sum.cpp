class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
      for (int i =0; i<nums.size();i++){
         for (int j=i+1;j<nums.size();j++){
            if(nums[i]==target-nums[j]) return {i,j};
        }
        }
        return {};
    }
    
};
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int, int> seen;
      for (int i =0; i<nums.size();i++){
        int need = target - nums[i];
            if(seen.count(need) ){
                return {seen[need],i};
            }else{
                seen[nums[i]]=i;
            }
        }
        return {};
    }
    
};
