import numpy as np

# Прямое дискретное преобразование Фурье (O(n^2))
def DFT(x):
    n = len(x)
    X = np.zeros(n, dtype=complex)
    for k in range(n):
        for t in range(n):
            X[k] += x[t] * np.exp(-2j * np.pi * k * t / n)
    return X

# Обратное дискретное преобразование Фурье (O(n^2))
def IDFT(X):
    n = len(X)
    x = np.zeros(n, dtype=complex)
    for t in range(n):
        for k in range(n):
            x[t] += X[k] * np.exp(2j * np.pi * k * t / n)
    x /= n  # деление на n происходит один раз после цикла
    return x

# Функция свёртки через ДПФ
def convolution_dft(signal, kernel):

    n = len(signal) + len(kernel) - 1
  
    signal_padded = np.pad(signal, (0, n - len(signal)), mode='constant')
    kernel_padded = np.pad(kernel, (0, n - len(kernel)), mode='constant')
    
    signal_dft = DFT(signal_padded)
    kernel_dft = DFT(kernel_padded)
    
    result_dft = signal_dft * kernel_dft
    
    result = IDFT(result_dft)
    
    return np.real(result)

signal = [1, 2, 3, 4, 5]
kernel = [1, 0, -1]

result = convolution_dft(signal, kernel)

print("Сигнал:", signal)
print("Ядро:", kernel)
print("Результат свёртки через ДПФ:", result)

print("\nДополнение нулями:")
print(f"Дополненный сигнал: {np.pad(signal, (0, len(result) - len(signal)), mode='constant')}")
print(f"Дополненное ядро: {np.pad(kernel, (0, len(result) - len(kernel)), mode='constant')}")

print("\nПеремножение в частотной области:")
signal_dft = DFT(np.pad(signal, (0, len(result) - len(signal)), mode='constant'))
kernel_dft = DFT(np.pad(kernel, (0, len(result) - len(kernel)), mode='constant'))
print("Спектр сигнала:", signal_dft)
print("Спектр ядра:", kernel_dft)
print("\nОбратное ДПФ:")
result_dft = signal_dft * kernel_dft
result_time = IDFT(result_dft)
print("Результат после обратного ДПФ:", result_time)
