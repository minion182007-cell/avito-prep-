// Evaluate Reverse Polish Notation — https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Тема: стек
// Идея: число → в стек. Оператор → снять два числа: первым снимается ПРАВЫЙ операнд (n1),
//       вторым левый (n2); положить n2 op n1. В конце в стеке лежит ответ.
// Сложность: O(n) время, O(n) память
// Заметка: строку в число — stoi(s) (аналог strconv.Atoi). Оператор проверять сравнением
//          всей строки (i == "-"), а не по первому символу: "-11" — это число.
//          i[0] — char ('-'), "-" — строка: сравнивать их нельзя. Порядок важен для - и /.
//          На LeetCode функция должна быть внутри class Solution { public: ... }.
// Статус: с подсказками
// Дата: 10.10.2026

class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<int> token;
        for (string i : tokens) {
            if (i == "+" || i == "-" || i == "*" || i == "/") {
                int n1 = token.top();
                token.pop();
                int n2 = token.top();
                token.pop();
                if (i == "+") {
                    token.push(n1 + n2);
                } else if (i == "*") {
                    token.push(n1 * n2);
                } else if (i == "/") {
                    token.push(n2 / n1);
                } else {
                    token.push(n2 - n1);
                }
            } else {
                token.push(stoi(i));
            }
        }
        return token.top();
    }
};
