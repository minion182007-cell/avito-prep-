// Remove Nth Node From End of List — https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Тема: связный список, два указателя (fast и slow) + dummy узел
// Идея: чтобы удалить n-й элемент с конца за один проход, запускаем два указателя с зазором в n+1.
//       Fast уходит вперед на n+1 шагов. Затем оба двигаются, пока fast не станет nullptr.
//       В этот момент slow стоит ровно на узле ПЕРЕД удаляемым. Меняем slow->next = slow->next->next.
// Сложность: O(n) время (один проход), O(1) память.
// Заметка: dummy узел (созданный через new) обязателен по двум причинам:
//          1) Если удаляем первый элемент (head), slow будет указывать на dummy, и мы легко перепрыгнем head.
//          2) Платформа LeetCode на C++ требует удалить dummy через delete перед возвратом,
//             иначе будет heap-use-after-free (если вернуть dummy) или Memory Limit (если забыть delete).
//          Главная ловушка в цикле: условие должно быть fast != nullptr, а НЕ fast->next != nullptr.
//          Если использовать fast->next, fast остановится на последнем узле, slow не дойдет до конца,
//          и удалится не тот элемент (например, для [1..5], n=2 удалится 3 вместо 4).
//          Возвращаем res = dummy->next, а не head (head мог быть удален).
// Статус: c подсказкой, но логику понял сам (Accepted, C++, 100% Beats)
// Дата: 07.11.2026

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

ListNode* removeNthFromEnd(ListNode* head, int n) {

    ListNode *dummy=new ListNode(0);
    dummy->next=head;
    ListNode * fasts=dummy;
    ListNode * slow=dummy;
    while (n>=0) {
        fasts=fasts->next;
        n--;
    }

    while (fasts!=nullptr ) {
        slow=slow->next;
        fasts=fasts->next;
    }
    slow->next=slow->next->next;
    ListNode *res=dummy->next;
    delete dummy;
    return res;
}
