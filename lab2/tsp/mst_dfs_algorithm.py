from math import inf
import heapq

class ApproximateTSP:
    def __init__(self, matrix, start_city):
        self.matrix = matrix
        self.n = len(matrix)
        self.start = start_city

    def find_mst(self):
        # Алгоритм Прима для построения MST
        mst = []  # Список рёбер MST
        visited = set()
        heap = []  # Куча для выбора рёбер

        # Начинаем с начальной вершины
        visited.add(self.start)
        for next_city in range(self.n):
            if next_city != self.start and self.matrix[self.start][next_city] != inf:
                heapq.heappush(heap, (self.matrix[self.start][next_city], self.start, next_city))

        while heap:
            cost, u, v = heapq.heappop(heap)
            if v not in visited:
                visited.add(v)
                mst.append((u, v))
                for next_city in range(self.n):
                    if next_city not in visited and self.matrix[v][next_city] != inf:
                        heapq.heappush(heap, (self.matrix[v][next_city], v, next_city))
        return mst

    def dfs_traversal(self, mst):
        # Обход MST в глубину (DFS)
        adjacency_list = {i: [] for i in range(self.n)}
        for u, v in mst:
            adjacency_list[u].append(v)
            adjacency_list[v].append(u)

        path = []
        stack = [self.start]
        visited = set()

        while stack:
            node = stack.pop()
            if node not in visited:
                visited.add(node)
                path.append(node)
                for neighbor in reversed(adjacency_list[node]):
                    if neighbor not in visited:
                        stack.append(neighbor)
        return path

    def solve(self):
        # Шаг 1: Построение MST
        mst = self.find_mst()

        # Шаг 2: Обход MST
        path = self.dfs_traversal(mst)

        # Шаг 3: Добавление возврата в стартовую вершину
        path.append(self.start)

        # Шаг 4: Вычисление стоимости
        cost = 0
        for i in range(len(path) - 1):
            cost += self.matrix[path[i]][path[i + 1]]

        return path, cost
