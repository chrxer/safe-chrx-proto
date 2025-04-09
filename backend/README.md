## Dependencies
Requires [golang](https://go.dev/doc/install) to be installed.

#### Linux/Debian
```bash
sudo ./deps.sh
```

## Running

start server (in ./backend/server)

```bash
cd server
go run chrxCryptServer
```

start client (in ./backend)

```bash
go run serverTest
```

## Building
```bash
go build $package
./$package
``` 

# References

- [AES-256 encrypt/decrypt in Go](https://gist.github.com/donvito/efb2c643b724cf6ff453da84985281f8)
- [fyne-io/fyne](https://github.com/fyne-io/fyne)
