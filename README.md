
# 🏛️ API da Biblioteca - CRUD Completo

API REST completa para gerenciamento de biblioteca com suporte a PostgreSQL e MongoDB.

## 🚀 Funcionalidades

- **CRUD Completo** para 4 entidades: Usuários, Autores, Livros e Empréstimos
- **Suporte a 2 Bancos de Dados**: PostgreSQL e MongoDB
- **API REST** com endpoints padronizados
- **CORS habilitado** para integração com frontend
- **Validação de dados** e tratamento de erros

## 📋 Tabela Completa de Rotas

### 🔵 PostgreSQL Endpoints

#### Usuários
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/postgres/usuarios` | Listar todos os usuários |
| POST | `/postgres/usuarios` | Criar novo usuário |
| PUT | `/postgres/usuarios/:cpf` | Atualizar usuário |
| DELETE | `/postgres/usuarios/:cpf` | Deletar usuário |

#### Autores
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/postgres/autores` | Listar todos os autores |
| POST | `/postgres/autores` | Criar novo autor |
| PUT | `/postgres/autores/:id` | Atualizar autor |
| DELETE | `/postgres/autores/:id` | Deletar autor |

#### Livros
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/postgres/livros` | Listar todos os livros |
| POST | `/postgres/livros` | Criar novo livro |
| PUT | `/postgres/livros/:isbn` | Atualizar livro |
| DELETE | `/postgres/livros/:isbn` | Deletar livro |

#### Empréstimos
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/postgres/emprestimos` | Listar todos os empréstimos |
| POST | `/postgres/emprestimos` | Criar novo empréstimo |
| PUT | `/postgres/emprestimos/:id` | Atualizar empréstimo |
| DELETE | `/postgres/emprestimos/:id` | Deletar empréstimo |

### 🟢 MongoDB Endpoints

#### Usuários
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/mongo/usuarios` | Listar todos os usuários |
| POST | `/mongo/usuarios` | Criar novo usuário |
| PUT | `/mongo/usuarios/:id` | Atualizar usuário |
| DELETE | `/mongo/usuarios/:id` | Deletar usuário |

#### Autores
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/mongo/autores` | Listar todos os autores |
| POST | `/mongo/autores` | Criar novo autor |
| PUT | `/mongo/autores/:id` | Atualizar autor |
| DELETE | `/mongo/autores/:id` | Deletar autor |

#### Livros
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/mongo/livros` | Listar todos os livros |
| POST | `/mongo/livros` | Criar novo livro |
| PUT | `/mongo/livros/:id` | Atualizar livro |
| DELETE | `/mongo/livros/:id` | Deletar livro |

#### Empréstimos
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/mongo/emprestimos` | Listar todos os empréstimos |
| POST | `/mongo/emprestimos` | Criar novo empréstimo |
| PUT | `/mongo/emprestimos/:id` | Atualizar empréstimo |
| DELETE | `/mongo/emprestimos/:id` | Deletar empréstimo |

## 🛠️ Instalação e Configuração

### Pré-requisitos
- Go 1.23+
- PostgreSQL
- MongoDB
- Arquivo `.env` configurado

### 1. Clone o repositório
```bash
git clone <url-do-repositorio>
cd crud-biblioteca
```

### 2. Configure as variáveis de ambiente
Crie um arquivo `.env` na raiz do projeto:

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

### 3. Execute o projeto
```bash
go mod tidy
go run main.go
```

## 📝 Exemplos de Uso

### Criar um Usuário (PostgreSQL)
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

### Listar Todos os Livros (MongoDB)
```bash
curl -X GET http://localhost:8080/mongo/livros
```

### Atualizar um Autor (PostgreSQL)
```bash
curl -X PUT http://localhost:8080/postgres/autores/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "primeiro_nome": "Maria",
    "sobrenome": "Santos"
  }'
```

### Deletar um Empréstimo (MongoDB)
```bash
curl -X DELETE http://localhost:8080/mongo/emprestimos/1
```

## 🏗️ Estrutura do Projeto

```
crud-biblioteca/
├── main.go                 # Servidor HTTP com Gin
├── handlers/
│   └── handlers.go         # Handlers HTTP para todas as rotas
├── model/
│   └── models.go           # Estruturas de dados
├── repository/
│   ├── interfaces.go       # Interfaces dos repositórios
│   ├── postgres/           # Implementações PostgreSQL
│   └── mongo/              # Implementações MongoDB
├── database/
│   └── connection.go       # Conexões com bancos de dados
└── README.md
```

## 🔧 Tecnologias Utilizadas

- **Go 1.23+** - Linguagem principal
- **Gin** - Framework web
- **PostgreSQL** - Banco relacional
- **MongoDB** - Banco NoSQL
- **pgx** - Driver PostgreSQL
- **mongo-driver** - Driver MongoDB

## 🌐 Endpoints Especiais

### Health Check
```bash
GET /health
```
Retorna o status da API e documentação dos endpoints.

## 📊 Modelos de Dados

### Usuário
```json
{
  "cpf": "string",
  "primeiro_nome": "string",
  "sobrenome": "string",
  "data_nascimento": "datetime"
}
```

### Autor
```json
{
  "id": "integer",
  "primeiro_nome": "string",
  "sobrenome": "string"
}
```

### Livro
```json
{
  "isbn": "string",
  "titulo": "string",
  "edicao": "string",
  "num_paginas": "integer",
  "editora_cnpj": "string",
  "funcionario_matricula": "integer",
  "autores": ["array of Autor"]
}
```

### Empréstimo
```json
{
  "id": "integer",
  "data_emprestimo": "datetime",
  "status": "string",
  "quant_livros": "integer",
  "cliente_usuario_cpf": "string"
}
```

## 🚀 Próximos Passos

1. **Frontend**: Criar interface web para consumir a API
2. **Autenticação**: Implementar JWT ou OAuth
3. **Validação**: Adicionar validações mais robustas
4. **Testes**: Implementar testes unitários e de integração
5. **Logs**: Adicionar sistema de logs estruturados
6. **Documentação**: Swagger/OpenAPI

## 📞 Suporte

Para dúvidas ou problemas, abra uma issue no repositório.

---

**Desenvolvido com ❤️ em Go**
