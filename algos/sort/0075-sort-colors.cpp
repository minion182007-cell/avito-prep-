// Sort Colors — https://leetcode.com/problems/sort-colors/
// Тема: сортировки, три указателя (голландский флаг)
// Идея: три указателя. low — граница нулей, high — граница двоек, mid — текущий элемент.
//       Инвариант: [0, low) — нули, [low, mid) — единицы, (high, n) — двойки,
//       [mid, high] — ещё не просмотрено. Смотрим nums[mid]: 0 — обмен с low, оба вперёд;
//       1 — просто mid++; 2 — обмен с high, high--, mid НЕ двигаем.
// Сложность: O(n) время, один проход; O(1) память
// Заметка: ловушки — после обмена с high mid не двигаем (пришедший элемент не проверен),
//          после обмена с low двигаем (там была единица или тот же элемент);
//          цикл while (mid <= high), не строго меньше.
// Статус: с подсказками (подглядел)
// Дата: 22.11.2026

class Solution {
public:
    void sortColors(vector<int> &nums) {
        int low = 0;
        int high = nums.size()-1;
        int mid=0;
        while (mid<=high) {
            if (nums[mid]==0) {
                swap(nums[low],nums[mid]);
                low++;
                mid++;
            }
            else if(nums[mid]==1) {
                mid++;
            }
            else if (nums[mid]==2) {
                swap(nums[high],nums[mid]);
                high--;
            }
        }
    }
};
