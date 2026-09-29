// Best Time to Buy and Sell Stock — https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Тема: один проход
// Идея: идём слева направо, держим минимум из увиденного;
//       на каждом шаге проверяем прибыль от продажи сегодня.
// Сложность: O(n) время, O(1) память
// Заметка: имена max и min совпадают с std::max / std::min — лучше best и minPrice.
// Статус: сам
// Дата: 23.09.2026

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int max=0;
        int min=INT_MAX;
        for (int price:prices){
            if(price<min){
                min=price;
            }
            if (max<price-min){
                max=price-min;
            }
        }
        return max;
    }
};
