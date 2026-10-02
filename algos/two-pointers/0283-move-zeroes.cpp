// Move Zeroes — https://leetcode.com/problems/move-zeroes/
// Тема: два указателя (в одну сторону)
// Идея: zero — граница записи: слева от неё уже стоят все найденные ненулевые в исходном порядке.
//       Встретил ненулевой — swap с позицией zero и сдвинуть границу. Нули сами уходят вправо.
// Сложность: O(n) время, O(1) память
// Заметка: num — это индекс, а не число: лучше i; zero лучше назвать write/pos.
//          swap при zero == num меняет элемент сам с собой — безвредно.
// Статус: сам (с подсказкой из плана «указатели в одну сторону»)
// Дата: 02.10.2026

class Solution {
public:
    void moveZeroes(vector<int>& nums) {
    int zero = 0;
    for (int num = 0; num < nums.size(); ++num) {
        if (nums[num] != 0) {
            swap(nums[zero], nums[num]);
            ++zero;
        }
    }
}
};
