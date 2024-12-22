from enigma import Enigma
from config import EnigmaConfig

message = "AABDCDHGKFKFHTJKFKJGKFJGKFKzgjkfJKJKJKGJKFJKGJKFGJKF"

encodedMessage = Enigma(EnigmaConfig()).encodeMessage(message)
decodedMessage = Enigma(EnigmaConfig()).encodeMessage(encodedMessage)

print(encodedMessage, decodedMessage)