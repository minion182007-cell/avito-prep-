// Palindrome Linked List — https://leetcode.com/problems/palindrome-linked-list/
// Тема: связный список, два указателя, разворот списка
// Идея:
//   Решение 1 (твое, O(n) память): скопировать все значения в vector, сравнить с двух концов.
//   Решение 2 (эталон, O(1) память):
//     1. Найти середину списка (slow и fast).
//     2. Развернуть вторую половину (как в 206 Reverse Linked List).
//     3. Сравнить первую половину (head) и развернутую вторую.
// Сложность:
//   Время: O(n)
//   Память: O(1) для решения с разворотом, O(n) для решения с vector.
// Заметка:
//   Главная ловушка на собеседовании — спросят решение за O(1) по памяти.
//   Нужно уметь делать in-place разворот (prev, curr, next).
//   Если меняем список, не забудь восстановить его обратно (хотя в LeetCode это не обязательно, но на собеседовании спросят).
// Статус: сам (Accepted, C++, 93/93)
// Дата: 09.11.2026
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    bool isPalindrome(ListNode *head) {
    ListNode *second =head;
    vector<int> poli;
    while (second!= nullptr) {
        poli.push_back(second->val);
        second = second->next;
    }
    int left = 0;
    int right = poli.size() - 1;
    while (left<right) {
        if (poli[left]!=poli[right]) {
            return false;
        }
        left++;
        right--;
    }


    return true;
}