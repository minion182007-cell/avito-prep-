// Minimum Size Subarray Sum — https://leetcode.com/problems/minimum-size-subarray-sum/
// Тема: скользящее окно (переменной длины)
// Идея: right расширяет окно (sum += nums[right]). ПОКА sum >= target — окно подходит:
//       записать длину, потом сжать слева (sum -= nums[left], left++), ищем короче.
// Сложность: O(n) время, O(1) память
// Заметка: условие задачи — сумма >= target, а не == target. Сначала ловил только
//          точное равенство и выкидывал подходящие окна (sum > target) не записав длину.
//          Запись длины — ВНУТРИ while, до сжатия. Длину можно считать как right-left+1.
// Статус: с подсказками
// Дата: 17.10.2026

class Solution {
public:
    int minSubArrayLen(int target, vector<int> &nums) {
        int left = 0;
        int num = 0;
        int sum = 0;
        int min1 = INT_MAX;

        for (int i = 0; i < nums.size(); ++i) {
            num += 1;
            sum += nums[i];
            while (target <= sum) {
                min1 = min(num, min1);
                sum -= nums[left];
                ++left;
                --num;
            }
        }
        if (min1 == INT_MAX) return 0;
        return min1;
    }
};
