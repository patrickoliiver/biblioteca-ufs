package repository

import (
	"context"
	"crud-biblioteca/model"
)

type UsuarioRepository interface {
	Create(ctx context.Context, usuario model.Usuario) error
	GetByCPF(ctx context.Context, cpf string) (*model.Usuario, error)
	GetAll(ctx context.Context) ([]model.Usuario, error)
	Update(ctx context.Context, usuario model.Usuario) error
	Delete(ctx context.Context, cpf string) error
}

type AutorRepository interface {
	Create(ctx context.Context, autor model.Autor) error
	GetByID(ctx context.Context, id int) (*model.Autor, error)
	GetAll(ctx context.Context) ([]model.Autor, error)
	Update(ctx context.Context, autor model.Autor) error
	Delete(ctx context.Context, id int) error
}

type LivroRepository interface {
	Create(ctx context.Context, livro model.Livro) error
	GetByISBN(ctx context.Context, isbn string) (*model.Livro, error)
	GetAll(ctx context.Context) ([]model.Livro, error)
	Update(ctx context.Context, livro model.Livro) error
	Delete(ctx context.Context, isbn string) error
	
	AddAutor(ctx context.Context, isbn string, autor model.Autor) error
	RemoveAutor(ctx context.Context, isbn string, autorID int) error
}

type EmprestimoRepository interface {
	Create(ctx context.Context, emprestimo model.Emprestimo) error
	GetByID(ctx context.Context, id int) (*model.Emprestimo, error)
	GetAll(ctx context.Context) ([]model.Emprestimo, error)
	Update(ctx context.Context, emprestimo model.Emprestimo) error
	Delete(ctx context.Context, id int) error
}
