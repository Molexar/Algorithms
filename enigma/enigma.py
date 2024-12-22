from typing import List
from config import EnigmaConfig

class Rotor:
    def __init__(self, config: List[int], state: int = 0):
        self.config = config
        self.state = state
    
    def rotate(self) -> int:
        self.config = self.config[len(self.config)-1:] + self.config[:len(self.config)-1]
        self.state += 1
        if self.state == 26:
            self.state = 0
            return 1
        return 0

    
    def forward(self, char: int) -> int:
        return self.config[char]
    
    def backward(self, char: int) -> int:
        return self.config.index(char)

class Reflector:
    def __init__(self, config: List[int]):
        self.config = config
    
    def reflect(self, char: int) -> int:
        return self.config.index(char)
 
class Enigma:
    def __init__(self, config: EnigmaConfig):
        self.config = config
        self.rotorI = Rotor(config.rotorIConfig)
        self.rotorII = Rotor(config.rotorIConfig)
        self.rotorIII = Rotor(config.rotorIConfig)
        self.reflector = Reflector(config.reflectorConfig)

    def forward(self, char: int) -> int:
        ext = self.rotorI.rotate()
        if ext:
            ext = self.rotorII.rotate()
            if ext:
                self.rotorIII.rotate()


        rotorI = self.rotorI.forward(char)
        rotorII = self.rotorII.forward(rotorI)
        rotorIII = self.rotorIII.forward(rotorII)
        reflected = self.reflector.reflect(rotorIII)

        return reflected
    
    def backward(self, char: int) -> int:
        rotorIII = self.rotorIII.backward(char)
        rotorII = self.rotorII.backward(rotorIII)
        rotorI = self.rotorI.backward(rotorII)

        return rotorI
    
 
    def encodeSymbol(self, char: int) -> int:
        forward = self.forward(char)

        backward = self.backward(forward)

        return backward
 
    
    def encodeMessage(self, message: str) -> str:
        encoded = ''
 
        for char in message:
            if char in self.config.alphabet:
                encoded += self.config.alphabet[self.encodeSymbol(self.config.alphabet.index(char))]
            elif char == ' ':
                encoded += char
        return encoded