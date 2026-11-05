// Middle of the Linked List — https://leetcode.com/problems/middle-of-the-linked-list/
// Тема: связный список (быстрый и медленный указатели)
// Идея: slow идёт на 1 узел, fast — на 2. Когда fast дошёл до конца, slow прошёл ровно
//       половину и стоит на середине. Цикл: пока fast != nullptr && fast->next != nullptr.
//       При чётной длине slow остаётся на второй из двух середин — как требует условие.
// Сложность: O(n) время, O(1) память
// Заметка: нужны ОБЕ проверки: fast прыгает через узел, и fast->next->next при
//          fast->next == nullptr — обращение по пустому адресу. Порядок важен: сначала
//          fast != nullptr, иначе fast->next упадёт; && не вычисляет вторую часть, если
//          первая false.
// Статус: сам (приём и план дня даны заранее, вопросов не было)
// Дата: 05.11.2026

class Solution {
public:
    ListNode* middleNode(ListNode* head) {
        ListNode* slow = head;
        ListNode* fast = head;
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
        }
        return slow;
    }
};
