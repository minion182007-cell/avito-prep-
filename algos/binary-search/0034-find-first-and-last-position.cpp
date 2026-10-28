// Find First and Last Position of Element in Sorted Array — https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Тема: бинарный поиск (lower bound)
// Идея: lowerBound(x) — первый индекс, где nums[i] >= x («куда встала бы x», x может не быть в массиве).
//       first = lowerBound(target). Если first == size или nums[first] != target → [-1, -1].
//       last  = lowerBound(target + 1) - 1: первый элемент больше target, шаг назад — последний равный.
// Сложность: O(log n) время, O(1) память
// Заметка: шаблон 278: right = size() (не -1), while (left < right), >= → right = mid, иначе left = mid + 1.
//          Ошибки по дороге: искал равенство внутри цикла (ветка else недостижима при <= и >=),
//          менял входной массив, смешал шаблоны 704 и 278, равенство уводил вправо (это upper bound),
//          потом поменял условие, но не действие. Правило: куда уходит равенство — то и ищешь.
//          Проверка границы first == size ДО nums[first], иначе выход за массив (assert на LeetCode).
// Статус: с подсказками
// Дата: 28.10.2026

class Solution {
public:
    int lowerBound(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size();
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] >= target) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        return left;
    }

    vector<int> searchRange(vector<int>& nums, int target) {
        int first = lowerBound(nums, target);
        if (first == nums.size() || nums[first] != target) {
            return {-1, -1};
        }
        int last = lowerBound(nums, target + 1) - 1;
        return {first, last};
    }
};
