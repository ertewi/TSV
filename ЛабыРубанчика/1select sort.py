def recursive_selection_sort(arr, n=None, index=0):
    # Инициализация размера массива
    if n is None:
        n = len(arr)
    
    # Базовый случай: если достигнут конец массива
    if index == n:
        return
    
    # Найти минимальный элемент в неотсортированной части массива
    min_index = index
    for i in range(index + 1, n):
        if arr[i] < arr[min_index]:
            min_index = i
    
    # Поменять местами текущий элемент с найденным минимальным
    arr[index], arr[min_index] = arr[min_index], arr[index]
    
    # Рекурсивный вызов для сортировки оставшейся части массива
    recursive_selection_sort(arr, n, index + 1)

# Пример использования:
arr = [64, 25, 12, 22, 11]
print("Исходный массив:", arr)
recursive_selection_sort(arr)
print("Отсортированный массив:", arr)
