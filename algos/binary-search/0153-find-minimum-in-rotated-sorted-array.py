# Find Minimum in Rotated Sorted Array — https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
# Тема: бинарный поиск (поиск границы в повёрнутом массиве)
# Идея: сравнивать nums[mid] с nums[right]. Если nums[mid] > nums[right] — «обрыв» (минимум)
#       правее mid, и сам mid точно не минимум → left = mid + 1. Иначе минимум — mid или левее
#       → right = mid (mid не выбрасываем). Шаблон 278: while left < right, ответ nums[left].
# Сложность: O(log n) время, O(1) память
# Заметка: сравнивать именно с nums[right], а не с nums[left]: правый край всегда говорит,
#          в какой половине обрыв. Решено на Python; в C++ mid = left + (right - left) / 2
#          (в Python переполнения нет, поэтому (left + right) // 2 допустимо).
# Статус: с подсказками (план сравнения дан, действия выбрал сам)
# Дата: 02.11.2026

class Solution:
    def findMin(self, nums: list[int]) -> int:
        left, right = 0, len(nums) - 1
        while left < right:
            mid = (left + right) // 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        return nums[left]
