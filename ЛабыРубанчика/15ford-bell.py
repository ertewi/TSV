class Graph:
    def __init__(self, num_vertices):
        # Конструктор: инициализация графа с количеством вершин
        self.num_vertices = num_vertices  # количество вершин
        self.edges = []  # список рёбер в формате (начало, конец, вес)

    def add_edge(self, start, end, weight):
        # Добавляем ребро между вершинами start и end с весом weight
        self.edges.append((start, end, weight))

    def bellman_ford(self, source):
        # Инициализация расстояний: все бесконечны, кроме исходной вершины
        distances = [float('inf')] * self.num_vertices
        distances[source] = 0

        # Расслабляем все рёбра (V-1) раз, где V — количество вершин
        for _ in range(self.num_vertices - 1):
            for start, end, weight in self.edges:
                # Если кратчайший путь до вершины start известен
                # и можно улучшить расстояние до вершины end, то обновляем его
                if distances[start] != float('inf') and distances[start] + weight < distances[end]:
                    distances[end] = distances[start] + weight

        # Проверка на наличие циклов с отрицательным весом
        for start, end, weight in self.edges:
            if distances[start] != float('inf') and distances[start] + weight < distances[end]:
                return "Граф содержит цикл с отрицательным весом"

        # Возвращаем итоговые кратчайшие расстояния
        return distances

    def print_distances(self, source, distances):
        # Выводим кратчайшие расстояния от исходной вершины
        print(f"Кратчайшие расстояния от вершины {source}:")
        for vertex, distance in enumerate(distances):
            if distance == float('inf'):
                print(f"Вершина {vertex}: недостижима")
            else:
                print(f"Вершина {vertex}: {distance}")


# Пример использования
g = Graph(4)
g.add_edge(0, 1, 1)
g.add_edge(0, 3, 4)
g.add_edge(1, 2, 2)
g.add_edge(3, 2, 3)
g.add_edge(2, 3, -1)

# Запуск алгоритма Беллмана-Форда из вершины 0
result = g.bellman_ford(0)

# Проверка результата и вывод
if isinstance(result, str):
    print(result)  # Если был найден цикл с отрицательным весом
else:
    g.print_distances(0, result)  # Если успешный расчёт, выводим расстояния
