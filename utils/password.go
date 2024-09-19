package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/pbkdf2"
)

func GetPBKDF2Digest() func() hash.Hash {
	digest := os.Getenv("PBKDF2_DIGEST")
	switch digest {
	case "sha512":
		return sha512.New
	case "sha256":
		return sha256.New
	default:
		log.Fatalf("PBKDF2_DIGEST invalid, must be 'sha256' or 'sha512'")
		return sha512.New
	}
}

func HashPassword(password string) string {
	fmt.Println("Execute function HashPassword")
	salt := []byte(os.Getenv("PBDKF2_SALT_ENCRYPT"))
	iterations, _ := strconv.Atoi(os.Getenv("PBKDF2_ITERATIONS"))
	keyLen, _ := strconv.Atoi(os.Getenv("PBKDF2_KEYLEN"))
	digest := GetPBKDF2Digest()

	// create hashpassword using PBDKF2
	hash := pbkdf2.Key([]byte(password), salt, iterations, keyLen, digest)
	return hex.EncodeToString(hash)
}

func VerifyPassword(password string, dbPassword string) bool {
	fmt.Println("Execute function verifyPassword")

	hashPasswd := HashPassword(password)
	return hashPasswd == dbPassword
}
