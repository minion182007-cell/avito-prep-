# Алгоритмы

Решения задач LeetCode на C++, разложенные по темам. В начале каждого файла есть шапка: идея, сложность, статус и дата.

## Задачи

| №    | Задача | Тема | Сложность | Статус | Файл |
|------|--------|------|-----------|--------|------|
| 1    | [Two Sum](https://leetcode.com/problems/two-sum/) | хеш-таблица | O(n) / O(n) | brute force — сам, map — с подсказками | [hashing/0001-two-sum.cpp](hashing/0001-two-sum.cpp) |
| 217  | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | хеш-множество | O(n) / O(n) | с подсказками | [hashing/0217-contains-duplicate.cpp](hashing/0217-contains-duplicate.cpp) |
| 242  | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | подсчёт частот | O(n) / O(1) | с подсказками | [hashing/0242-valid-anagram.cpp](hashing/0242-valid-anagram.cpp) |
| 49   | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | хеш-таблица, ключ | O(n·k) / O(n·k) | с разбором → сам после разбора | [hashing/0049-group-anagrams.cpp](hashing/0049-group-anagrams.cpp), [заметки](hashing/0049-group-anagrams.md) |
| 121  | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | один проход | O(n) / O(1) | сам | [one-pass/0121-best-time-to-buy-and-sell-stock.cpp](one-pass/0121-best-time-to-buy-and-sell-stock.cpp) |
| 125  | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | два указателя | O(n) / O(1) | с подсказками | [two-pointers/0125-valid-palindrome.cpp](two-pointers/0125-valid-palindrome.cpp) |
| 167  | [Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | два указателя | O(n) / O(1) | map — сам, указатели — с подсказкой | [two-pointers/0167-two-sum-ii.cpp](two-pointers/0167-two-sum-ii.cpp) |
| 11   | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | два указателя | O(n) / O(1) | сам | [two-pointers/0011-container-with-most-water.cpp](two-pointers/0011-container-with-most-water.cpp) |
| 283  | [Move Zeroes](https://leetcode.com/problems/move-zeroes/) | два указателя | O(n) / O(1) | сам (с подсказкой из плана) | [two-pointers/0283-move-zeroes.cpp](two-pointers/0283-move-zeroes.cpp) |
| 977  | [Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/) | два указателя | O(n) / O(1) | insert — сам, с конца — с подсказкой | [two-pointers/0977-squares-of-a-sorted-array.cpp](two-pointers/0977-squares-of-a-sorted-array.cpp) |

## Статусы

- **сам** — решил без подсказок
- **с подсказками** — нужен был намёк на идею
- **с разбором** — разобрал готовое решение, нужно перерешать
- **сам после разбора** — перерешал сам после разбора

## Как повторять

1. Открыть файл и прочитать только строку «Идея».
2. Решить задачу с нуля в чистом файле, не подглядывая в код.
3. Если решил сам, обновить статус и дату. Если нет, повторить через 2–3 дня.
