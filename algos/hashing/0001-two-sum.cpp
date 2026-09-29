// Two Sum — https://leetcode.com/problems/two-sum/
// Тема: хеш-таблица
// Идея: храним уже виденные числа с индексами; для каждого x ищем в мапе target - x.
//       Сначала проверяем, потом кладём — иначе элемент найдёт сам себя.
// Версии: 1) перебор пар  — O(n²) время, O(1) память
//         2) хеш-таблица  — O(n) время,  O(n) память
// Статус: перебор — сам; хеш-таблица — с подсказками
// Дата: 22.09.2026

// ---------- Версия 1: перебор пар, O(n²) ----------

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

// ---------- Версия 2: хеш-таблица, O(n) ----------

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
