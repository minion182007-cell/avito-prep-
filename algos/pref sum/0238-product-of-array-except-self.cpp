// Range Sum Query - Immutable — https://leetcode.com
// Тема: префиксные суммы
// Идея: массив не меняется, запросов много. Один раз в конструкторе считаем pref[i] — сумму
//       элементов с 0-го по (i-1)-й. Сумма на отрезке [left, right] = pref[right+1] - pref[left].
// Сложность: O(n) на конструктор, O(1) на запрос; O(n) память
// Заметка: ловушка — left == 0: pref[left-1] выходит за границу. Решается сдвигом:
//          слайс pref размера n+1, pref[i+1] = pref[i] + nums[i], ответ pref[right+1] - pref[left].
// Статус: сам
// Дата: 18.11.2026

class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int n=nums.size();
        vector<int> pref(nums.size()+1,1);
        vector<int> suf(nums.size()+1,1);
        vector<int> answer(nums.size(),0);

        for (int i=0;i<n;++i) {
            pref[i+1]=pref[i]*nums[i];
        }
        for (int i = n - 2; i >= 0; --i) {
            suf[i] = suf[i + 1] * nums[i + 1];
        }
        for (int i=0;i<n;++i) {
            answer[i]=pref[i]*suf[i];
        }
        return answer;
    }
};
