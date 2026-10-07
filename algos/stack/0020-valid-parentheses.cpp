// Valid Parentheses — https://leetcode.com/problems/valid-parentheses/
// Тема: стек
// Идея: открывающие скобки кладём в стек; закрывающая должна совпасть с вершиной
//       (иначе false); в конце стек должен быть пуст.
// Сложность: O(n) время, O(n) память
// Заметка: сначала было flag - i → разность всегда отрицательная, всё возвращало false.
//          Пара определяется по разности ASCII: ')'-'(' = 1, ']'-'[' = '}'-'{' = 2.
//          Читаемее писать сами символы '(' ')' вместо 40, 41 — на повторе писать так.
//          Перед top()/pop() всегда проверять empty().
// Статус: с подсказкой
// Дата: 07.10.2026

class Solution {
public:
    bool isValid(string s) {
    stack<char> str;
    for (char i : s) {
        if (i == 40 || i == 91 ||i == 123  ) {
            str.push(i);
        }else if (i == 41 || i == 93 ||i == 125  ) {
            if (str.empty()) return false;
            char flag=str.top();
            str.pop();
            int sum=i-flag;
            if (sum!=1 && sum!=2) {
                return false;
            }
        }
    }
    return str.empty();
}
};
