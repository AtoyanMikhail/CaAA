from matrix_reader import MatrixReader
from tsp import LittleTSP, ApproximateTSP
from math import inf

reader = MatrixReader("matrix.txt")
matrix = reader.read_matrix()

start_city = 0

solver = ApproximateTSP(matrix, start_city)
path, cost = solver.solve()

print(f"Приближённый путь: {path}")
print(f"Стоимость: {cost}")

if matrix:
    start_vertex = 0

    tsp = LittleTSP(matrix, start_vertex)
    path, cost = tsp.solve()
    
    print(f"Оптимальный путь: {path}")
    print(f"Стоимость: {cost}")