package utils

import (
	"encoding/hex"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateRandomString(length int) string {
	rand.Seed(uint64(time.Now().Unix()))
	randStr := make([]byte, length)
	rand.Read(randStr)
	return hex.EncodeToString(randStr)
}
