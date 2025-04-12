# Usage

```bash
Usage of ./chrxCryptServer:
  -port int
        port to serve on (default 3333)
  -reset
        Reset the password. All currently encrypted data will be lost
```
A AES265 key has to be passed base64 encoded to the server over `stdin`

```
Usage of ./serverTest:
Usage of ./serverTest:
  -conn-key string
        (optional) base64 encoded 256bit AES connection key for testing. Won't start server process if provided.
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

### Logging
Logs can be found in the default `%tmp%` directory, under the filename `chrxCryptServerLog.log` (`/tmp/chrxCryptServerLog.log` on Linux)

# References

- [AES-256 encrypt/decrypt in Go](https://gist.github.com/donvito/efb2c643b724cf6ff453da84985281f8)
- [fyne-io/fyne](https://github.com/fyne-io/fyne)
