from typing import List

def longestCommonPrefix(strs: List[str]) -> str:
    substr = strs[0]
    for s in strs:
        b = False
        i = 0
        for e, ee in zip(substr, s):
            i += 1
            if e != ee:
                substr = substr[:i-1]
                b = True
        if not b and len(s) < len(substr):
            substr = s
    return substr


print(longestCommonPrefix(['amir', 'a', 'abbas']))