from docker.arvancloud.ir/golang:1.22.0-alpine3.19
expose 1111
workdir /app
copy main.go main.go 
run go mod init myapp 
run go mod tidy
cmd go run main.go 
