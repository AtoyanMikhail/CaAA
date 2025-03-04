import timeit
from math import inf
from mst_dfs_algorithm import ApproximateTSP

def generate_symmetric_matrix(size):
    """Генерация симметричной матрицы с возрастающими весами"""
    matrix = [[inf] * size for _ in range(size)]
    for i in range(size):
        for j in range(i+1, size):
            matrix[i][j] = matrix[j][i] = (i + j) * 10
    return matrix

def benchmark_case(matrix, start=0, iterations=5):
    """Замер времени для конкретной матрицы"""
    stmt = f'ApproximateTSP({matrix}, {start}).solve()'
    setup = 'from __main__ import ApproximateTSP; from math import inf'
    times = timeit.repeat(stmt, setup, number=1, repeat=iterations)
    return min(times), sum(times)/iterations, max(times)

def print_results(size, results):
    """Форматированный вывод результатов"""
    print(f"\n{'Размер матрицы':^17}|{'Минимум':^14}|{'Среднее':^14}|{'Максимум':^14}")
    print(f"{'-'*17}|{'-'*14}|{'-'*14}|{'-'*14}")
    print(f"{size:^17}|{results[0]:^12.8f}s |{results[1]:^12.8f}s |{results[2]:^12.8f}s")

def run_benchmarks():
    """Набор тестовых матриц разных размеров"""
    test_cases = {
        10: generate_symmetric_matrix(10),
        20: generate_symmetric_matrix(20),
        50: generate_symmetric_matrix(50),
        100: generate_symmetric_matrix(100),
        200: generate_symmetric_matrix(200)
    }

    sizes = []
    avg_times = []
    
    for size, matrix in test_cases.items():
        results = benchmark_case(matrix, iterations=5)
        print_results(size, results)
        sizes.append(size)
        avg_times.append(results[1])

if __name__ == '__main__':
    print("🚀 Запуск бенчмарка приближённого алгоритма TSP")
    run_benchmarks()
    print("\n✅ Все тесты завершены.")