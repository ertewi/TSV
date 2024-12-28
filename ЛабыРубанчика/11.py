def stolbik(num1, num2):
    # Если одно из чисел равно "0", результат всегда "0"
    if num1 == "0" or num2 == "0":
        return "0"

    n1, n2 = len(num1), len(num2)
    result = [0] * (n1 + n2)

    # Цикл для умножения каждой цифры
    for i in range(n1 - 1, -1, -1):
        for j in range(n2 - 1, -1, -1):
            # Умножение цифр
            mul = (ord(num1[i]) - ord('0')) * (ord(num2[j]) - ord('0'))
            # Индексы для сложения с существующим результатом
            p1, p2 = i + j, i + j + 1
            # Суммируем результат умножения с текущей позицией
            total = mul + result[p2]

            # Обновляем текущую и следующую позицию
            result[p2] = total % 10
            result[p1] += total // 10

    # Преобразуем список результата в строку, убирая ведущие нули
    result_str = ''.join(map(str, result)).lstrip('0')
    
    # Если результат пустой (в случае, если все нули), возвращаем "0"
    return result_str if result_str else "0"

num1 = "5"
num2 = "5"
print(stolbik(num1, num2))
