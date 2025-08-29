package postgres

import (
	"context"
	"crud-biblioteca/model"
	"github.com/jackc/pgx/v5"
)

type UsuarioRepository struct {
	DB *pgx.Conn
}

func NewUsuarioRepository(db *pgx.Conn) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (r *UsuarioRepository) Create(ctx context.Context, usuario model.Usuario) error {
	query := `INSERT INTO "Projeto Logico".Usuario (cpf, data_nascimento, sobrenome, primeiro_nome) 
	          VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(ctx, query, usuario.CPF, usuario.DataNascimento, usuario.Sobrenome, usuario.PrimeiroNome)
	return err
}

func (r *UsuarioRepository) GetByCPF(ctx context.Context, cpf string) (*model.Usuario, error) {
	query := `SELECT cpf, data_nascimento, sobrenome, primeiro_nome 
	          FROM "Projeto Logico".Usuario WHERE cpf = $1`
	row := r.DB.QueryRow(ctx, query, cpf)
	var u model.Usuario
	err := row.Scan(&u.CPF, &u.DataNascimento, &u.Sobrenome, &u.PrimeiroNome)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UsuarioRepository) GetAll(ctx context.Context) ([]model.Usuario, error) {
	query := `SELECT cpf, data_nascimento, sobrenome, primeiro_nome 
	          FROM "Projeto Logico".Usuario ORDER BY cpf`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []model.Usuario
	for rows.Next() {
		var usuario model.Usuario
		if err := rows.Scan(&usuario.CPF, &usuario.DataNascimento, &usuario.Sobrenome, &usuario.PrimeiroNome); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, rows.Err()
}

func (r *UsuarioRepository) Update(ctx context.Context, usuario model.Usuario) error {
	query := `UPDATE "Projeto Logico".Usuario 
	          SET data_nascimento = $1, sobrenome = $2, primeiro_nome = $3 
			  WHERE cpf = $4`
	_, err := r.DB.Exec(ctx, query, usuario.DataNascimento, usuario.Sobrenome, usuario.PrimeiroNome, usuario.CPF)
	return err
}

func (r *UsuarioRepository) Delete(ctx context.Context, cpf string) error {
	query := `DELETE FROM "Projeto Logico".Usuario WHERE cpf = $1`
	_, err := r.DB.Exec(ctx, query, cpf)
	return err
}
