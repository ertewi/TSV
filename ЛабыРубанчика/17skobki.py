import sys

# Функция для вычисления минимального числа операций для умножения матриц
def calculate_min_operations(dimensions):
    num_matrices = len(dimensions) - 1  # количество матриц
    min_operations = [[0] * num_matrices for _ in range(num_matrices)]  # таблица для хранения минимальных операций
    split_points = [[0] * num_matrices for _ in range(num_matrices)]  # таблица для хранения позиций разбиений

    # l - длина цепочки матриц, начиная с 2
    for chain_length in range(2, num_matrices + 1):
        for start in range(num_matrices - chain_length + 1):
            end = start + chain_length - 1
            min_operations[start][end] = sys.maxsize  # инициализируем большим значением

            # k - точка разбиения цепочки матриц на две подзадачи
            for split in range(start, end):
                operations = (min_operations[start][split] + 
                              min_operations[split + 1][end] + 
                              dimensions[start] * dimensions[split + 1] * dimensions[end + 1])

                # Если текущее количество операций меньше предыдущего, обновляем минимальное значение
                if operations < min_operations[start][end]:
                    min_operations[start][end] = operations
                    split_points[start][end] = split

    return min_operations, split_points

# Функция для восстановления оптимального порядка умножения
def get_optimal_parenthesization(split_points, start, end):
    if start == end:
        return f"A{start + 1}"  # Одиночная матрица
    else:
        left = get_optimal_parenthesization(split_points, start, split_points[start][end])
        right = get_optimal_parenthesization(split_points, split_points[start][end] + 1, end)
        return f"({left} x {right})"  # Возвращаем строку с расстановкой скобок

# Пример использования
dimensions = [10, 20, 50, 1, 100]  # Размерности матриц
min_operations, split_points = calculate_min_operations(dimensions)

optimal_order = get_optimal_parenthesization(split_points, 0, len(dimensions) - 2)
min_operations_count = min_operations[0][len(dimensions) - 2]

print(f"Минимальное количество операций: {min_operations_count}")
print(f"Оптимальный порядок умножения: {optimal_order}")
