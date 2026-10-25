// First Bad Version — https://leetcode.com/problems/first-bad-version/
// Тема: бинарный поиск (поиск первого true)
// Идея: версии выглядят как false false ... true true. Ищем первую true.
//       Если mid плохая — ответ mid или левее: right = mid (mid НЕ выкидываем).
//       Если хорошая — ответ правее: left = mid + 1. Цикл while (left < right),
//       выходим, когда left == right — это и есть первая плохая версия.
// Сложность: O(log n) время, O(1) память
// Заметка: отличие от 704 — три детали: left < right (не <=), right = mid (не mid - 1),
//          нет проверки на равенство, ответ — left после цикла.
//          mid = left + (right - left) / 2: n до 2^31 - 1, left + right переполнит int.
// Статус: сам
// Дата: 25.10.2026

// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int left = 1;
        int right = n;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (isBadVersion(mid)) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        return left;
    }
};
