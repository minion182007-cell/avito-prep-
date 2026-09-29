// Contains Duplicate — https://leetcode.com/problems/contains-duplicate/
// Тема: хеш-множество
// Идея: множество уже виденных чисел; встретили повтор — сразу true.
// Сложность: O(n) время, O(n) память
// Статус: с подсказками
// Дата: 22.09.2026

class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        unordered_set<int> seen;
        for (int i =0;i<nums.size();i++){
            if (seen.count(nums[i])){
                return true;
            }else{
                seen.insert(nums[i]); 
            }
        }
        return false;
    }
};
