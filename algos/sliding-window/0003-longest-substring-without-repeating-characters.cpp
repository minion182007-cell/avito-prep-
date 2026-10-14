// Longest Substring Without Repeating Characters — https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Тема: скользящее окно (переменной длины)
// Идея: right идёт по строке. ПОКА s[right] уже есть в окне — выкидываем символ слева
//       (left++). Потом добавляем s[right]. Длина окна = right - left + 1, берём максимум.
// Сложность: O(n·k) с окном-строкой (k — размер алфавита), O(n) с unordered_set
// Заметка: find возвращает позицию, а не bool — сравнивать с string::npos (позиция 0
//          даёт false, «не найдено» даёт true). wind.substr(1), а не substr(left+1):
//          у окна своя нумерация. right < s.size() (без -1), длина right-left+1.
//          Первая попытка сравнивала только s[left] и s[right] — повтор может быть в середине.
// Статус: с подсказками
// Дата: 14.10.2026

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int left = 0;
        int max1 = 0;
        string wind = "";
        int n = s.size();
        for (int right = 0; right <= n - 1; ++right) {
            while (wind.find(s[right]) != string::npos) {
                wind = wind.substr(1);
                ++left;
            }
            wind += s[right];
            max1 = max(max1, right - left + 1);
        }
        return max1;
    }
};

// Вариант с unordered_set (классический, O(n)):
// unordered_set<char> window;
// for (int right = 0; right < n; ++right) {
//     while (window.count(s[right])) { window.erase(s[left]); ++left; }
//     window.insert(s[right]);
//     best = max(best, right - left + 1);
// }
