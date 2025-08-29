package postgres

import (
	"context"
	"crud-biblioteca/model"
	"github.com/jackc/pgx/v5"
)

type AutorRepository struct {
	DB *pgx.Conn
}

func NewAutorRepository(db *pgx.Conn) *AutorRepository {
	return &AutorRepository{DB: db}
}

func (r *AutorRepository) Create(ctx context.Context, autor model.Autor) error {
	query := `INSERT INTO "Projeto Logico".Autor (id, primeiro_nome, sobrenome) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING`
	_, err := r.DB.Exec(ctx, query, autor.ID, autor.PrimeiroNome, autor.Sobrenome)
	return err
}

func (r *AutorRepository) GetByID(ctx context.Context, id int) (*model.Autor, error) {
	query := `SELECT id, primeiro_nome, sobrenome FROM "Projeto Logico".Autor WHERE id = $1`
	row := r.DB.QueryRow(ctx, query, id)
	var a model.Autor
	err := row.Scan(&a.ID, &a.PrimeiroNome, &a.Sobrenome)
	return &a, err
}

func (r *AutorRepository) GetAll(ctx context.Context) ([]model.Autor, error) {
	query := `SELECT id, primeiro_nome, sobrenome FROM "Projeto Logico".Autor ORDER BY id`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var autores []model.Autor
	for rows.Next() {
		var autor model.Autor
		if err := rows.Scan(&autor.ID, &autor.PrimeiroNome, &autor.Sobrenome); err != nil {
			return nil, err
		}
		autores = append(autores, autor)
	}
	return autores, rows.Err()
}

func (r *AutorRepository) Update(ctx context.Context, autor model.Autor) error {
	query := `UPDATE "Projeto Logico".Autor SET primeiro_nome = $2, sobrenome = $3 WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, autor.ID, autor.PrimeiroNome, autor.Sobrenome)
	return err
}

func (r *AutorRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM "Projeto Logico".Autor WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
