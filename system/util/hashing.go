package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/blake2b"

	"crypto/sha256"
)

func Sha256(data []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(data))
}

// Return a compact but collision resistant hash
func Sha256Base64(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func Sha256RandomToken(seeds ...string) string {
	seed := "379^43_40 n /a-#!h"
	if len(seeds) > 0 {
		seed = seeds[0]
	}
	str := fmt.Sprintf("~(%s&>$)%d_$@!-*", seed, time.Now().Unix())
	return Sha256([]byte(str))
}

// Adapted from Beego
func RandomToken(ln int) (token string, err error) {
	b := make([]byte, ln)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return token, errors.New("Could not read from the system CSPRNG - " + err.Error())
	}
	return hex.EncodeToString(b), nil
}

func Base64RandomToken(length int) (token string, err error) {
	b := make([]byte, length)
	n, err := rand.Read(b)
	if err != nil {
		return token, errors.New("Could not read from the system CSPRNG - " + err.Error())
	}
	if n != len(b) {
		return token, errors.New("Length of bytes read from rand does not match expected length")
	}
	return base64.URLEncoding.EncodeToString(b), err
}

func Blake384(data string) string {
	h := blake2b.Sum384([]byte(data))
	return base64.URLEncoding.EncodeToString(h[:])
}

func Blake512(data string) string {
	h := blake2b.Sum512([]byte(data))
	return base64.URLEncoding.EncodeToString(h[:])
}
