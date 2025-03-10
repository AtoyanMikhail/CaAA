import random
import time
import sys
from kmp import search as kmp
from naive_search import search as naive

def generate_test_data(haystack_len, needle_len):
    base = ['x'] * haystack_len
    needle = ''.join(random.choices('abcdefghijklmno', k=needle_len))
    
    occurrences = haystack_len//len(needle)//2
    step = haystack_len // (occurrences + 1)
    for i in range(occurrences):
        pos = i * step + random.randint(0, step - needle_len)
        base[pos:pos+needle_len] = list(needle)
        
    haystack = ''.join(base)
    return haystack, needle, occurrences

def benchmark(f):
    random.seed(42)
    tests = [
        (180000, 20000),
        (360000, 40000),
        (540000, 60000),
        (720000, 80000),
        (900000, 100000),
        (1080000, 120000),
    ]
    
    print(f"{'Haystack len':<12} | {'Needle len':<10} | {'Time (ms)':<10} | {'Occurrences':<12} | {'Mem (MB)':<10}")
    print("-" * 70)

    for haystack_len, needle_len in tests:
        haystack, needle, expected = generate_test_data(haystack_len, needle_len)
        
        start = time.perf_counter()
        result = f(needle, haystack)
        duration = (time.perf_counter() - start) * 1000
        
        assert len(result) == expected, \
            f"Ошибка: найдено {len(result)}, ожидалось {expected}"
        
        mem_usage = (sys.getsizeof(haystack) + sys.getsizeof(needle)) / (1024**2)
        
        print(
            f"{haystack_len:<12} | {needle_len:<10} | "
            f"{duration:<10.2f} | {len(result):<12} | "
            f"{mem_usage:.2f} MB"
        )

if __name__ == "__main__":
    print("Алгоритм Кнута-Морриса-Прутта:\n")
    benchmark(kmp)
    print("\nНаивный перебор:\n")
    benchmark(naive)