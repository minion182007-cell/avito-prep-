// Implement Queue using Stacks — https://leetcode.com/problems/implement-queue-using-stacks/
// Тема: стек, очередь
// Идея: два стека: in — для push, out — для pop/peek. Перекладывать из in в out
//       только когда out пуст: порядок переворачивается и получается очередь (FIFO).
// Сложность: амортизированно O(1) на операцию (каждый элемент перекладывается
//            один раз), O(n) память
// Заметка: сначала перекладывал в push — элементы застревали в in, peek падал на пустом out.
//          Перекладка должна быть там, где берём (pop/peek). return перед out.pop() —
//          строка после return не выполняется. empty = оба стека пусты.
// Статус: с подсказками
// Дата: 07.10.2026

class MyQueue {
private:
    stack<int> in;
    stack<int> out;
    void move() {
        if (out.empty()) {
            while (!in.empty()) {
                out.push(in.top());
                in.pop();
            }
        }
    }
public:
    MyQueue() {
    }

    void push(int x) {
        in.push(x);
    }

    int pop() {
        move();
        int t=out.top();
        out.pop();
        return t;
    }

    int peek() {
        move();
        return out.top();
    }

    bool empty() {
        return out.empty() && in.empty();
    }
};
