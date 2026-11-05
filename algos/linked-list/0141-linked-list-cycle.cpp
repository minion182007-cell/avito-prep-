// Linked List Cycle — https://leetcode.com/problems/linked-list-cycle/
// Тема: связный список (быстрый и медленный указатели, алгоритм Флойда)
// Идея: slow на 1 узел, fast на 2. Если цикла нет — fast дойдёт до nullptr и цикл while
//       закончится → false. Если цикл есть — fast зайдёт на круг и догонит slow сзади:
//       после очередного шага slow == fast → true.
// Сложность: O(n) время, O(1) память (без set для запоминания узлов)
// Заметка: pos в примерах LeetCode функции не передаётся — это только описание теста.
//          Сравниваем адреса узлов (slow == fast), а не значения: одинаковые числа
//          бывают в разных узлах. Проверку делать ПОСЛЕ шагов: до шагов slow == fast == head.
//          Обычный проход «пока cur != nullptr» на списке с циклом никогда не кончится.
// Статус: сам (по условию, план дня не читал)
// Дата: 05.11.2026

class Solution {
public:
    bool hasCycle(ListNode* head) {
        ListNode* slow = head;
        ListNode* fast = head;
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
            if (fast == slow) {
                return true;
            }
        }
        return false;
    }
};
