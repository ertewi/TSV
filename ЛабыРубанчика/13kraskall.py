def find(parent, u):
    if parent[u] != u:
        parent[u] = find(parent, parent[u])  # сжатие пути
    return parent[u]

def union(parent, rank, u, v):
    root_u, root_v = find(parent, u), find(parent, v)
    if root_u != root_v:
        if rank[root_u] > rank[root_v]:
            parent[root_v] = root_u
        else:
            parent[root_u] = root_v
            if rank[root_u] == rank[root_v]:
                rank[root_v] += 1

def kraskall(n, edges):
    edges.sort(key=lambda x: x[2])  # сортировка по весу
    parent = list(range(n))  # родительский элемент для каждого узла
    rank = [0] * n  # ранг (высота дерева) для каждого узла
    mst = []

    for u, v, weight in edges:
        if find(parent, u) != find(parent, v):  # проверка на циклы
            union(parent, rank, u, v)
            mst.append((u, v, weight))
    
    return mst

edges = [(0, 1, 1), (0, 2, 2), (1, 2, 1)]
n = 3
print("MST:", kraskall(n, edges))

edges2 = [(0, 1, 4), (0, 2, 1), (1, 2, 2), (1, 3, 5), (2, 3, 3)]
n2 = 4
print("MST 2:", kraskall(n2, edges2))
