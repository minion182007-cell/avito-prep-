// Binary Search — https://leetcode.com/problems/binary-search/
// Тема: бинарный поиск
// Идея: массив отсортирован. Смотрим середину: равна — нашли; меньше target — ищем справа
//       (left = mid + 1); больше — слева (right = mid - 1). Пока left <= right.
// Сложность: O(log n) время, O(1) память
// Заметка: mid = left + (right - left) / 2 — без переполнения int (в отличие от (left+right)/2).
//          Условие left <= right (а не <): иначе пропустим отрезок из одного элемента.
//          mid ± 1 — иначе бесконечный цикл, когда left == right.
// Статус: сам
// Дата: 20.10.2026

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] == target) {
                return mid;
            }
            if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return -1;
    }
};
