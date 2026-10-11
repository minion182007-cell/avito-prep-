// Daily Temperatures — https://leetcode.com/problems/daily-temperatures/
// Тема: монотонный стек
// Идея: в стеке индексы дней, которые ещё ждут потепления (температуры по убыванию).
//       Новый день i: ПОКА он теплее дня на вершине — тот дождался: j = top,
//       answer[j] = i - j, pop. Потом push(i). Кто остался в стеке — ответ 0.
// Сложность: O(n) время (каждый индекс кладётся и снимается один раз), O(n) память
// Заметка: int i без = 0 — мусор (в C++ локальные не обнуляются, в отличие от Go).
//          Знак: нужен temperatures[i] > temperatures[top], с < ищет более холодный.
//          while, а не if: один тёплый день закрывает сразу несколько ожидающих.
//          answer/pop — внутри while, а не после него.
// Статус: с подсказками
// Дата: 11.10.2026

class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        stack<int> temp;
        int n = temperatures.size();
        vector<int> answer(n, 0);
        for (int i = 0; i < n; ++i) {
            while (!temp.empty() && temperatures[i] > temperatures[temp.top()]) {
                int j = temp.top();
                answer[j] = i - j;
                temp.pop();
            }
            temp.push(i);
        }
        return answer;
    }
};
