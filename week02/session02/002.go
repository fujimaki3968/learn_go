package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	return rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
}
