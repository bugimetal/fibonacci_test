package fibonacci

import (
	"math/big"
	"sync"
)

type matrix [][]*big.Int

type matrixCache struct {
	values map[int]*matrix
	mutex  *sync.Mutex
}

func CreateMatrixCache() *matrixCache {
	cache := make(map[int]*matrix)
	return &matrixCache{values: cache, mutex: &sync.Mutex{}}
}

func (c *matrixCache) Get(number int) (*matrix, error) {
	c.mutex.Lock()
	value := c.values[number]
	c.mutex.Unlock()
	if value == nil {
		return value, ErrNotFound
	} else {
		return value, nil
	}
}

func (c *matrixCache) Add(number int, value *matrix) {
	c.mutex.Lock()
	c.values[number] = value
	c.mutex.Unlock()
}

func Qmatrix(n int, cache *matrixCache) *big.Int {
	return qmatrixInternal(n, cache)
}

func qmatrixInternal(n int, cache *matrixCache) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	var q matrix
	q = append(q, []*big.Int{big.NewInt(1), big.NewInt(1)})
	q = append(q, []*big.Int{big.NewInt(1), big.NewInt(0)})

	var matricies []*matrix
	i := 0
	for n > 0 {
		exp := n & 0x1
		n = n >> 1
		pow := exp << uint(i)
		i += 1
		if exp != 0 {
			matricies = append(matricies, matrixPowerOfTwo(q, pow, cache))
		}
	}

	for len(matricies) > 1 {
		m1 := matricies[len(matricies)-1]
		matricies = matricies[:len(matricies)-1]
		m2 := matricies[len(matricies)-1]
		matricies = matricies[:len(matricies)-1]
		matricies = append(matricies, matrixMultiply(m1, m2))
	}
	return (*matricies[0])[0][1]
}

func matrixMultiply(a, b *matrix) *matrix {
	c11 := new(big.Int).Add(new(big.Int).Mul((*a)[0][0], (*b)[0][0]), new(big.Int).Mul((*a)[0][1], (*b)[1][0]))
	c12 := new(big.Int).Add(new(big.Int).Mul((*a)[0][0], (*b)[0][1]), new(big.Int).Mul((*a)[0][1], (*b)[1][1]))
	c21 := new(big.Int).Add(new(big.Int).Mul((*a)[1][0], (*b)[0][0]), new(big.Int).Mul((*a)[1][1], (*b)[1][0]))
	c22 := new(big.Int).Add(new(big.Int).Mul((*a)[1][0], (*b)[0][1]), new(big.Int).Mul((*a)[1][1], (*b)[1][1]))
	var c matrix
	c = append(c, []*big.Int{c11, c12})
	c = append(c, []*big.Int{c21, c22})
	return &c
}

func matrixPowerOfTwo(m matrix, p int, cache *matrixCache) *matrix {
	if p == 1 {
		return &m
	}
	cachedMatrix, err := cache.Get(p)

	if err == nil {
		return cachedMatrix
	}

	k := matrixPowerOfTwo(m, p/2, cache)
	res := matrixMultiply(k, k)

	cache.Add(p, res)

	return res
}
