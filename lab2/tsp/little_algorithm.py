from math import inf
import heapq

class LittleTSP:
    def __init__(self, matrix, start_city):
        self.matrix = matrix
        self.n = len(matrix)
        self.start = start_city
        self.best_cost = inf
        self.best_path = []

    class Node:
        def __init__(self, path, visited, cost):
            self.path = path          # Текущий путь (список городов)
            self.visited = visited    # Множество посещённых городов
            self.cost = cost          # Текущая стоимость пути

        def __lt__(self, other):
            return self.cost < other.cost  # Для работы с heapq

    def solve(self):
        root = self.Node(
            path=[self.start],
            visited={self.start},
            cost=0
        )
        
        # Приоритетная очередь для хранения узлов
        heap = []
        heapq.heappush(heap, root)

        while heap:
            current_node = heapq.heappop(heap)

            if len(current_node.path) == self.n:
                final_cost = current_node.cost + self.matrix[current_node.path[-1]][self.start]
                if final_cost < self.best_cost:
                    self.best_cost = final_cost
                    self.best_path = current_node.path + [self.start]
                continue

            if current_node.cost >= self.best_cost:
                continue

            last_city = current_node.path[-1]
            for next_city in range(self.n):
                if next_city not in current_node.visited and self.matrix[last_city][next_city] != inf:
                    new_path = current_node.path.copy()
                    new_path.append(next_city)
                    
                    new_visited = current_node.visited.copy()
                    new_visited.add(next_city)
                    
                    new_cost = current_node.cost + self.matrix[last_city][next_city]
                    
                    child = self.Node(new_path, new_visited, new_cost)
                    heapq.heappush(heap, child)

        return list(x + 1 for x in self.best_path), self.best_cost

