def search(needle, haystack):
    if not needle:
        return []
    
    occurrences = []
    
    for i in range(len(haystack) - len(needle) + 1):
        if haystack[i:i+len(needle)] == needle:
            occurrences.append(i)
                
    return occurrences