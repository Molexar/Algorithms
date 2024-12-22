package main

import (
	"fmt"

	"molexar.org/crypto/internal"
)

func main() {

	bitSize := 10
	p, _ := internal.GenerateRandomPrime(bitSize)
	g := internal.FindPrimitiveRoot(p)

	// Генерация случайных приватных ключей для Алисы и Боба
	privateKeyA := internal.GenerateRandomNumber(p)
	privateKeyB := internal.GenerateRandomNumber(p)

	// Вычисление публичных ключей
	publicKeyA := internal.ModExp(g, privateKeyA, p)
	publicKeyB := internal.ModExp(g, privateKeyB, p)

	// Обмен публичными ключами
	sharedSecretA := internal.ModExp(publicKeyB, privateKeyA, p)
	sharedSecretB := internal.ModExp(publicKeyA, privateKeyB, p)

	// Проверка, что общий секрет совпадает
	if sharedSecretA.Cmp(sharedSecretB) == 0 {
		fmt.Println("Общий секрет совпадает:", sharedSecretA)
	} else {
		fmt.Println("Ошибка: общий секрет не совпадает")
	}
}
