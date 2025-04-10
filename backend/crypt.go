package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func encrypt(b []byte, mP []byte) []byte {
	if len(b) == 0 {
		return []byte("")
	}
	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(mP)
	if err != nil {
		panic(err.Error())
	}

	// Create a new GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypt the data using aesGCM.Seal
	// We don't save the nonce (in a database) -> add it as a prefix to the encrypted data -> first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, b, nil)

	return ciphertext
}

func decrypt(b []byte, mP []byte) []byte {
	if len(b) == 0 {
		return []byte("")
	}

	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(mP)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data (-> it is the prefix of the encrypted password)
	nonce, ciphertext := b[:nonceSize], b[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func randKey() []byte {
	// 256-bit AES key (32 bytes)
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}