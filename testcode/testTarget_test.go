package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	rst := square(4)
	if rst != 16 {
		t.Errorf("square(4)=%d; want %d", rst, 4)
	}
}
