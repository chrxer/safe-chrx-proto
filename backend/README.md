# Usage

```bash
Usage of ./chrxCryptServer:
  -conn-key string
        (optional) base64 encoded 256bit AES connection key for testing. Should be passed over stdin instead
  -port int
        port to serve on (default 3333)
  -reset
        Reset the password. All currently encrypted data will be lost
```
```
Usage of ./serverTest:
  -conn-key string
        (optional) base64 encoded 256bit AES connection key for testing
  -port int
        port on which server runs (default 3333)
  -server string
        (optional) Path to server executable
```

# Building
### Dependencies
Requires [golang](https://go.dev/doc/install) to be installed.
```bash
go get serverTest
cd server
go get chrxCryptServer
cd ..
```

#### Linux/Debian
```bash
sudo ./deps.sh
```


### Build executable
```bash
go build serverTest
cd server
go build chrxCryptServer
cd ..
``` 

# References

- [AES-256 encrypt/decrypt in Go](https://gist.github.com/donvito/efb2c643b724cf6ff453da84985281f8)
- [fyne-io/fyne](https://github.com/fyne-io/fyne)
