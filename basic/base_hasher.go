package basic

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
)

type IHasher interface {
	GetHash(content string) (string, error)
}

type HasherSha1 struct{}

func (HasherSha1) GetHash(content string) (string, error) {
	hasher := sha1.New()
	_, err := hasher.Write([]byte(content))

	if err != nil {
		log.Print(err)
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
