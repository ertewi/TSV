import numpy as np

# Функция половинно-быстрого преобразования Фурье (HFFT)
def half_fast_fourier_transform(f):
    N = len(f)
    p1 = int(np.sqrt(N))
    p2 = N // p1

    # Убедимся, что длина N является произведением p1 и p2
    assert p1 * p2 == N, "Длина входного массива должна быть произведением p1 и p2."

    # Шаг 1: Первое преобразование
    A_1 = np.zeros((p1, p2), dtype=complex)
    for k2 in range(p2):
        for k1 in range(p1):
            summation = 0
            for j1 in range(p1):
                j = j1 + p1 * k2
                exponent = -2j * np.pi * k1 * j1 / p1
                summation += f[j] * np.exp(exponent)
            A_1[k1, k2] = summation

    # Шаг 2: Второе преобразование
    A_2 = np.zeros((p1, p2), dtype=complex)
    for k1 in range(p1):
        for k2 in range(p2):
            summation = 0
            for j2 in range(p2):
                j = k1 + p1 * j2
                exponent = -2j * np.pi * k2 * j2 / p2
                summation += A_1[k1, j2] * np.exp(exponent)
            A_2[k1, k2] = summation

    # Преобразование в одномерный массив
    A = np.zeros(N, dtype=complex)
    for k1 in range(p1):
        for k2 in range(p2):
            k = k1 + p1 * k2
            A[k] = A_2[k1, k2]

    return A

# Функция обратного половинно-быстрого преобразования Фурье
def inverse_half_fast_fourier_transform(F):
    N = len(F)
    p1 = int(np.sqrt(N))
    p2 = N // p1

    assert p1 * p2 == N, "Длина входного массива должна быть произведением p1 и p2."

    # Изменение формы F
    F_reshaped = np.zeros((p1, p2), dtype=complex)
    for k1 in range(p1):
        for k2 in range(p2):
            k = k1 + p1 * k2
            F_reshaped[k1, k2] = F[k]

    # Шаг 1: Первое обратное преобразование
    A_inv_1 = np.zeros((p1, p2), dtype=complex)
    for j2 in range(p2):
        for j1 in range(p1):
            summation = 0
            for k2 in range(p2):
                exponent = 2j * np.pi * k2 * j2 / p2
                summation += F_reshaped[j1, k2] * np.exp(exponent)
            A_inv_1[j1, j2] = summation

    # Шаг 2: Второе обратное преобразование
    A_inv_2 = np.zeros((p1, p2), dtype=complex)
    for j1 in range(p1):
        for j2 in range(p2):
            summation = 0
            for k1 in range(p1):
                exponent = 2j * np.pi * k1 * j1 / p1
                summation += A_inv_1[k1, j2] * np.exp(exponent)
            A_inv_2[j1, j2] = summation

    # Преобразование обратно в одномерный массив
    A_inv = np.zeros(N, dtype=complex)
    for j1 in range(p1):
        for j2 in range(p2):
            j = j1 + p1 * j2
            A_inv[j] = A_inv_2[j1, j2] / N

    return A_inv

# Функция свёртки через HFFT
def half_fast_fourier_convolution(f, g):
    N = len(f) + len(g) - 1
    f_padded = np.pad(f, (0, N - len(f)), mode='constant')
    g_padded = np.pad(g, (0, N - len(g)), mode='constant')

    # Преобразование Фурье
    F = half_fast_fourier_transform(f_padded)
    G = half_fast_fourier_transform(g_padded)

    # Произведение преобразований Фурье
    FG = F * G

    # Обратное преобразование
    conv_result = inverse_half_fast_fourier_transform(FG)

    return conv_result

# Пример использования и красивый вывод
f = np.array([1, 2, 3, 4], dtype=complex)
g = np.array([2, 0, 1], dtype=complex)

# Результат свёртки
convolution_result = half_fast_fourier_convolution(f, g)

# Вывод результата с округлением для удобства
print("Результат свёртки (округлён до 2 знаков):")
for i, value in enumerate(np.round(convolution_result, 2)):
    print(f"Элемент {i}: {value}")
