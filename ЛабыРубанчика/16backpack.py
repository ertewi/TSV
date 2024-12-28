import numpy as np

def rukzak(M, price: dict):
    # n - количество предметов
    n = len(price["m"])
    
    # dp - массив для хранения максимальных стоимостей для каждой массы от 0 до M
    dp = np.zeros(M + 1)

    # Проходим по всем предметам
    for i in range(n):
        # Обновляем массив dp в обратном порядке, чтобы избежать переиспользования обновленных значений
        for j in range(M, price["m"][i] - 1, -1):
            dp[j] = max(dp[j], dp[j - price["m"][i]] + price["c"][i])

    return dp[M]

price = {
    "m" : [3, 5, 8],
    "c" : [8, 14, 23]
}

M = 15
print(f"Максимальная стоимость набора товаров при грузоподъемности {M}: {rukzak(M, price)}")
