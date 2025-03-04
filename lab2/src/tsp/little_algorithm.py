from math import inf
import heapq
import copy

class LittleTSP:
    def __init__(self, matrix, start_city):
        self.original_matrix = matrix
        self.n = len(matrix)
        self.start = start_city
        self.best_cost = inf
        self.best_path = []

    class Node:
        def __init__(self, path, matrix, cost):
            self.path = path          # Текущий путь
            self.matrix = matrix      # Текущая редуцированная матрица
            self.cost = cost          # Нижняя граница стоимости

        def __lt__(self, other):
            return self.cost < other.cost

    def reduce_matrix(self, matrix):
        reduction_cost = 0

        for i in range(len(matrix)):
            min_val = min(matrix[i])
            if min_val != inf and min_val > 0:
                reduction_cost += min_val
                matrix[i] = [x - min_val for x in matrix[i]]
        
        for j in range(len(matrix[0])):
            col = [matrix[i][j] for i in range(len(matrix))]
            min_val = min(col)
            if min_val != inf and min_val > 0:
                reduction_cost += min_val
                for i in range(len(matrix)):
                    matrix[i][j] -= min_val
        
        return matrix, reduction_cost

    def solve(self):
        initial_matrix = copy.deepcopy(self.original_matrix)
        reduced_matrix, reduction_cost = self.reduce_matrix(initial_matrix)
        
        root = self.Node(
            path=[self.start],
            matrix=reduced_matrix,
            cost=reduction_cost
        )
        
        heap = []
        heapq.heappush(heap, root)

        while heap:
            current_node = heapq.heappop(heap)

            if len(current_node.path) == self.n:
                final_cost = current_node.cost + current_node.matrix[current_node.path[-1]][self.start]
                if final_cost < self.best_cost:
                    self.best_cost = final_cost
                    self.best_path = current_node.path + [self.start]
                continue
            
            if current_node.cost >= self.best_cost:
                continue

            last_city = current_node.path[-1]
            for next_city in range(self.n):
                if current_node.matrix[last_city][next_city] != inf:
                    # Создаём новую матрицу для дочернего узла
                    new_matrix = copy.deepcopy(current_node.matrix)
                    
                    # Удаляем строку last_city и столбец next_city
                    for i in range(self.n):
                        new_matrix[i][next_city] = inf
                    new_matrix[last_city] = [inf] * self.n
                    
                    # Редуцируем новую матрицу
                    reduced_new_matrix, reduction = self.reduce_matrix(new_matrix)
                    new_cost = current_node.cost + current_node.matrix[last_city][next_city] + reduction
                    
                    # Создаём новый узел
                    new_path = current_node.path.copy()
                    new_path.append(next_city)
                    
                    child = self.Node(
                        path=new_path,
                        matrix=reduced_new_matrix,
                        cost=new_cost
                    )
                    
                    heapq.heappush(heap, child)

        return list(x + 1 for x in self.best_path), self.best_cost