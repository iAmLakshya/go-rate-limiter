# go-rate-limiter

a simple rate limiter written in go. tracks requests per ip using a thread-safe in-memory store.

## how it works

spins up an http server on port 8080 and counts requests from each ip address.

## run it

```bash
go run main.go
```

## licence

mit
