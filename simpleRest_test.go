package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestVersionHandler(t *testing.T) {
	if strings.Compare(version, "v1.0.0") != 0 {
		t.Fatal("The expected and actual versions don't match...")
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	randomNumber := randomGenerator.Intn(100)
	if randomNumber%3 == 0 {
		t.Fatalf("The random number %d is multiple of 3...", randomNumber)
	}
}
