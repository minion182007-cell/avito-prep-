// Range Sum Query - Immutable — https://leetcode.com/problems/range-sum-query-immutable/
// Тема: префиксные суммы
// Идея: массив не меняется, запросов много. Один раз в конструкторе считаем pref[i] — сумму
//       элементов с 0-го по i-й. Сумма на отрезке [left, right] = pref[right] - pref[left-1].
// Сложность: O(n) на конструктор, O(1) на запрос; O(n) память
// Заметка: ловушка — left == 0: pref[left-1] выходит за границу. Решается веткой
//          или сдвигом: pref размера n+1, pref[i+1] = pref[i] + nums[i], ответ pref[right+1] - pref[left].
// Статус: сам
// Дата: 16.11.2026

class NumArray {
private:
    vector<int> pref;

public:
    NumArray(vector<int> &nums) :pref(nums.size()+1,0){
        for (int i = 0; i < nums.size(); ++i) {
             pref[i+1] = pref[i] + nums[i];
        }
    }

    int sumRange(int left, int right) {
        return pref[right+1]-pref[left];
    }
};
