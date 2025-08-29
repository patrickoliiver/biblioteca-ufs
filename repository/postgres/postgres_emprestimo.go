package postgres

import (
	"context"
	"crud-biblioteca/model"

	"github.com/jackc/pgx/v5"
)

type EmprestimoRepository struct {
	DB *pgx.Conn
}

func NewEmprestimoRepository(db *pgx.Conn) *EmprestimoRepository {
	return &EmprestimoRepository{DB: db}
}

func (r *EmprestimoRepository) Create(ctx context.Context, emprestimo model.Emprestimo) error {
	query := `INSERT INTO "Projeto Logico".Emprestimo (id, data_emprestimo, status, quant_livros, cliente_usuario_cpf) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(ctx, query, emprestimo.ID, emprestimo.DataEmprestimo, emprestimo.Status, emprestimo.QuantLivros, emprestimo.ClienteUsuarioCPF)
	return err
}

func (r *EmprestimoRepository) GetByID(ctx context.Context, id int) (*model.Emprestimo, error) {
	query := `SELECT id, data_emprestimo, status, quant_livros, cliente_usuario_cpf FROM "Projeto Logico".Emprestimo WHERE id = $1`
	row := r.DB.QueryRow(ctx, query, id)
	var e model.Emprestimo
	err := row.Scan(&e.ID, &e.DataEmprestimo, &e.Status, &e.QuantLivros, &e.ClienteUsuarioCPF)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *EmprestimoRepository) GetAll(ctx context.Context) ([]model.Emprestimo, error) {
	query := `SELECT id, data_emprestimo, status, quant_livros, cliente_usuario_cpf FROM "Projeto Logico".Emprestimo ORDER BY id`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emprestimos []model.Emprestimo
	for rows.Next() {
		var emprestimo model.Emprestimo
		if err := rows.Scan(&emprestimo.ID, &emprestimo.DataEmprestimo, &emprestimo.Status, &emprestimo.QuantLivros, &emprestimo.ClienteUsuarioCPF); err != nil {
			return nil, err
		}
		emprestimos = append(emprestimos, emprestimo)
	}
	return emprestimos, rows.Err()
}

func (r *EmprestimoRepository) Update(ctx context.Context, emprestimo model.Emprestimo) error {
	query := `UPDATE "Projeto Logico".Emprestimo SET data_emprestimo = $1, status = $2, quant_livros = $3, cliente_usuario_cpf = $4 WHERE id = $5`
	_, err := r.DB.Exec(ctx, query, emprestimo.DataEmprestimo, emprestimo.Status, emprestimo.QuantLivros, emprestimo.ClienteUsuarioCPF, emprestimo.ID)
	return err
}

func (r *EmprestimoRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM "Projeto Logico".Emprestimo WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
