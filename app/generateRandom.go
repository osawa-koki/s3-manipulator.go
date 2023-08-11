package app

import (
	"math/rand"
	"time"
)

/*
`i1`と`i2`の間のランダムな整数を返します。
ただし、`i1`は含み、`i2`は含みません。
*/
func GenerateRandom(i1, i2 int) int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return r.Intn(i2-i1) + i1
}
