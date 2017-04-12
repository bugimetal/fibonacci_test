package fibonacci

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("Index not found")
)

func NaiveImplementations(number int) int {
	if number == 0 || number == 1 {
		return number
	}

	return NaiveImplementations(number-2) + NaiveImplementations(number-1)
}

type computedValues struct {
	values map[int]int
	mutex  *sync.Mutex
}

func CreateCache() *computedValues {
	values := make(map[int]int)
	return &computedValues{values: values, mutex: &sync.Mutex{}}
}

func (c *computedValues) Get(number int) (int, error) {
	c.mutex.Lock()
	value := c.values[number]
	c.mutex.Unlock()
	if value == 0 {
		return value, ErrNotFound
	} else {
		return value, nil
	}
}

func (c *computedValues) Add(number int, value int) {
	c.mutex.Lock()
	c.values[number] = value
	c.mutex.Unlock()
}

func WithCache(number int, cache *computedValues) int {
	if number == 0 || number == 1 {
		return number
	}

	// In case when we already know the result
	currentNumberCache, currentNumberErr := cache.Get(number)
	if currentNumberErr == nil {
		return currentNumberCache
	}

	firstPrecedingValue, err := cache.Get(number - 1)
	if err != nil {
		firstPrecedingValue = WithCache(number-1, cache)
		cache.Add(number-1, firstPrecedingValue)
	}

	secondPrecedingValue, err := cache.Get(number - 2)
	if err != nil {
		secondPrecedingValue = WithCache(number-2, cache)
		cache.Add(number-1, secondPrecedingValue)
	}

	if currentNumberErr != nil {
		cache.Add(number, firstPrecedingValue+secondPrecedingValue)
	}

	return firstPrecedingValue + secondPrecedingValue
}

func Sequence(n int) []int {
	x, y := 1, 1
	results := make([]int, n)
	for i := 0; i < n; i++ {
		results[i] = x
		x, y = y, x+y
	}
	return results
}
