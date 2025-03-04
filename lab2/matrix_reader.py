from math import inf

class MatrixReader:
    def __init__(self, file_path):
        self.file_path = file_path
        self.matrix = []

    def read_matrix(self):
        try:
            with open(self.file_path, 'r') as file:
                for line in file:
                    row = list(int(x) if x != "inf" else inf for x in line.strip().split())
                    self.matrix.append(row)
            return self.matrix
        except FileNotFoundError:
            print(f"Файл {self.file_path} не найден.")
            return None
        except ValueError:
            print("Ошибка: файл содержит некорректные данные (ожидались числа).")
            return None

    def print_matrix(self):
        if not self.matrix:
            print("Матрица пуста.")
            return
        for row in self.matrix:
            print(" ".join(map(str, row)))


