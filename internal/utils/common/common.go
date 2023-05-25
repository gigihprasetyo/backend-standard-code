package common

import (
	"math/rand"

	"github.com/google/uuid"
)

func Unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func CheckFreeShipping(sellerCity uint64) bool {
	isDiscounted := false

	return isDiscounted
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
