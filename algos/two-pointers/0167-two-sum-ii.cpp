// Two Sum II - Input Array Is Sorted — https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Тема: два указателя
// Идея: массив отсортирован: сумма мала — двигаем левый, велика — правый.
//       Отброшенный элемент не даёт target ни с кем из оставшихся.
// Версии: 1) хеш-таблица  — O(n) время, O(n) память (нарушает условие O(1) памяти)
//         2) два указателя — O(n) время, O(1) память
// Заметка: если ответа нет, честнее return {}, а не пару индексов.
// Статус: хеш-таблица — сам; два указателя — с подсказкой
// Дата: 29.09.2026

// ---------- Версия 1: хеш-таблица, O(n) памяти ----------

class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
    vector<int> result;
    unordered_map<int,int>seen;
    for (int i =0; i<numbers.size();i++){
        int need=target-numbers[i];
        if (seen.count(need)){
            result.push_back(seen[need]);
            result.push_back(i+1);
            return result;
        }else {
            seen[numbers[i]]=i+1;
        }
    }
    return result;
}
};

// ---------- Версия 2: два указателя, O(1) памяти ----------

class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
            int left=0;
            int right= numbers.size()-1 ;
            while (left<right) {
                auto sum=numbers[left]+numbers[right];
                if (sum==target) {
                return {left+1,right+1};
                }
                else if (sum>target){
                right-=1;
                }
                else {
                    left+=1;
                }
            }
        return {left+1,right+1};
    }
};
