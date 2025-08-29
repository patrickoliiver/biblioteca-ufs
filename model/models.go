package model

import "time"

// Usuario representa a tabela/coleção Usuario
type Usuario struct {
	CPF            string    `bson:"_id"` // CPF como ID no Mongo
	DataNascimento time.Time `bson:"data_nascimento"`
	Sobrenome      string    `bson:"sobrenome"`
	PrimeiroNome   string    `bson:"primeiro_nome"`
}

// Autor representa um autor, que será embutido no Livro no modelo NoSQL
type Autor struct {
	ID           int    `bson:"_id"`
	PrimeiroNome string `bson:"primeiro_nome"`
	Sobrenome    string `bson:"sobrenome"`
}

// Livro representa a tabela/coleção Livro
type Livro struct {
	ISBN                 string  `bson:"_id"` // ISBN como ID
	Titulo               string  `bson:"titulo"`
	Edicao               string  `bson:"edicao"`
	NumPaginas           int     `bson:"num_paginas"`
	EditoraCNPJ          string  `bson:"editora_cnpj"`
	FuncionarioMatricula int     `bson:"funcionario_matricula"`
	Autores              []Autor `bson:"autores"` // relacionamento embutido para NoSQL
}

// Emprestimo representa a tabela no banco de dados
// Atualizado para refletir os campos reais
type Emprestimo struct {
	ID                 int       `bson:"_id" json:"id"`
	DataEmprestimo     time.Time `bson:"data_emprestimo" json:"data_emprestimo"`
	Status             string    `bson:"status" json:"status"`
	QuantLivros        int       `bson:"quant_livros" json:"quant_livros"`
	ClienteUsuarioCPF  string    `bson:"cliente_usuario_cpf" json:"cliente_usuario_cpf"`
}
