// Valid Anagram — https://leetcode.com/problems/valid-anagram/
// Тема: хеш-таблица, подсчёт частот
// Идея: разная длина — сразу false; иначе считаем частоты букв в обеих строках и сравниваем.
// Сложность: O(n) время, O(1) память — алфавит ограничен 26 буквами
// Заметка: можно одной мапой — по s прибавлять, по t вычитать, в конце все значения нули.
// Статус: с подсказками
// Дата: 22.09.2026

class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.length() != t.length()) return false;
        unordered_map<char,int> seen;
        unordered_map<char,int> seen1;
        for (int i=0;i<s.length();i++){
            seen[s[i]]++;
        } 
        for (int i=0;i<t.length();i++){
            seen1[t[i]]++;
        }
        if (seen1==seen){
            return true;
        }
        return false;
    }
};
