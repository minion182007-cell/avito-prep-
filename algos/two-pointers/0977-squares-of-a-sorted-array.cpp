// Squares of a Sorted Array — https://leetcode.com/problems/squares-of-a-sorted-array/
// Тема: два указателя (с двух концов)
// Идея: самый большой квадрат всегда на одном из краёв (отрицательные в квадрате тоже большие).
//       Указатели на оба конца, берём больший по модулю и пишем результат с конца.
// Сложность: версия 1 — O(n^2) из-за insert в начало; версия 2 — O(n) время, O(1) доп. память
// Заметка: insert в начало вектора сдвигает все элементы — O(n) на каждую вставку.
//          Создать vector(n) сразу и писать по индексу pos с конца.
// Статус: версия 1 — сам (с подсказкой из плана); версия 2 — с подсказкой
// Дата: 02.10.2026

// ---------- Версия 1: insert в начало, O(n^2) ----------
class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
    vector<int> sq2;
    int left=0;
    int right=size(nums)-1;
    while (left<right) {
        if (abs(nums[left])<=abs(nums[right])) {
            sq2.insert(sq2.begin(),nums[right]*nums[right]);
            right--;
        }else{
            sq2.insert(sq2.begin(),nums[left]*nums[left]);
            left++;
        }
    }
    sq2.insert(sq2.begin(),nums[left]*nums[left]);
    return sq2;
    }
};

// ---------- Версия 2: заполнение с конца, O(n) ----------
class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
    int n=size(nums);
    vector<int> sq2(n);
    int pos=n-1;
    int left=0;
    int right=n-1;
    while (left<=right) {
        if (abs(nums[left])<=abs(nums[right])) {
            sq2[pos]=nums[right]*nums[right];
            right--;

        }else{
            sq2[pos]=nums[left]*nums[left];
            left++;
        }
        pos--;
    }
    return sq2;
    }
};
