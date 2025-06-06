# Portfolio

This website is made with Go and is designed to replicate the appearance of my Linux desktop environment, mainly <a href="https://github.com/Alexays/Waybar">@waybar</a>. It is fully synchronized with my <a href="https://github.com/mihasket/dotfiles">@dotfiles</a>, ensuring that any changes I push to my configuration are immediately reflected in the website’s look and feel.

## Generate TLS certificate

Using `crypto/tls` from go

```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

You can also use <a href="https://github.com/FiloSottile/mkcert">mkcert</a>.

## Local build

You can use <a href="https://github.com/air-verse/air">air</a> for live reloading.
```bash
air
```

Or normal build
```bash
# addr flag is optional, defaults to 4000

# Just run the web app
go run ./cmd/web -addr=":8000"

# Or build and run
go build -o ./tmp/main ./cmd/web/
./main -addr="8000"
```

## Docker build

Build using:
```bash
docker build -t portfolio .
```

Run using:
```bash
# where <path to your folder with certificates>:/root/tls
docker run -p 8080:8080 -v ./tls:/root/tls portfolio
```
