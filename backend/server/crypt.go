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

// import () if necessary

func encrypt(b []byte) []byte {
	if len(b) == 0 {
		return []byte("")
	}

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
	fmt.Printf("cleartext:%s\nciphertext: %s\n", b, ciphertext)

	return ciphertext
}

func decrypt(b []byte) []byte {
	if len(b) == 0 {
		return []byte("")
	}

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
	fmt.Printf("decrypt: %v", plaintext)
	return plaintext
}

func getMasterPassword() []byte {
	/* FOR TESTING PURPOSES */
	// masterKey = NewSHA256([]byte("a"))
	/* ********************* */

	if len(masterKey) == 0 {
		for(len(userPassword) == 0) {
			wg.Add(1)
			go func() {
				myWindow.Show()
			}()
			wg.Wait() // wg.Done() is run when the correct password is given by the user on Main()
		}
		masterKey = NewSHA256([]byte(userPassword))
	}
	return masterKey
}

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

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
    fmt.Println("file written successfully.")
}