# 🚀 Instruções para Executar a API da Biblioteca

## 📋 Pré-requisitos

1. **Go 1.23+** instalado
2. **PostgreSQL** rodando
3. **MongoDB** rodando
4. Arquivo `.env` configurado

## ⚙️ Configuração

### 1. Configure o arquivo `.env`

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
# PostgreSQL
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=seu_usuario
POSTGRES_PASSWORD=sua_senha
POSTGRES_DB=biblioteca

# MongoDB
MONGO_URI=mongodb://localhost:27017
MONGO_DB=bibliotecaDB

# API
PORT=8080
```

### 2. Prepare os bancos de dados

#### PostgreSQL
- Certifique-se de que o banco `biblioteca` existe
- As tabelas serão criadas automaticamente quando necessário
- Para testes, você pode usar o CNPJ `11222333000144` para editoras

#### MongoDB
- O banco `bibliotecaDB` será criado automaticamente
- As coleções serão criadas conforme necessário

## 🏃‍♂️ Executando a API

### Opção 1: Executar diretamente
```bash
go run main.go
```

### Opção 2: Compilar e executar
```bash
go build -o biblioteca-api.exe
./biblioteca-api.exe
```

### Opção 3: Usando go install
```bash
go install
biblioteca-api
```

## 🌐 Acessando a API

Após executar, a API estará disponível em:
- **URL Base**: `http://localhost:8080`
- **Health Check**: `http://localhost:8080/health`

## 🧪 Testando a API

### 1. Health Check
```bash
curl http://localhost:8080/health
```

### 2. Usando o arquivo de testes
O arquivo `test-api.http` contém exemplos de todas as requisições. Você pode usar:

- **VS Code**: Instale a extensão "REST Client" e execute as requisições diretamente
- **Postman**: Importe as requisições
- **cURL**: Copie e cole os comandos

### 3. Exemplos rápidos

#### Criar um usuário (PostgreSQL)
```bash
curl -X POST http://localhost:8080/postgres/usuarios \
  -H "Content-Type: application/json" \
  -d '{
    "cpf": "12345678901",
    "primeiro_nome": "João",
    "sobrenome": "Silva",
    "data_nascimento": "1990-01-01T00:00:00Z"
  }'
```

#### Listar todos os livros (MongoDB)
```bash
curl http://localhost:8080/mongo/livros
```

## 📊 Estrutura das Respostas

### Sucesso (200/201)
```json
{
  "message": "Operação realizada com sucesso",
  "data": { ... }
}
```

### Erro (400/500)
```json
{
  "error": "Descrição do erro"
}
```

## 🔍 Endpoints Disponíveis

### PostgreSQL
- `GET /postgres/usuarios` - Listar usuários
- `POST /postgres/usuarios` - Criar usuário
- `PUT /postgres/usuarios/:cpf` - Atualizar usuário
- `DELETE /postgres/usuarios/:cpf` - Deletar usuário

- `GET /postgres/autores` - Listar autores
- `POST /postgres/autores` - Criar autor
- `PUT /postgres/autores/:id` - Atualizar autor
- `DELETE /postgres/autores/:id` - Deletar autor

- `GET /postgres/livros` - Listar livros
- `POST /postgres/livros` - Criar livro
- `PUT /postgres/livros/:isbn` - Atualizar livro
- `DELETE /postgres/livros/:isbn` - Deletar livro

- `GET /postgres/emprestimos` - Listar empréstimos
- `POST /postgres/emprestimos` - Criar empréstimo
- `PUT /postgres/emprestimos/:id` - Atualizar empréstimo
- `DELETE /postgres/emprestimos/:id` - Deletar empréstimo

### MongoDB
- `GET /mongo/usuarios` - Listar usuários
- `POST /mongo/usuarios` - Criar usuário
- `PUT /mongo/usuarios/:id` - Atualizar usuário
- `DELETE /mongo/usuarios/:id` - Deletar usuário

- `GET /mongo/autores` - Listar autores
- `POST /mongo/autores` - Criar autor
- `PUT /mongo/autores/:id` - Atualizar autor
- `DELETE /mongo/autores/:id` - Deletar autor

- `GET /mongo/livros` - Listar livros
- `POST /mongo/livros` - Criar livro
- `PUT /mongo/livros/:id` - Atualizar livro
- `DELETE /mongo/livros/:id` - Deletar livro

- `GET /mongo/emprestimos` - Listar empréstimos
- `POST /mongo/emprestimos` - Criar empréstimo
- `PUT /mongo/emprestimos/:id` - Atualizar empréstimo
- `DELETE /mongo/emprestimos/:id` - Deletar empréstimo

## 🐛 Troubleshooting

### Erro de conexão com PostgreSQL
- Verifique se o PostgreSQL está rodando
- Confirme as credenciais no arquivo `.env`
- Teste a conexão: `psql -h localhost -U seu_usuario -d biblioteca`

### Erro de conexão com MongoDB
- Verifique se o MongoDB está rodando
- Confirme a URI no arquivo `.env`
- Teste a conexão: `mongosh`

### Erro de compilação
- Execute `go mod tidy` para resolver dependências
- Verifique se o Go 1.23+ está instalado: `go version`

### Erro de porta em uso
- Mude a porta no arquivo `.env`
- Ou mate o processo que está usando a porta 8080

## 📱 Integração com Frontend

A API está configurada com CORS habilitado, permitindo requisições de qualquer origem. Para integrar com um frontend:

```javascript
// Exemplo de requisição JavaScript
fetch('http://localhost:8080/postgres/usuarios', {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json'
  }
})
.then(response => response.json())
.then(data => console.log(data));
```

## 🎯 Próximos Passos

1. **Teste todos os endpoints** usando o arquivo `test-api.http`
2. **Crie um frontend** para consumir a API
3. **Implemente autenticação** se necessário
4. **Adicione validações** mais robustas
5. **Implemente testes** automatizados

---

**🎉 Sua API está pronta para uso!**
