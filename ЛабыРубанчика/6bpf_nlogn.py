import numpy as np

# Прямое быстрое преобразование Фурье (FFT)
def fft_1d(x):
    N = len(x)
    if N == 0:
        return []
    if N == 1:
        return x
    
    # Декомпозиция на четные и нечетные индексы
    even = fft_1d(x[::2])
    odd = fft_1d(x[1::2])
    
    # Расчет комплексных экспонент только для половины N
    factor = np.exp(-2j * np.pi * np.arange(N // 2) / N)
    
    # Комбинирование результатов
    return np.concatenate([even + factor * odd, even - factor * odd])

# Обратное быстрое преобразование Фурье (IFFT)
def ifft_1d(X):
    N = len(X)
    if N == 0:
        return []
    if N == 1:
        return X
    
    # Декомпозиция на четные и нечетные индексы
    even = ifft_1d(X[::2])
    odd = ifft_1d(X[1::2])
    
    # Расчет комплексных экспонент для обратного преобразования (со знаком '+')
    factor = np.exp(2j * np.pi * np.arange(N // 2) / N)
    
    # Комбинирование результатов
    result = np.concatenate([even + factor * odd, even - factor * odd])
    
    # Нормализация для обратного преобразования
    return result / N

# Пример одномерного массива
x = np.array([1, 2, 3, 4], dtype=complex)

# Вычисление FFT без библиотеки
fft_result_custom = fft_1d(x)

# Вычисление IFFT без библиотеки
ifft_result_custom = ifft_1d(fft_result_custom)

# Вычисление FFT и IFFT с использованием библиотеки numpy
fft_result_numpy = np.fft.fft(x)
ifft_result_numpy = np.fft.ifft(fft_result_numpy)

# Вывод результатов
print("Результат FFT (кастомная реализация, 1D):")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(fft_result_custom):
    print(f"{i:6} | {val:.4f}")

print("\nРезультат IFFT (кастомная реализация, 1D):")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(ifft_result_custom):
    print(f"{i:6} | {val:.4f}")

print("\nРезультат FFT (реализация numpy, 1D):")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(fft_result_numpy):
    print(f"{i:6} | {val:.4f}")

print("\nРезультат IFFT (реализация numpy, 1D):")
print("Индекс | Значение")
print("-------------------")
for i, val in enumerate(ifft_result_numpy):
    print(f"{i:6} | {val:.4f}")
