# TODO: Fix crypt.go

Current error:
```
go run .
Encrypting..
2025/02/16 14:45:39 http: panic serving [::1]:32780: cipher: message authentication failed
goroutine 7 [running]:
net/http.(*conn).serve.func1()
        /usr/local/go/src/net/http/server.go:1947 +0xbe
panic({0x691a60?, 0xc0000143a0?})
        /usr/local/go/src/runtime/panic.go:787 +0x132
main.decrypt({0xc00018c000, 0x15, 0x200})
        /home/user/VSCProjects/chrxer/safe-chrx-proto/backend/server/crypt.go:61 +0x13f
main.handleEncrypt({0x762bb0, 0xc0000000e0}, 0xc000180000)
        /home/user/VSCProjects/chrxer/safe-chrx-proto/backend/server/main.go:25 +0x10d
net/http.HandlerFunc.ServeHTTP(0x9505a0?, {0x762bb0?, 0xc0000000e0?}, 0xc000187b60?)
        /usr/local/go/src/net/http/server.go:2294 +0x29
net/http.(*ServeMux).ServeHTTP(0x416e85?, {0x762bb0, 0xc0000000e0}, 0xc000180000)
        /usr/local/go/src/net/http/server.go:2822 +0x1c4
net/http.serverHandler.ServeHTTP({0x761ac0?}, {0x762bb0?, 0xc0000000e0?}, 0x1?)
        /usr/local/go/src/net/http/server.go:3301 +0x8e
net/http.(*conn).serve(0xc0000cc240, {0x762fd0, 0xc00009b380})
        /usr/local/go/src/net/http/server.go:2102 +0x625
created by net/http.(*Server).Serve in goroutine 1
        /usr/local/go/src/net/http/server.go:3454 +0x485
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
