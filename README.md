
# Trabalho Prático de Banco de Dados - CRUD Biblioteca

## Integrantes do Grupo
- Patrick Oliveira
- João Pedro
- Paulo Henrique
- Lucas Santana

## Objetivo da Aplicação
Desenvolver um sistema de gerenciamento de biblioteca que realiza operações de CRUD (criar, ler, atualizar e deletar) para autores, livros, usuários e empréstimos, atendendo aos requisitos da disciplina de Banco de Dados. O sistema permite manipulação dos dados tanto em banco relacional (PostgreSQL) quanto NoSQL (MongoDB), incluindo o relacionamento entre livros e autores e o registro de empréstimos.

## Tecnologias Utilizadas
- Go (Golang)
- PostgreSQL
- MongoDB

## Estrutura do Projeto
```
go.mod
go.sum
main.go
database/
  database.go
model/
  models.go
repository/
  interfaces.go
  mongo/
    mongo_autor.go
    mongo_livro.go
    mongo_usuario.go
    mongo_emprestimo.go
  postgres/
    postgres_autor.go
    postgres_livro.go
    postgres_usuario.go
    postgres_emprestimo.go
```

## Como Configurar e Executar o Projeto
1. **Pré-requisitos:**
   - Instale o Go: https://go.dev/dl/
   - Instale o PostgreSQL e o MongoDB em sua máquina ou utilize instâncias remotas.

2. **Configuração do Banco de Dados:**
   - Crie as tabelas necessárias no PostgreSQL (consulte o script ou modelo do projeto).
   - Cadastre uma editora com o CNPJ `11222333000144` para testes.

3. **Configuração do Projeto:**
   - Crie um arquivo `.env` na raiz do projeto com a string de conexão do PostgreSQL:
     ```
     POSTGRES_CONN=postgresql://usuario:senha@host:5432/database?sslmode=require
     ```
   - O arquivo `.env` já está protegido pelo `.gitignore` e não será enviado ao GitHub.

4. **Executando o Projeto:**
   - No terminal, navegue até a pasta do projeto e execute:
     ```
     go run main.go
     ```
   - Siga o menu interativo para realizar as operações de CRUD.
   - O menu inclui opções para:
     - Usuário: criar, ler, atualizar, deletar
     - Livro: criar, ler, deletar
     - Autor: criar, ler, deletar, relacionar com livro
     - Empréstimo: criar, ler, atualizar, deletar

## CRUD de Empréstimo
No menu principal, utilize as opções 10 a 13 para:
- Criar empréstimo: informe ID (int), status (A/D/C), quantidade de livros, CPF do cliente/usuário
- Ler empréstimo por ID
- Atualizar empréstimo
- Deletar empréstimo

O campo status aceita apenas os valores: 'A', 'D', 'C'. O CPF deve existir na tabela Cliente.

## Observações
- Para testar todos os métodos, utilize o menu do sistema e confira o efeito das operações diretamente no banco de dados (usando pgAdmin, DBeaver ou MongoDB Compass).
- O projeto foi desenvolvido para fins acadêmicos e pode ser adaptado conforme a necessidade.

---
Desenvolvido para a disciplina de Banco de Dados - UFS
