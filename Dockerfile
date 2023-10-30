# Usar una imagen de Go como base para compilar la aplicación
FROM golang:1.19-alpine AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente de la aplicación al directorio de trabajo
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Crear una imagen de scratch para reducir el tamaño de la imagen
FROM scratch

# Copiar el ejecutable de la aplicación desde la imagen del compilador a la imagen scratch
COPY --from=builder /app/myapp /myapp

# Establecer el comando de entrada (entrypoint) para la imagen scratch
ENTRYPOINT ["/myapp"]