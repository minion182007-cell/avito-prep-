// Valid Palindrome — https://leetcode.com/problems/valid-palindrome/
// Тема: два указателя
// Идея: два указателя навстречу; всё, кроме букв и цифр, пропускаем,
//       сравниваем с приведением регистра.
// Версии: 1) очищенная строка + разворот — O(n) время, O(n) память
//         2) два указателя               — O(n) время, O(1) память
// Ловушки: проверка left < right во вложенных циклах; сдвиг обоих указателей после сравнения.
// Статус: с подсказками
// Дата: 24.09.2026

// ---------- Версия 1: очищенная строка, O(n) памяти ----------

class Solution {
public:
    bool isPalindrome(string s) {
        string clean;
        for (char c: s){
            if (isalnum(c)){
                clean+=tolower(c);
            }
        }
        string result = clean;
        reverse(result.begin(),result.end());
        return result==clean;
    }
};

// ---------- Версия 2: два указателя, O(1) памяти ----------

class Solution {
public:
    bool isPalindrome(string s) {
       int left=0;
       int right=s.size()-1;
       while(left<right){
            while(left<right && !isalnum(s[left])){
                left+=1;
            }
            while(left<right && !isalnum(s[right])){
                right-=1;
            }
            if (tolower(s[left])!=tolower(s[right])){
                    return false;
            }
            left+=1;
            right-=1;
       }
       return true;
    }
};
