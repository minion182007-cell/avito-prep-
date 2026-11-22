// Merge Intervals — https://leetcode.com/problems/merge-intervals/
// Тема: сортировка, слияние отрезков
// Идея: сортируем отрезки по началу, тогда пересекающиеся стоят подряд. Кладём первый
//       в ответ, дальше сравниваем каждый с последним слитым: если начало <= его конца,
//       расширяем конец до max, иначе добавляем отрезок как новый.
// Сложность: O(n log n) время (сортировка), O(n) память на ответ
// Заметка: ловушки — условие <=, а не <, иначе [1,4] и [4,5] не сольются;
//          конец берём через max, иначе вложенный отрезок [1,10],[2,3] укоротит ответ до [1,3].
//          sort для vector<vector<int>> сравнивает лексикографически, то есть по началу.
// Статус: сам
// Дата: 20.11.2026
class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
    if (intervals.empty()) return {};
    vector<vector<int>> answer;
    sort(intervals.begin(),intervals.end());
    answer.push_back(intervals[0]);
    for (int i=1;i<intervals.size();++i) {
        vector<int>& Last=answer.back();
        if (intervals[i][0]<=Last[1]) {
            Last[1]=max(intervals[i][1],Last[1]);
        }else {
            answer.push_back(intervals[i]);
        }
    }
    return answer;
}
};