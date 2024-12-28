def umn(x, y):
    # Базовый случай: если числа однозначные, возвращаем их произведение
    if x < 10 or y < 10:
        return x * y

    # Определение максимальной длины числа и половины для разбиения
    n = max(len(str(x)), len(str(y)))
    half_n = n // 2

    # Разбиваем числа на части
    x1, x0 = divmod(x, 10**half_n)
    y1, y0 = divmod(y, 10**half_n)

    # Рекурсивные вызовы для частей
    z0 = umn(x0, y0)
    z1 = umn(x1, y1)
    z2 = umn(x1 + x0, y1 + y0) - z1 - z0

    # Собираем результат
    return z1 * 10**(2 * half_n) + z2 * 10**half_n + z0

x = 25
y = 5
result = umn(x, y)
print(f"Результат умножения {x} и {y}: {result}")
