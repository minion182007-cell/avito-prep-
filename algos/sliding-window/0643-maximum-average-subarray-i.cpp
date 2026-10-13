// Maximum Average Subarray I — https://leetcode.com/problems/maximum-average-subarray-i/
// Тема: скользящее окно (фиксированной длины)
// Идея: сумма первых k элементов — первое окно. Сдвиг на шаг: + вошедший nums[right+1],
//       - вышедший nums[left]. Запоминаем лучшую сумму, в конце делим на k.
// Сложность: O(n) время, O(1) память
// Заметка: в лоб (пересчитывать каждое окно) — O(n·k). Сумма в double, чтобы деление
//          на k не стало целочисленным (int / int = 51 / 4 = 12, а нужно 12.75).
// Статус: сам
// Дата: 14.10.2026

class Solution {
public:
    double findMaxAverage(vector<int>& nums, int k) {
        int left = 0;
        int right = k - 1;
        double max1 = 0;
        for (int i = 0; i < k; ++i) {
            max1 += nums[i];
        }
        double result = max1;

        while (right != nums.size() - 1) {
            max1 = max1 - nums[left] + nums[right + 1];
            result = max(max1, result);
            ++left;
            ++right;
        }

        return result / k;
    }
};
