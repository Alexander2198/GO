# Usa una imagen base de Go
FROM golang:1.20-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo main.go al contenedor
COPY main.go .

# Compila la aplicación
RUN go build -o app main.go

# Expone el puerto que usará la aplicación (8080)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./app"]

