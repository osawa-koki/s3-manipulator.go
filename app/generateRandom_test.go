package app_test

import (
	"example.com/app"
	"testing"
)

func TestGenerateRandom(t *testing.T) {
	min := -10
	max := 10
	for i := 0; i < 100; i++ {
		i := app.GenerateRandom(min, max)
		if i < min || i >= max {
			t.Errorf("GenerateRandom(%d, %d) = %d", min, max, i)
		}
	}
}
