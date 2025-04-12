package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/emersion/go-appdir"
)

func encrypt(b []byte, mP []byte) []byte {
	if len(b) == 0 {
		return []byte("")
	}
	if len(mP) == 0{
		mP = getMasterPassword()
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
	if len(mP) == 0{
		mP = getMasterPassword()
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

func getMasterPassword() []byte {
	if bytes.Equal(masterKey,[]byte("")) {
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
		panic(err.Error())
	}
	return hash
}

func argonCheckPswd(pswd string, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(pswd, hash)
	if err != nil {
		panic(err.Error())
	}
	return match
}

/* FILE (password) read & write */

func getHashFile() string {
    dirs := appdir.New("chrx-safe-proto")
	p := dirs.UserConfig()
	if err := os.MkdirAll(p, 0700); err != nil {
		panic(err.Error())
	}
	fpath:= filepath.Join(p, "argon2.hash")
	
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
        f, err := os.Create(fpath)
        if err != nil {
            panic(err.Error())
        }
        defer f.Close()
    }

	return fpath
}

func fetchHash() string {
	
	hashf:=getHashFile()
	dat, err := os.ReadFile(hashf)
	if err != nil {
        panic(err.Error())
    }
    return string(dat)
}

func writeHash(hash string) {
	data := []byte(hash)
    err := os.WriteFile(getHashFile(), data, 0600)
    if err != nil {
        panic(err.Error())
    }
}

func readAESKeyFromStdin() []byte {
    // Create buffered reader for more reliable reading
    reader := bufio.NewReader(os.Stdin)
    
    // Read with timeout support
    result := make(chan []byte)
    go func() {
        line, err := reader.ReadString('\n')
        if err != nil {
            panic(err.Error())
        }
        
        key, err := base64.StdEncoding.DecodeString(strings.TrimSpace(line))
        if err != nil {
            panic(err.Error())
        }
        
        if len(key) != 32 {
            panic(err.Error())
        }
        
        result <- key
    }()

    // Wait for either the result or timeout
    select {
    case key := <-result:
        return key
    case <-time.After(5 * time.Second):
        panic("Timed out waiting for AES key input")
		panic("")
    }
}