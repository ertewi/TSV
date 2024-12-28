import numpy as np

# Полубыстрое прямое преобразование Фурье
def half_fast_fourier_transform(f):
    N = len(f)
    
    # Поиск оптимальных делителей для разбиения
    p1 = int(np.sqrt(N))
    while N % p1 != 0:
        p1 -= 1
    p2 = N // p1

    # Проверка корректности деления
    assert p1 * p2 == N, "Ошибка: не удалось корректно разделить массив"

    A = np.zeros(N, dtype=complex)

    # Предвычисление экспонент для ускорения
    exp1 = np.exp(-2j * np.pi * np.outer(np.arange(p1), np.arange(p1)) / p1)
    exp2 = np.exp(-2j * np.pi * np.outer(np.arange(p2), np.arange(p2)) / p2)

    # Применение двумерного FFT в разбиении на p1 и p2
    for k1 in range(p1):
        for k2 in range(p2):
            summation = 0
            for j1 in range(p1):
                for j2 in range(p2):
                    j = j1 + p1 * j2
                    exponent = exp1[k1, j1] * exp2[k2, j2]
                    summation += f[j] * exponent

            A[k1 + p1 * k2] = summation
    
    return A

# Полубыстрое обратное преобразование Фурье
def inverse_half_fast_fourier_transform(A):
    N = len(A)
    
    # Поиск оптимальных делителей для разбиения
    p1 = int(np.sqrt(N))
    while N % p1 != 0:
        p1 -= 1
    p2 = N // p1

    # Проверка корректности деления
    assert p1 * p2 == N, "Ошибка: не удалось корректно разделить массив"

    f = np.zeros(N, dtype=complex)

    # Предвычисление экспонент для обратного преобразования
    exp1 = np.exp(2j * np.pi * np.outer(np.arange(p1), np.arange(p1)) / p1)
    exp2 = np.exp(2j * np.pi * np.outer(np.arange(p2), np.arange(p2)) / p2)

    # Применение двумерного обратного FFT в разбиении на p1 и p2
    for j1 in range(p1):
        for j2 in range(p2):
            summation = 0
            for k1 in range(p1):
                for k2 in range(p2):
                    k = k1 + p1 * k2
                    exponent = exp1[j1, k1] * exp2[j2, k2]
                    summation += A[k] * exponent

            f[j1 + p1 * j2] = summation / N  # Нормировка
    
    return f

# Пример использования
f = np.array([1, 2, 3, 4], dtype=complex)

# Прямое полубыстрое преобразование Фурье
result = half_fast_fourier_transform(f)

# Обратное полубыстрое преобразование Фурье
inverse_result = inverse_half_fast_fourier_transform(result)

# Обычное преобразование Фурье для сравнения
fft_result = np.fft.fft(f)
ifft_result = np.fft.ifft(fft_result)

print("Результат полубыстрого преобразования Фурье (ПШФ):")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(result):
    print(f"{i:6} | {val:.4f}")

print("\nРезультат обратного ПШФ:")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(inverse_result):
    print(f"{i:6} | {val:.4f}")

print("\nРезультат обычного FFT для сравнения:")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(fft_result):
    print(f"{i:6} | {val:.4f}")

print("\nРезультат обратного FFT для сравнения:")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(ifft_result):
    print(f"{i:6} | {val:.4f}")
