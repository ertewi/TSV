def karatsuba(num1, num2):
    # Преобразуем числа в строки для дальнейшей работы
    str_num1, str_num2 = str(num1), str(num2)

    # Выравниваем длину чисел для корректного деления на части
    length = max(len(str_num1), len(str_num2))

    # Базовый случай рекурсии - если длина числа 1, выполняем обычное умножение
    if length == 1:
        return num1 * num2

    # Разделяем числа пополам
    half_length = length // 2

    # Разделяем числа на старшие и младшие разряды
    high1, low1 = divmod(num1, 10**half_length)
    high2, low2 = divmod(num2, 10**half_length)

    # Рекурсивные вызовы Карацубы для получения промежуточных произведений
    z_low = karatsuba(low1, low2)               # low1 * low2
    z_high = karatsuba(high1, high2)           # high1 * high2
    z_cross = karatsuba(low1 + high1, low2 + high2)  # (low1 + high1) * (low2 + high2)

    # Комбинируем результат по формуле Карацубы
    return z_high * 10**(2 * half_length) + (z_cross - z_high - z_low) * 10**half_length + z_low

x = 15
y = 15
result = karatsuba(x, y)

print(f"  {x}")
print(f"x {y}")
print("-" * max(len(str(x)), len(str(y)), len(str(result))))
print(f"  {result}")
