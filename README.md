# Лабораторные работы по дисциплине "Проектирование и анализ алгоритмов"

Добро пожаловать в репозиторий с лабораторными работами по курсу **"Проектирование и анализ алгоритмов"**! Здесь собраны реализации различных алгоритмов, их анализ и оптимизации, выполненные в рамках учебной программы.

---

## 📁 Структура репозитория

Репозиторий организован следующим образом:
```
.
├── lab1/ # Лабораторная работа 1
│ ├── src/ # Исходный код
│ ├── Atoyan_Mikhail_lb2.docx # Отчет по лабораторной работе в формате docx
│ └── Atoyan_Mikhail_lb1.pdf # Отчет по лабораторной работе в формате pdf
├── ...
└── README.md # Этот файл
```


Каждая лабораторная работа находится в отдельной папке, содержащей:
- Исходный код на языке программирования Go.
- Отчет в формате PDF с анализом алгоритмов, их временной и пространственной сложности.

---

## 📚 Список лабораторных работ

1. **Лабораторная работа 1: Разбиение квадрата на меньшие квадраты**
   - Задача: Покрыть квадратную таблицу размера \( n \times n \) минимальным количеством квадратов.
   - Алгоритмы: Backtracking с оптимизациями.
   - Оптимизации:
     - Для четных \( n \).
     - Для простых \( n \).
     - Для чисел вида \( n = 2^r - 1 \).
     - Для составных \( n \).

---

## 🛠️ Технологии и инструменты

- **Языки программирования:** Go.
- **Инструменты анализа:**
  - Профилирование времени выполнения.
  - Оценка временной и пространственной сложности.

---

## 📊 Анализ алгоритмов

В каждой лабораторной работе проводится подробный анализ алгоритмов:
- **Временная сложность:** Оценка времени выполнения в лучшем, среднем и худшем случаях.
- **Пространственная сложность:** Оценка использования памяти.
- **Оптимизации:** Описание примененных оптимизаций и их влияние на производительность.

---

## 📝 Как использовать

1. Перейдите в папку с интересующей вас лабораторной работой:
    ```bash
    cd lab1/
    ```
2. Ознакомьтесь с описанием задачи в файле Atoyan_Mikhail_lb*.pdf.

3. Ознакомьтесь с отчетом в файле Atoyan_Mikhail_lb*.pdf для подробного анализа.

4. Изучите исходный код в папке src/.

5. Запустите программу:
   ```bash
   go run . #Из папки с исходным кодом
   ```

Автор: Атоян Михаил Артурович

Группа: 3343

Год: 2025
