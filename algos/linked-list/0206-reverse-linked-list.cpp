// Reverse Linked List — https://leetcode.com/problems/reverse-linked-list/
// Тема: связный список (разворот на месте)
// Идея: три указателя. prev — уже развёрнутая часть (сначала nullptr), cur — текущий узел,
//       next — запасной. Пока cur != nullptr: запомнить next = cur->next → развернуть
//       стрелку cur->next = prev → prev = cur → cur = next. Ответ — prev (новая голова).
// Сложность: O(n) время, O(1) память
// Заметка: head — не массив, а адрес первого узла; [1,2,3] на LeetCode — только запись.
//          Указатель = адрес. `cur = next` — поменять записку в руках (куда смотрю),
//          `cur->next = prev` — поменять записку в узле (развернуть стрелку): это единственная
//          строка, которая меняет список. Сначала запомнить next, иначе хвост потерян.
//          Пустой список и один узел работают без отдельных проверок. Рисовать стрелки!
// Статус: с подсказками (план из 4 шагов и трассировка даны, код — сам)
// Дата: 03.11.2026

/**
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 * };
 */
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode* cur = head;
        ListNode* next = nullptr;
        ListNode* prev = nullptr;
        while (cur != nullptr) {
            next = cur->next;
            cur->next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }
};
