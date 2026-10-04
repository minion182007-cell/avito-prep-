// 3Sum — https://leetcode.com/problems/3sum/
// Тема: два указателя (сортировка + Two Sum II внутри цикла)
// Идея: отсортировать; фиксируем nums[i] и ищем справа пару с суммой -nums[i]
//       двумя указателями. Если nums[i] > 0 — дальше нулей не будет, break.
//       Дубли пропускаем в двух местах: одинаковые nums[i] во внешнем цикле
//       и одинаковые left/right после найденной тройки.
// Сложность: O(n^2) время, O(1) доп. память (не считая ответа и сортировки)
// Заметка: sum считать ВНУТРИ while — иначе решения по устаревшей сумме
//          (на [-3,0,1,2] тройка теряется). right лучше n-1, а не nums.size()-1.
// Статус: с подсказкой (идея + дубли)
// Дата: 04.10.2026

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> values;

        sort(nums.begin(),nums.end());
        int n = nums.size();
        for (int i=0;i<n-2;++i) {
                if (nums[i]>0) break;
                int left=i+1;
                int right=nums.size()-1;
                int target=-nums[i];
                if (i>0 && nums[i]==nums[i-1]) {
                        continue;
                }
                while (left<right) {
                    int sum =nums[left]+nums[right];
                    if (sum==target) {
                        values.push_back({nums[i],nums[left],nums[right]});
                        while (left<right && nums[left]==nums[left+1])left++;
                        while (left<right && nums[right]==nums[right-1])right--;
                        left++;
                        right--;
                    }else if(sum<target){
                        left++;
                    }else{
                        right--;
                    }
                }
            }
        return values;
}
};
