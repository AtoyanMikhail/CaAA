import time 
import os

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

def search(needle, haystack):
    if not needle:
        return []

    pi = compute_prefix(needle)
    
    occurrences = []
    j = 0
    for i in range(len(haystack)): 
        if j == 0:      #Код для визуализации
            start = i   #
        
        while j > 0 and haystack[i] != needle[j]:
            j = pi[j-1]
        if haystack[i] == needle[j]:
            j += 1
            if j == len(needle):
                occurrences.append(i - j + 1)
                j = pi[j-1]
        
        
        os.system("clear")                                                          #Код для визуализации
        print(haystack)                                                             #
        print(" "*(i-j+1) + (haystack[start:start + j] if j != 0 else "_"))         #
        for o in reversed(occurrences):                                             #
            print(" "*(o) + needle + "\n" + " "*(o) + "^" + "\n" + " "*(o) + str(o))#
        time.sleep(1)                                                               #
        
    return occurrences

s1 = input()
s2 = input()
search(s1, s2)