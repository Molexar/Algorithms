from typing import List
from itertools import chain
import random
 
def initRotorConfig() -> List[int]:
    alphabet = list(range(26))
    random.shuffle(alphabet)
    return alphabet
 
class EnigmaConfig:
    rotorIII: int = ord('Q') % 65
    rotorII: int = ord('U') % 65
    rotorI: int = ord('C') % 65
    alphabet: List[int] = [
        chr(i) for i in chain(range(65, 91)) # A-Z
    ]
    rotorIIIConfig: List[int] = [
        4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9
    ]
    rotorIIConfig: List[int] = [
        0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4
    ]
    rotorIConfig: List[int] = [
        1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14
    ]
    reflectorConfig: List[int] = [
        24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19
    ]