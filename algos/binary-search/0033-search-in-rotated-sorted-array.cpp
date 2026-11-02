// Search in Rotated Sorted Array — https://leetcode.com/problems/search-in-rotated-sorted-array/
// Тема: бинарный поиск (повёрнутый массив)
// Идея: в два шага. 1) Найти pivot — индекс минимума (как 153: сравнивать nums[mid] с nums[right]).
//       Массив распадается на две отсортированные части: [0 .. pivot-1] и [pivot .. n].
//       2) Если nums[pivot] <= target <= nums[n] — обычный бинпоиск (704) в правой части,
//       иначе — в левой. Бинпоиск вынесен в функцию с границами (lo, hi).
// Сложность: O(log n) время, O(1) память
// Заметка: споткнулся на логике — сначала думал, что pivot и есть ответ (цикл без target),
//          а на самом деле pivot только делит массив, и искать надо ВНУТРИ нужного подмассива.
//          Ещё по дороге: параметры границ объявил, но не использовал; после цикла right == pivot,
//          а не конец массива (конец — n); `<` вместо `<=` (target может равняться минимуму).
//          Бинпоиск работает только на отсортированном куске: левая часть строго до pivot-1.
//          Лучше назвать параметры lo, hi и передавать в порядке «начало, конец»; int pivot = left.
//          Способ за один проход (какая половина от mid отсортирована) — для перерешивания.
// Статус: сам с оговоркой (план «pivot + поиск в части» дан подсказкой, код и отладка — сам)
// Дата: 02.11.2026

class Solution {
public:
    int searchbin(vector<int>& nums, int ir, int lf, int target) {
        int left = lf;
        int right = ir;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid + 1;
            } else if (nums[mid] > target) {
                right = mid - 1;
            } else {
                return mid;
            }
        }
        return -1;
    }

    int search(vector<int>& nums, int target) {
        int left = 0;
        int n = nums.size() - 1;
        int right = n;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > nums[right]) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        if (nums[left] <= target && target <= nums[n]) {
            return searchbin(nums, n, left, target);
        } else {
            return searchbin(nums, left, 0, target);
        }
    }
};
