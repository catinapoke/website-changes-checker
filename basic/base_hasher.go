package basic

import (
	"crypto/sha1"
	"encoding/hex"
)

type IHasher interface {
	GetHash(content string) string
}

type HasherSha1 struct{}

func (HasherSha1) GetHash(content string) string {
	hasher := sha1.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}
