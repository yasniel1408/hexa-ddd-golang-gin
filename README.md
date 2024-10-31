# 1- Generate doc swagger

### ver en http://localhost:8080/api/swagger/index.html

```
swag init -g cmd/main.go
```

# 2- Correr el proyecto

```
go run cmd/main.go
```

# 3- hot reload

go install github.com/cosmtrek/air@latest
go install github.com/air-verse/air@latest

### generar esto la primera ves

go build -o ./tmp/main ./cmd/main.go

# 4- resolver problemas de paquete

```
go mod tidy
```
