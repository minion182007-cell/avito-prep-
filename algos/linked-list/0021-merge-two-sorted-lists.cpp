// Merge Two Sorted Lists — https://leetcode.com/problems/merge-two-sorted-lists/
// Тема: связный список (фиктивный узел)
// Идея: узел-пустышка dummy в начале нового списка + указатель tail на его последний узел.
//       Пока ОБА списка не пусты: прицепить к tail меньший узел (tail->next = ...),
//       сдвинуть этот список, сдвинуть tail. Потом прицепить остаток другого целиком.
//       Ответ — dummy.next.
// Сложность: O(n + m) время, O(1) память (узлы не создаются, только перецепляются)
// Заметка: сначала пытался переставлять стрелки внутри самих списков — терялись узлы
//          и падало на пустом списке. Потом ошибка в условиях: проверял list->next вместо
//          самого list и || вместо &&. «Список не пуст» = p != nullptr, а не p->next != nullptr;
//          перед p->... убедиться, что p не nullptr. Оба пустых на входе — прицепится nullptr,
//          ответ пустой, отдельная проверка не нужна. dummy — сам узел (dummy.next),
//          tail — адрес узла (tail->next).
// Статус: с подсказками (приём dummy и план даны, код и исправления условий — сам)
// Дата: 05.11.2026

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode zp;
        ListNode* tail = &zp;
        while (list1 != nullptr && list2 != nullptr) {
            if (list1->val <= list2->val) {
                tail->next = list1;
                list1 = list1->next;
            } else {
                tail->next = list2;
                list2 = list2->next;
            }
            tail = tail->next;
        }
        if (list1 == nullptr) {
            tail->next = list2;
        } else {
            tail->next = list1;
        }
        return zp.next;
    }
};
