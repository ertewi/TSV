import numpy as np

def dft(f):
    N = len(f)
    A = np.zeros(N, dtype=complex)
    for k in range(N):
        for j in range(N):
            A[k] += f[j] * np.exp(-2j * np.pi * k * j / N)
        A[k] /= N
    return A

f = [1, 2, 3, 4]  
A = dft(f)
print("Результат ДПФ:", A)
