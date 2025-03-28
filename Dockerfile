# Etapa de Build
FROM golang:1.24.1 AS builder

WORKDIR /app

# Copiar todos os arquivos necessários para o contêiner
COPY src/ /app/src/
COPY docs/ /app/docs/
COPY dependencies/ /app/dependencies/
COPY go.mod go.sum /app/
COPY .env /app/

# Baixar as dependências do Go
RUN go mod download

# Copiar e compilar o código
COPY main.go /app/
COPY dependencies/init_dependencies.go /app/dependencies/
RUN go build -o apigo main.go

# Etapa de Execução
FROM golang:1.24.1 AS runner

WORKDIR /app

# Copiar o binário compilado da etapa anterior
COPY --from=builder /app/apigo /app/apigo

# Expor a porta 8080
EXPOSE 8080

# Executar o binário apigo
CMD ["./apigo"]
