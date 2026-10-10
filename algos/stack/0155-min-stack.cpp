// Min Stack — https://leetcode.com/problems/min-stack/
// Тема: стек
// Идея: второй стек minimum хранит минимумы. В push кладём value в minimum, если
//       value <= minimum.top(). В pop снимаем из minimum, только если снимаемый
//       элемент равен minimum.top(). getMin = minimum.top().
// Сложность: O(1) на каждую операцию, O(n) память
// Заметка: сначала было `>` вместо `<=` — повторный минимум (0,1,0) не попадал в minimum,
//          и после pop терялся. Потом pop без проверки снимал из minimum всегда → падение
//          на пустом стеке. Вариант проще: один stack<pair<int,int>> {значение, минимум}.
// Статус: с подсказками
// Дата: 09.10.2026

class MinStack {
private:
    stack<int> stack1;
    stack<int> minimum;

public:
    MinStack() {

    }

    void push(int value) {
        stack1.push(value);
        if (minimum.empty()) {
            minimum.push(value);
        } else if (value <= minimum.top()) {
            minimum.push(value);
        }
    }

    void pop() {
        if (stack1.top() == minimum.top()) {
            stack1.pop();
            minimum.pop();
        } else {
            stack1.pop();
        }
    }

    int top() {
        return stack1.top();
    }

    int getMin() {
        return minimum.top();
    }
};
