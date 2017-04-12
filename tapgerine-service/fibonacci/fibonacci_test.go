// +build unit

package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NaiveImplementations(t *testing.T) {
	assert.Equal(t, NaiveImplementations(10), 55)
	assert.Equal(t, NaiveImplementations(20), 6765)
	assert.Equal(t, NaiveImplementations(30), 832040)
}

func Test_Sequence(t *testing.T) {
	s1 := Sequence(10)
	s2 := Sequence(20)
	s3 := Sequence(30)
	assert.Equal(t, s1[len(s1)-1], 55)
	assert.Equal(t, s2[len(s2)-1], 6765)
	assert.Equal(t, s3[len(s3)-1], 832040)
}

func Test_WithCache(t *testing.T) {
	cache := CreateCache()
	assert.Equal(t, WithCache(10, cache), 55)
	assert.Equal(t, WithCache(20, cache), 6765)
	assert.Equal(t, WithCache(30, cache), 832040)
	assert.Equal(t, WithCache(92, cache), 7540113804746346429)
	assert.Equal(t, WithCache(30, cache), 832040)
}
