// Search Insert Position — https://leetcode.com/problems/search-insert-position/
// Тема: бинарный поиск
// Идея: обычный бинпоиск. Если не нашли — вернуть left: после цикла left указывает
//       на первый элемент больше target, т.е. ровно туда, куда его надо вставить.
// Сложность: O(log n) время, O(1) память
// Заметка: отличие от 704 — одна строка: return left вместо return -1.
//          Сначала добавил sort(): массив по условию уже отсортирован, а sort — это O(n log n),
//          он убивает весь смысл бинпоиска (O(log n)) и меняет входные данные.
// Статус: сам
// Дата: 20.10.2026

class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
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
        return left;
    }
};
