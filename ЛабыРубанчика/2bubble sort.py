def recursive_bubble_sort(arr, n=None):
    # Инициализация размера массива
    if n is None:
        n = len(arr)

    # если размер массива равен 1
    if n == 1:
        return

    # Пройти по массиву и поменять местами соседние элементы, если они не по порядку
    for i in range(n - 1):
        if arr[i] > arr[i + 1]:
            arr[i], arr[i + 1] = arr[i + 1], arr[i]

    # Рекурсивный вызов для оставшейся части массива
    recursive_bubble_sort(arr, n - 1)

arr = [64, 22, 25, 12, 22, 11, 90]
print("Исходный массив:", arr)
recursive_bubble_sort(arr)
print("Отсортированный массив:", arr)
