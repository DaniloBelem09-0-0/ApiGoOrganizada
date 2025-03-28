# ApiGoOrganizada

Este projeto é uma API RESTful desenvolvida em **Go (Golang)** utilizando o framework **Gin** e o banco de dados **MongoDB**. A API segue uma estrutura organizada e modular, facilitando a manutenção e escalabilidade.

## 🚀 Tecnologias Utilizadas

- **Go (Golang)**
- **Gin (Web Framework)**
- **MongoDB (Banco de Dados NoSQL)**
- **Swagger (Documentação da API)**
- **Docker (Containerização)**
- **godotenv (Gerenciamento de variáveis de ambiente)**

## 📁 Estrutura do Projeto

```
ApiGoOrganizada/
│-- dependencies/            # Injeção de dependências
│-- docs/                    # Documentação gerada pelo Swagger
│-- src/
│   ├── configuration/
│   │   ├── database/        # Configuração do banco MongoDB
│   │   ├── logger/          # Configuração de logs
│   ├── controller/
│   ├── routes/              # Definição de rotas da API
│   ├── models/              # Modelos de dados
│-- .env                     # Variáveis de ambiente
│-- Dockerfile               # Configuração para Docker
│-- go.mod                   # Dependências do Go
│-- go.sum                   # Checksum das dependências
│-- main.go                  # Ponto de entrada da aplicação
```

## 🔧 Instalação e Execução

### 1️⃣ Clonar o repositório
```sh
git clone https://github.com/DaniloBelem09-0-0/ApiGoOrganizada.git
cd ApiGoOrganizada
```

### 2️⃣ Configurar o ambiente
Crie um arquivo `.env` na raiz do projeto e configure as variáveis necessárias:
```ini
MONGO_URI=mongodb://localhost:27017
DATABASE_NAME=meubanco
PORT=8080
```

### 3️⃣ Instalar as dependências
```sh
go mod download
```

### 4️⃣ Gerar a documentação Swagger
```sh
swag init
```

### 5️⃣ Rodar a aplicação
```sh
go run main.go
```

A API estará disponível em **http://localhost:8080** e a documentação em **http://localhost:8080/swagger/index.html**.

## 🐳 Executando com Docker

### Construir a imagem Docker
```sh
docker build -t apigo .
```

### Criar e rodar um container
```sh
docker run -d -p 8080:8080 --name apigo-container apigo
```

### Ver logs do container
```sh
docker logs -f apigo-container
```

## 📖 Documentação da API

A documentação da API é gerada automaticamente com **Swagger**. Para acessar, após iniciar o servidor, visite:

🔗 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

## 📌 Endpoints Principais

| Método  | Rota            | Descrição          |
|---------|----------------|--------------------|
| `GET`   | `/users`       | Lista todos usuários |
| `POST`  | `/users`       | Cria um novo usuário |
| `GET`   | `/users/:id`   | Busca um usuário pelo ID |
| `PUT`   | `/users/:id`   | Atualiza um usuário |
| `DELETE`| `/users/:id`   | Remove um usuário |

## ✨ Contribuições

Contribuições são bem-vindas! Para contribuir:
1. **Fork** este repositório
2. Crie um **branch** (`git checkout -b minha-feature`)
3. Faça um **commit** (`git commit -m 'Minha nova feature'`)
4. Envie para o **remote** (`git push origin minha-feature`)
5. Abra um **Pull Request** 🚀

---

Feito por Danilo Belém

