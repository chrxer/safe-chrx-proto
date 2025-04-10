package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

func _req(s []byte, port int, endpoint string) []byte {
	endpoint_url := "http://localhost:" + strconv.Itoa(port) + "/" + endpoint
	
	// Set up the HTTP client with a 10-second timeout
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// Try for maximum 10 s
	start := time.Now()
	var resp *http.Response
	for time.Since(start).Seconds() < 60 {
		req, err := http.NewRequest(http.MethodPost, endpoint_url, bytes.NewBuffer(s))
		if err != nil {
			fmt.Printf("%s", err.Error())
			return []byte("")
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err = client.Do(req)
		if err == nil && resp.StatusCode == 200 {
			// Successfully got a valid response
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return []byte("")
			}
			return body
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Failed to connect after 60 seconds\n")
	return []byte("")
}

func serverEncrypt(s []byte,port int, connKey []byte) []byte {
	s = encrypt(s, connKey)
	res:= _req(s, port,"encrypt")
	return decrypt(res, connKey)
}

func serverDecrypt(s []byte, port int, connKey []byte) []byte {
	s = encrypt(s, connKey)
	res:=_req(s, port, "decrypt")
	return decrypt(res, connKey)
}

func main() {
	port:= flag.Int("port", 3333, "port on which server runs")
	var serverPath = flag.String("server", "", "(optional) Path to server executable")
	connBase64Key:=flag.String("conn-key", "", "(optional) base64 encoded 256bit AES connection key for testing. Won't start server process if provided.")
	flag.Parse()

	// pass connection encryption key to server
	var connKey []byte = []byte("")
	if len(*connBase64Key) != 0{
		connKey, _ = base64.StdEncoding.DecodeString(*connBase64Key)
	}
	if len(connKey) == 0{
		connKey = randKey()
	}
	
	encodedKey := base64.StdEncoding.EncodeToString(connKey)
	println("Connection AES key: ", encodedKey)

	// start the server
	if len(*serverPath) == 0 {
		executablePath, err := os.Executable()
		if err != nil {
			fmt.Printf("Error getting executable path: %v\n", err)
			return
		}
		p := filepath.Join(filepath.Dir(executablePath), "server", "chrxCryptServer")
		serverPath = &p
	}
	
	var cmd *exec.Cmd
	if len(*connBase64Key) == 0 {
		cmd = exec.Command(*serverPath, fmt.Sprintf("--port=%d", *port))
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // Ensure process can be killed by signal
		stdin, err := cmd.StdinPipe()
		if err != nil {
			fmt.Printf("Error setting up stdin pipe: %v\n", err)
			return
		}
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatalf("Error setting up stdout pipe: %v\n", err)
		}
		stderrPipe, err := cmd.StderrPipe()
		if err != nil {
			log.Fatalf("Error setting up stderr pipe: %v\n", err)
		}
		
		err = cmd.Start()
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
			return
		}

		// read one line from server (wait till ready to receive)
		stdoutReader := bufio.NewReader(stdoutPipe)
		line, err := stdoutReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from stdout: %v\n", err)
			cmd.Process.Kill()
			return
		}
		fmt.Printf("Server says: %s", line)

		_, err = stdin.Write([]byte(encodedKey + "\n"))
		if err != nil {
			// Print and return error
			fmt.Println("Error writing to stdin:", err)
			if err := cmd.Process.Kill(); err != nil {
				fmt.Printf("Error killing server: %v\n", err)
			}
		}
		fmt.Printf("%s\n", encodedKey)
		go logOutput(stdoutPipe, "stdout")
		go logOutput(stderrPipe, "stderr")
	}
	
	
	// test encryption & decryption
	data := []byte("password")
	var decrypted = []byte("")
	var encrypted = []byte("")
	encrypted = serverEncrypt(data, *port, connKey)
	fmt.Printf("\nPlaintext: %s\nEncrypted (hex): %x (%s)\n", data, encrypted,encrypted)
	serverDecrypt(encrypted, *port, connKey)
	serverDecrypt(encrypted, *port, connKey)
	decrypted = serverDecrypt(encrypted,*port, connKey)

	fmt.Printf("Decrypted: %s\n\n", decrypted)

	// kill server
	if len(*connBase64Key) == 0 {
		if err := cmd.Process.Kill(); err != nil {
		fmt.Printf("Error killing server: %v\n", err)
		}
	}
}

func logOutput(reader io.Reader, outputType string) {
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from %s: %v\n", outputType, err)
			}
			return
		}
		fmt.Printf("chrxCryptServer: %s", string(buf[:n]))
	}
}
