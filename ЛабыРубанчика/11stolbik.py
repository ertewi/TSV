def multiply_strings(num1, num2):
    if num1 == "0" or num2 == "0":
        return "0"
    
    len1, len2 = len(num1), len(num2)
    result = [0] * (len1 + len2)

    # Вспомогательная функция для умножения цифр
    def multiply_digits(i, j):
        return (ord(num1[i]) - ord('0')) * (ord(num2[j]) - ord('0'))

    # Вспомогательная функция для добавления промежуточных результатов
    def add_to_result(pos1, pos2, product):
        total = product + result[pos2]
        result[pos2] = total % 10
        result[pos1] += total // 10

    # Основной цикл умножения
    for i in range(len1 - 1, -1, -1):
        for j in range(len2 - 1, -1, -1):
            mul = multiply_digits(i, j)
            add_to_result(i + j, i + j + 1, mul)

    # Преобразование массива результата в строку
    result_str = ''.join(map(str, result)).lstrip('0')
    return result_str if result_str else "0"


num1 = "123"
num2 = "456"
result = multiply_strings(num1, num2)

print(f" {num1}")
print(f"x {num2}")
print("--")
print(f"{result}")
