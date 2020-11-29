package main

import (
	b64 "encoding/base64"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
	"os"
)

const HEADER = "XMAS_CRYPTOR_ENCRYPTED_CIPHERTEXT:"
const HEADER_LEN = len(HEADER)

func encrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func decrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func encryptString(key, plaintext string) {
	keyB := []byte(key)
	ciphertextB, err := encrypt(keyB, []byte(plaintext))
	if err != nil {
		log.Fatalf("Error encrypting chiphertext: %v", err)
	}
	fmt.Print(HEADER)
	fmt.Println(b64.StdEncoding.EncodeToString(ciphertextB))
}

func decryptString(key, ciphertext string) {
	if len(ciphertext) < HEADER_LEN || ciphertext[0:HEADER_LEN] != HEADER {
		log.Fatalf("Chiphertext was not encrypted with cryptor.")
	}

	keyB := []byte(key)
	ciphertextB, err := b64.StdEncoding.DecodeString(ciphertext[HEADER_LEN:])
	if err != nil {
		log.Fatalf("Error decoding chiphertext: %v", err)
	}
	plaintext, err := decrypt(keyB, ciphertextB)
	if err != nil {
		log.Fatalf("Error decrypting chiphertext: %v", err)
	}
	fmt.Println(string(plaintext))
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage:\n%s encrypt [key] [plaintext]\n%s decrypt [key] [ciphertext]\n\nExample:\n%s encrypt secure_password \"secret text\"\n", os.Args[0], os.Args[0], os.Args[0])
		os.Exit(2)
	}

	switch os.Args[1] {
	case "encrypt":
		encryptString(os.Args[2], os.Args[3])
	case "decrypt":
		decryptString(os.Args[2], os.Args[3])
	default:
		log.Fatalf("Unknown command: %s\n", os.Args[1])
	}
}
