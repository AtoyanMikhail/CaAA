def compute_prefix(pattern):
    n = len(pattern)
    pi = [0] * n
    j = 0
    for i in range(1, n):
        while j > 0 and pattern[i] != pattern[j]:
            j = pi[j - 1]
        if pattern[i] == pattern[j]:
            j += 1
            pi[i] = j
        else:
            pi[i] = 0
    return pi

def find_cyclic_shift(a, b):
    if len(a) != len(b):
        return -1
    n = len(a)
    if n == 0:
        return 0
    if a == b:
        return 0
    
    pi = compute_prefix(a)
    j = 0
    for i in range(2 * n):
        current_char = b[i % n]
        while j > 0 and current_char != a[j]:
            j = pi[j - 1]
        if current_char == a[j]:
            j += 1
        if j == n:
            start_index = i - n + 1
            if 0 <= start_index < n:
                return (n - start_index) % n
    return -1