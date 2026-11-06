// Remove Duplicates from Sorted List — https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// Тема: связный список, один указатель
// Идея: список отсортирован → дубликаты идут подряд. Держим cur и сравниваем cur.Val с cur.Next.Val.
//       Если равны — перепрыгиваем узел: cur.Next = cur.Next.Next (сам cur не двигаем, вдруг следующий тоже дубль).
//       Если не равны — двигаем cur = cur.Next. Идём, пока cur != nil && cur.Next != nil.
// Сложность: O(n) время, O(1) память (переставляем указатели, новые узлы не создаём)
// Заметка: сначала забыл '=' в "cur->next = cur->next->next" — выражение вычислялось и выбрасывалось,
//          ничего не удалялось. В C++ это компилируется (warning), в Go — ошибка "evaluated but not used".
//          Вторая ловушка — внутренний while без проверки cur.Next != nil → разыменование nil → паника.
//          Порядок в && важен: сначала cur.Next != nil, потом cur.Next.Val — short-circuit спасает.
//          Приём «внутренний while ест все дубликаты подряд, потом сдвиг» эквивалентен «if + не двигать cur»,
//          оба O(n)/O(1). Главное — не двигать cur, пока есть дубли после него.
// Статус: сам (Accepted, C++, 168/168)
// Дата: 06.11.2026

struct ListNode {
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr) {
    }

    ListNode(int x) : val(x), next(nullptr) {
    }

    ListNode(int x, ListNode *next) : val(x), next(next) {
    }
};

ListNode *deleteDuplicates(ListNode *head) {
    ListNode *cur = head;
    while (cur != nullptr && cur->next != nullptr) {
        while (cur->next != nullptr &&cur->val == cur->next->val) {
            cur->next = cur->next->next;
        }
        cur = cur->next;
    }
    return head;
}
