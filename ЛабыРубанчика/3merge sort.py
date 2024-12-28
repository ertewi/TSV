def merge(arr, left, mid, right):
    # временные массивы для левой и правой частей
    left_part = arr[left:mid + 1]
    right_part = arr[mid + 1:right + 1]

    i = 0
    j = 0
    k = left

    # Объединяем временные массивы обратно в основной массив arr
    while i < len(left_part) and j < len(right_part):
        if left_part[i] <= right_part[j]:
            arr[k] = left_part[i]
            i += 1
        else:
            arr[k] = right_part[j]
            j += 1
        k += 1

    # Копируем оставшиеся элементы из left_part, если они есть
    while i < len(left_part):
        arr[k] = left_part[i]
        i += 1
        k += 1

    # Копируем оставшиеся элементы из right_part, если они есть
    while j < len(right_part):
        arr[k] = right_part[j]
        j += 1
        k += 1

def recursive_merge_sort(arr, left=None, right=None):
    if left is None and right is None:
        left = 0
        right = len(arr) - 1

    # если массив состоит из одного элемента
    if left >= right:
        return

    # Находим середину массива
    mid = (left + right) // 2

    # Рекурсивно сортируем левую и правую части
    recursive_merge_sort(arr, left, mid)
    recursive_merge_sort(arr, mid + 1, right)

    # Сливаем отсортированные части
    merge(arr, left, mid, right)

arr = [38, 27, 43, 3, 9, 82, 10]
print("Исходный массив:", arr)
recursive_merge_sort(arr)
print("Отсортированный массив:", arr)
