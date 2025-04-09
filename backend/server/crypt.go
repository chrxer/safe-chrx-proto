package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/alexedwards/argon2id"
)

func encrypt(b []byte) []byte {
	if len(b) == 0 {
		return []byte("")
	}
	mP := getMasterPassword()
	
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

func decrypt(b []byte) []byte {
	if len(b) == 0 {
		return []byte("")
	}
	mP := getMasterPassword()

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

func getMasterPassword() []byte {
	if len(masterKey) == 0 {
		// In case the user attempts closing the window
		for(len(userPassword) == 0) {
			wg.Add(1)
			myWindow.Show()
			wg.Wait() // wg.Done() is run on Main() on correct password given or if the window is closed (=> reason for the for loop)
		}
		masterKey = NewSHA256([]byte(userPassword))
	}
	return masterKey
}

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

/* ARGON2ID */

func argonHash(pswd string) string {
	hash, err := argon2id.CreateHash(pswd, argon2id.DefaultParams)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	return hash
}

func argonCheckPswd(pswd string, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(pswd, hash)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	return match
}

/* FILE (password) read & write */

func fetchHash() string {
	dat, err := os.ReadFile("./password.txt")
	if err != nil {
        panic(err)
    }
    return string(dat)
}

func writeHash(hash string) {
	data := []byte(hash)
    err := os.WriteFile("./password.txt", data, 0666)
    if err != nil {
        fmt.Printf("%s", err.Error())
    }
}