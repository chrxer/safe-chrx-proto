package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// import () if necessary

func encrypt(b []byte) []byte {
	mP := getMasterPassword()
	//Create a new Cipher Block from the key

	block, err := aes.NewCipher(mP)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, b, nil)
	return ciphertext
}

func decrypt(b []byte) []byte {
	mP := getMasterPassword()
	block, err := aes.NewCipher(mP)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := b[:nonceSize], b[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}

var mockKey = make([]byte, 32)

func getMasterPassword() []byte {
	if len(masterPassword) == 0 {
		// get Masterpassword
		// return []byte("ThisIsSomeReallyLongKey")
		return mockKey
	} else {
		return masterPassword
	}
}
