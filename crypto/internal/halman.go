package internal

import (
	"crypto/rand"
	"math/big"
)

// Функция для генерации случайного числа в заданном диапазоне
func GenerateRandomNumber(max *big.Int) *big.Int {
	num, _ := rand.Int(rand.Reader, max)
	return num
}

// Функция для вычисления (base^exp) % mod
func ModExp(base, exp, mod *big.Int) *big.Int {
	result := new(big.Int).Exp(base, exp, mod)
	return result
}

// Функция для генерации случайного простого числа
func GenerateRandomPrime(bitSize int) (*big.Int, error) {
	prime, err := rand.Prime(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}
	return prime, nil
}

// Функция для нахождения примитивного корня по модулю p
func FindPrimitiveRoot(p *big.Int) *big.Int {
	var fact []*big.Int
	phi := new(big.Int).Sub(p, big.NewInt(1))
	n := new(big.Int).Set(phi)

	// Факторизация phi
	for i := big.NewInt(2); i.Mul(i, i).Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		if n.Mod(n, i).Cmp(big.NewInt(0)) == 0 {
			fact = append(fact, new(big.Int).Set(i))
			for n.Mod(n, i).Cmp(big.NewInt(0)) == 0 {
				n.Div(n, i)
			}
		}
	}
	if n.Cmp(big.NewInt(1)) > 0 {
		fact = append(fact, new(big.Int).Set(n))
	}

	// Проверка каждого возможного примитивного корня
	for res := big.NewInt(2); res.Cmp(p) < 0; res.Add(res, big.NewInt(1)) {
		ok := true
		for _, factor := range fact {
			if ModExp(res, new(big.Int).Div(phi, factor), p).Cmp(big.NewInt(1)) == 0 {
				ok = false
				break
			}
		}
		if ok {
			return res
		}
	}
	return big.NewInt(-1)
}
