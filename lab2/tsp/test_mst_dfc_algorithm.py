import unittest
from math import inf
from mst_dfs_algorithm import ApproximateTSP

class TestApproximateTSP(unittest.TestCase):
    def setUp(self):
        self.sample_matrix = [
            [inf, 10, 15, 20],
            [10, inf, 25, 5],
            [15, 25, inf, 30],
            [20, 5, 30, inf]
        ]
        
        self.single_city = [[0]]
        
        self.disconnected_matrix = [
            [inf, inf, inf, inf],
            [inf, inf, 15, 25],
            [inf, 15, inf, inf],
            [inf, 25, inf, inf]
        ]
        
    def calculate_mst_cost(self, matrix, mst_edges):
        """Вычисляет стоимость MST по списку рёбер"""
        return sum(matrix[u][v] for u, v in mst_edges)

    def test_basic_functionality(self):
        """Проверка 2-приближения"""
        tsp = ApproximateTSP(self.sample_matrix, 0)
        path, cost = tsp.solve()
        mst = tsp.find_mst()
        
        mst_cost = self.calculate_mst_cost(self.sample_matrix, mst)
        upper_bound = 2 * mst_cost
        
        self.assertEqual(path[0], 0)
        self.assertEqual(path[-1], 0)
        self.assertEqual(len(set(path)), 4)
        self.assertLessEqual(cost, upper_bound)  # Основной критерий
        self.assertTrue(60 <= upper_bound <= 100)  # Доп. проверка для примера

    def test_single_city(self):
        """Один город - нулевая стоимость"""
        tsp = ApproximateTSP(self.single_city, 0)
        path, cost = tsp.solve()
        self.assertEqual(path, [0,0])
        self.assertEqual(cost, 0)

    def test_disconnected_graph(self):
        """Несвязный граф - бесконечная стоимость"""
        tsp = ApproximateTSP(self.disconnected_matrix, 0)
        path, cost = tsp.solve()
        self.assertEqual(cost, inf)
        self.assertEqual(path, [0,0])

    def test_approximation_ratio(self):
    """Проверка коэффициента аппроксимации для разных случаев"""
    test_cases = [
        (
            [
                [inf, 10, 15, 20],
                [10, inf, 35, 25],
                [15, 35, inf, 30],
                [20, 25, 30, inf]
            ],
            90  # Ожидаемая верхняя граница 2*MST
        ),
        (
            [
                [inf, 12, 18, 23],
                [12, inf, 14, 19],
                [18, 14, inf, 16],
                [23, 19, 16, inf]
            ],
            70  # Ожидаемая верхняя граница 2*MST
        )
    ]
    
    for matrix, expected_upper in test_cases:
        tsp = ApproximateTSP(matrix, 0)
        path, cost = tsp.solve()
        mst = tsp.find_mst()
        
        mst_cost = self.calculate_mst_cost(matrix, mst)
        upper_bound = 2 * mst_cost
        
        with self.subTest(matrix=matrix):
            self.assertLessEqual(cost, upper_bound, 
                f"Стоимость {cost} превышает 2*MST ({upper_bound})")
            self.assertLessEqual(upper_bound, expected_upper, 
                f"Верхняя граница {upper_bound} больше ожидаемой {expected_upper}")

if __name__ == '__main__':
    unittest.main()