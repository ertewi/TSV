def multiply_large_numbers(A: str, B: str):
    # Вычисляем количество разрядов
    Na = len(A)
    Nb = len(B)
    
    # Инициализация массивов для хранения чисел
    Array_A = [int(A[i]) for i in range(Na)]
    Array_B = [int(B[i]) for i in range(Nb)]
    
    # Результирующий массив, максимальная длина которого будет Na + Nb
    Array_C = [0] * (Na + Nb)
    
    # Умножение двух чисел в виде массивов цифр
    for i in range(Na):
        for j in range(Nb):
            Array_C[i + j + 1] += Array_A[i] * Array_B[j]

    # Обработка переносов
    for i in range(Na + Nb - 1, 0, -1):
        if Array_C[i] >= 10:
            Array_C[i - 1] += Array_C[i] // 10
            Array_C[i] %= 10
    
    # Найдем первый ненулевой разряд
    k = 0
    while k < len(Array_C) and Array_C[k] == 0:
        k += 1
    
    # Вывод результата
    if k == len(Array_C):  # Если все цифры нулевые
        print("0")
    else:
        result = ''.join(map(str, Array_C[k:]))
        print(result)
    
# Пример использования
A = "123"
B = "456"
multiply_large_numbers(A, B)
