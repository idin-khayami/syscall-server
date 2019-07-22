#For Run Server
```go run server-with-syscall.go```

#To Check In Other Terminal Run Below Commands
```ss -pantl ```

#Curl
```curl 127.0.0.1:1074```

#To Benchmark
```go get -u github.com/codesenberg/bombardier```

```bombardier -c 125 -n 1000 127.0.0.1:1074```
