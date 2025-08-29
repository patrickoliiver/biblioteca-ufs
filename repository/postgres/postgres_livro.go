package postgres

import (
	"context"
	"crud-biblioteca/model"
	"github.com/jackc/pgx/v5"
)

type LivroRepository struct {
	DB *pgx.Conn
}

func NewLivroRepository(db *pgx.Conn) *LivroRepository {
	return &LivroRepository{DB: db}
}

func (r *LivroRepository) Create(ctx context.Context, livro model.Livro) error {
	query := `INSERT INTO "Projeto Logico".Livro (isbn, titulo, edicao, num_paginas, editora_cnpj, funcionario_matricula)
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.Exec(ctx, query, livro.ISBN, livro.Titulo, livro.Edicao, livro.NumPaginas, livro.EditoraCNPJ, livro.FuncionarioMatricula)
	return err
}

func (r *LivroRepository) GetByISBN(ctx context.Context, isbn string) (*model.Livro, error) {
	query := `SELECT isbn, titulo, edicao, num_paginas, editora_cnpj, funcionario_matricula FROM "Projeto Logico".Livro WHERE isbn = $1`
	row := r.DB.QueryRow(ctx, query, isbn)
	var l model.Livro
	err := row.Scan(&l.ISBN, &l.Titulo, &l.Edicao, &l.NumPaginas, &l.EditoraCNPJ, &l.FuncionarioMatricula)
	return &l, err
}

func (r *LivroRepository) GetAll(ctx context.Context) ([]model.Livro, error) {
	query := `SELECT isbn, titulo, edicao, num_paginas, editora_cnpj, funcionario_matricula FROM "Projeto Logico".Livro ORDER BY isbn`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var livros []model.Livro
	for rows.Next() {
		var livro model.Livro
		if err := rows.Scan(&livro.ISBN, &livro.Titulo, &livro.Edicao, &livro.NumPaginas, &livro.EditoraCNPJ, &livro.FuncionarioMatricula); err != nil {
			return nil, err
		}
		livros = append(livros, livro)
	}
	return livros, rows.Err()
}

func (r *LivroRepository) Update(ctx context.Context, livro model.Livro) error {
	query := `UPDATE "Projeto Logico".Livro SET titulo = $1, edicao = $2 WHERE isbn = $3`
	_, err := r.DB.Exec(ctx, query, livro.Titulo, livro.Edicao, livro.ISBN)
	return err
}

func (r *LivroRepository) Delete(ctx context.Context, isbn string) error {
	query := `DELETE FROM "Projeto Logico".Livro WHERE isbn = $1`
	_, err := r.DB.Exec(ctx, query, isbn)
	return err
}

// implementação do relacionamento para postgres (tabela Escreve)
func (r *LivroRepository) AddAutor(ctx context.Context, isbn string, autor model.Autor) error {
	query := `INSERT INTO "Projeto Logico".Escreve (livro_isbn, autor_id) VALUES ($1, $2)`
	_, err := r.DB.Exec(ctx, query, isbn, autor.ID)
	return err
}

func (r *LivroRepository) RemoveAutor(ctx context.Context, isbn string, autorID int) error {
	query := `DELETE FROM "Projeto Logico".Escreve WHERE livro_isbn = $1 AND autor_id = $2`
	_, err := r.DB.Exec(ctx, query, isbn, autorID)
	return err
}
