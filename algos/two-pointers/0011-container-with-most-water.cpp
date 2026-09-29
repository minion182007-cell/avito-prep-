// Container With Most Water — https://leetcode.com/problems/container-with-most-water/
// Тема: два указателя
// Идея: площадь ограничена меньшей высотой; двигаем меньшую —
//       сдвиг большей уменьшит ширину, а высоту не поднимет, площадь только упадёт.
// Сложность: O(n) время, O(1) память
// Заметка: abs не нужен — внутри цикла right > left; площадь считать один раз и брать max.
// Статус: сам
// Дата: 29.09.2026

class Solution {
public:
    int maxArea(vector<int>& height) {
    int left=0;
    int right=height.size()-1;
    int best=0;
    while (left<right) {
        if (best<abs(left-right)*min(height[left],height[right])){
            best=abs(left-right)*min(height[left],height[right]);
        }
        if( height[left]<height[right]) {
            left+=1;
        }else {
            right-=1;
        }

    }
    return best;
}
};
