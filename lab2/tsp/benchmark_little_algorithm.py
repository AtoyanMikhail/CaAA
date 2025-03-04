import timeit
from math import inf
from little_algorithm import LittleTSP

def benchmark_case(matrix, start=0, iterations=5):
    """Замер времени для конкретной матрицы"""
    stmt = f'LittleTSP({matrix}, {start}).solve()'
    setup = f'from __main__ import LittleTSP; from math import inf'
    times = timeit.repeat(stmt, setup, number=1, repeat=iterations)
    return min(times), sum(times)/iterations, max(times)

def print_results(size, results):
    """Форматированный вывод результатов"""
    print(f"\n{'Размер матрицы':^16}|{'Минимум':^13}|{'Среднее':^13}|{'Максимум':^13}")
    print(f"{'-'*16}|{'-'*13}|{'-'*13}|{'-'*13}")
    print(f" {size:^15}| {results[0]:^10.8f}s | {results[1]:^10.8f}s | {results[2]:^10.8f}s")

def run_benchmarks():
    """Набор тестовых матриц разных размеров"""
    test_cases = {
        # Матрица 3x3 (оптимальный путь 1-2-3-1 = 80)
        3: [
            [inf, 10, 15],
            [10, inf, 35],
            [15, 35, inf]
        ]
        ,
        # Матрица 4x4 (стандартный тест)
        4: [
            [inf, 10, 15, 20],
            [10, inf, 35, 25],
            [15, 35, inf, 30],
            [20, 25, 30, inf]
        ],
        # Матрица 5x5 с шаблоном
        5: [
            [inf, 20, 30, 10, 15],
            [20, inf, 25, 35, 40],
            [30, 25, inf, 45, 50],
            [10, 35, 45, inf, 60],
            [15, 40, 50, 60, inf]
        ],
        # Матрица 6x6 с возрастающими значениями
        6: [
            [inf, 12, 18, 23, 27, 31],
            [12, inf, 14, 19, 25, 29],
            [18, 14, inf, 16, 22, 26],
            [23, 19, 16, inf, 17, 24],
            [27, 25, 22, 17, inf, 21],
            [31, 29, 26, 24, 21, inf]
        ],
        # Матрица 7x7 (полностью связанная)
        7: [
            [inf, 10, 15, 20, 25, 30, 35],
            [10, inf, 12, 18, 22, 28, 32],
            [15, 12, inf, 14, 16, 24, 30],
            [20, 18, 14, inf, 15, 20, 25],
            [25, 22, 16, 15, inf, 18, 22],
            [30, 28, 24, 20, 18, inf, 12],
            [35, 32, 30, 25, 22, 12, inf]
        ]
    }

    for size, matrix in test_cases.items():
        results = benchmark_case(matrix, iterations=3)
        print_results(size, results)

if __name__ == '__main__':
    print("🚀 Запуск комплексного бенчмарка алгоритма Литтла")
    run_benchmarks()
    print("\n✅ Все тесты завершены")