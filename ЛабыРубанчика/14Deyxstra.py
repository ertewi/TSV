import heapq

def deyxtra(graph, start):
    D = {node: float('inf') for node in graph}
    D[start] = 0
    priority_queue = [(0, start)]

    while priority_queue:
        current_distance, w = heapq.heappop(priority_queue)

        if current_distance > D[w]:
            continue

        for v, cost in graph[w].items():
            new_distance = current_distance + cost
            if new_distance < D[v]:
                D[v] = new_distance
                heapq.heappush(priority_queue, (new_distance, v))

    return D

graph = {
    0: {1: 4, 2: 1},
    1: {3: 1},
    2: {1: 2, 3: 5},
    3: {}
}

distances = deyxtra(graph, 0)
print(distances)
