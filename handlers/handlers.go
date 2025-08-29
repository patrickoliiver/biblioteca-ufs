package handlers

import (
	"crud-biblioteca/model"
	"crud-biblioteca/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ===== HANDLERS PARA USUÁRIOS =====

// ListarUsuariosPostgres - GET /postgres/usuarios
func ListarUsuariosPostgres(c *gin.Context, userRepo repository.UsuarioRepository) {
	usuarios, err := userRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar usuários: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// CriarUsuarioPostgres - POST /postgres/usuarios
func CriarUsuarioPostgres(c *gin.Context, userRepo repository.UsuarioRepository) {
	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := userRepo.Create(c.Request.Context(), usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso", "usuario": usuario})
}

// AtualizarUsuarioPostgres - PUT /postgres/usuarios/:cpf
func AtualizarUsuarioPostgres(c *gin.Context, userRepo repository.UsuarioRepository) {
	cpf := c.Param("cpf")
	
	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	usuario.CPF = cpf // Garante que o CPF da URL seja usado

	if err := userRepo.Update(c.Request.Context(), usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso", "usuario": usuario})
}

// DeletarUsuarioPostgres - DELETE /postgres/usuarios/:cpf
func DeletarUsuarioPostgres(c *gin.Context, userRepo repository.UsuarioRepository) {
	cpf := c.Param("cpf")

	if err := userRepo.Delete(c.Request.Context(), cpf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}

// ListarUsuariosMongo - GET /mongo/usuarios
func ListarUsuariosMongo(c *gin.Context, userRepo repository.UsuarioRepository) {
	usuarios, err := userRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar usuários: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// CriarUsuarioMongo - POST /mongo/usuarios
func CriarUsuarioMongo(c *gin.Context, userRepo repository.UsuarioRepository) {
	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := userRepo.Create(c.Request.Context(), usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso", "usuario": usuario})
}

// AtualizarUsuarioMongo - PUT /mongo/usuarios/:id
func AtualizarUsuarioMongo(c *gin.Context, userRepo repository.UsuarioRepository) {
	id := c.Param("id")
	
	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	usuario.CPF = id // No MongoDB, o ID é o CPF

	if err := userRepo.Update(c.Request.Context(), usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso", "usuario": usuario})
}

// DeletarUsuarioMongo - DELETE /mongo/usuarios/:id
func DeletarUsuarioMongo(c *gin.Context, userRepo repository.UsuarioRepository) {
	id := c.Param("id")

	if err := userRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}

// ===== HANDLERS PARA AUTORES =====

// ListarAutoresPostgres - GET /postgres/autores
func ListarAutoresPostgres(c *gin.Context, autorRepo repository.AutorRepository) {
	autores, err := autorRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar autores: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, autores)
}

// CriarAutorPostgres - POST /postgres/autores
func CriarAutorPostgres(c *gin.Context, autorRepo repository.AutorRepository) {
	var autor model.Autor
	if err := c.ShouldBindJSON(&autor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := autorRepo.Create(c.Request.Context(), autor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar autor: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Autor criado com sucesso", "autor": autor})
}

// AtualizarAutorPostgres - PUT /postgres/autores/:id
func AtualizarAutorPostgres(c *gin.Context, autorRepo repository.AutorRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	var autor model.Autor
	if err := c.ShouldBindJSON(&autor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	autor.ID = id

	if err := autorRepo.Update(c.Request.Context(), autor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar autor: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Autor atualizado com sucesso", "autor": autor})
}

// DeletarAutorPostgres - DELETE /postgres/autores/:id
func DeletarAutorPostgres(c *gin.Context, autorRepo repository.AutorRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := autorRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar autor: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Autor deletado com sucesso"})
}

// ListarAutoresMongo - GET /mongo/autores
func ListarAutoresMongo(c *gin.Context, autorRepo repository.AutorRepository) {
	autores, err := autorRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar autores: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, autores)
}

// CriarAutorMongo - POST /mongo/autores
func CriarAutorMongo(c *gin.Context, autorRepo repository.AutorRepository) {
	var autor model.Autor
	if err := c.ShouldBindJSON(&autor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := autorRepo.Create(c.Request.Context(), autor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar autor: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Autor criado com sucesso", "autor": autor})
}

// AtualizarAutorMongo - PUT /mongo/autores/:id
func AtualizarAutorMongo(c *gin.Context, autorRepo repository.AutorRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	var autor model.Autor
	if err := c.ShouldBindJSON(&autor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	autor.ID = id

	if err := autorRepo.Update(c.Request.Context(), autor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar autor: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Autor atualizado com sucesso", "autor": autor})
}

// DeletarAutorMongo - DELETE /mongo/autores/:id
func DeletarAutorMongo(c *gin.Context, autorRepo repository.AutorRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := autorRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar autor: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Autor deletado com sucesso"})
}

// ===== HANDLERS PARA LIVROS =====

// ListarLivrosPostgres - GET /postgres/livros
func ListarLivrosPostgres(c *gin.Context, livroRepo repository.LivroRepository) {
	livros, err := livroRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar livros: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, livros)
}

// CriarLivroPostgres - POST /postgres/livros
func CriarLivroPostgres(c *gin.Context, livroRepo repository.LivroRepository) {
	var livro model.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := livroRepo.Create(c.Request.Context(), livro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar livro: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Livro criado com sucesso", "livro": livro})
}

// AtualizarLivroPostgres - PUT /postgres/livros/:isbn
func AtualizarLivroPostgres(c *gin.Context, livroRepo repository.LivroRepository) {
	isbn := c.Param("isbn")
	
	var livro model.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	livro.ISBN = isbn

	if err := livroRepo.Update(c.Request.Context(), livro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar livro: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro atualizado com sucesso", "livro": livro})
}

// DeletarLivroPostgres - DELETE /postgres/livros/:isbn
func DeletarLivroPostgres(c *gin.Context, livroRepo repository.LivroRepository) {
	isbn := c.Param("isbn")

	if err := livroRepo.Delete(c.Request.Context(), isbn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar livro: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso"})
}

// ListarLivrosMongo - GET /mongo/livros
func ListarLivrosMongo(c *gin.Context, livroRepo repository.LivroRepository) {
	livros, err := livroRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar livros: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, livros)
}

// CriarLivroMongo - POST /mongo/livros
func CriarLivroMongo(c *gin.Context, livroRepo repository.LivroRepository) {
	var livro model.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := livroRepo.Create(c.Request.Context(), livro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar livro: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Livro criado com sucesso", "livro": livro})
}

// AtualizarLivroMongo - PUT /mongo/livros/:id
func AtualizarLivroMongo(c *gin.Context, livroRepo repository.LivroRepository) {
	id := c.Param("id")
	
	var livro model.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	livro.ISBN = id // No MongoDB, o ID é o ISBN

	if err := livroRepo.Update(c.Request.Context(), livro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar livro: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro atualizado com sucesso", "livro": livro})
}

// DeletarLivroMongo - DELETE /mongo/livros/:id
func DeletarLivroMongo(c *gin.Context, livroRepo repository.LivroRepository) {
	id := c.Param("id")

	if err := livroRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar livro: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso"})
}

// ===== HANDLERS PARA EMPRÉSTIMOS =====

// ListarEmprestimosPostgres - GET /postgres/emprestimos
func ListarEmprestimosPostgres(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	emprestimos, err := emprestimoRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar empréstimos: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, emprestimos)
}

// CriarEmprestimoPostgres - POST /postgres/emprestimos
func CriarEmprestimoPostgres(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	var emprestimo model.Emprestimo
	if err := c.ShouldBindJSON(&emprestimo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Define a data do empréstimo como agora se não fornecida
	if emprestimo.DataEmprestimo.IsZero() {
		emprestimo.DataEmprestimo = time.Now()
	}

	if err := emprestimoRepo.Create(c.Request.Context(), emprestimo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar empréstimo: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Empréstimo criado com sucesso", "emprestimo": emprestimo})
}

// AtualizarEmprestimoPostgres - PUT /postgres/emprestimos/:id
func AtualizarEmprestimoPostgres(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	var emprestimo model.Emprestimo
	if err := c.ShouldBindJSON(&emprestimo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	emprestimo.ID = id

	if err := emprestimoRepo.Update(c.Request.Context(), emprestimo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar empréstimo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empréstimo atualizado com sucesso", "emprestimo": emprestimo})
}

// DeletarEmprestimoPostgres - DELETE /postgres/emprestimos/:id
func DeletarEmprestimoPostgres(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := emprestimoRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar empréstimo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empréstimo deletado com sucesso"})
}

// ListarEmprestimosMongo - GET /mongo/emprestimos
func ListarEmprestimosMongo(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	emprestimos, err := emprestimoRepo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar empréstimos: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, emprestimos)
}

// CriarEmprestimoMongo - POST /mongo/emprestimos
func CriarEmprestimoMongo(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	var emprestimo model.Emprestimo
	if err := c.ShouldBindJSON(&emprestimo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Define a data do empréstimo como agora se não fornecida
	if emprestimo.DataEmprestimo.IsZero() {
		emprestimo.DataEmprestimo = time.Now()
	}

	if err := emprestimoRepo.Create(c.Request.Context(), emprestimo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar empréstimo: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Empréstimo criado com sucesso", "emprestimo": emprestimo})
}

// AtualizarEmprestimoMongo - PUT /mongo/emprestimos/:id
func AtualizarEmprestimoMongo(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	var emprestimo model.Emprestimo
	if err := c.ShouldBindJSON(&emprestimo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	emprestimo.ID = id

	if err := emprestimoRepo.Update(c.Request.Context(), emprestimo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar empréstimo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empréstimo atualizado com sucesso", "emprestimo": emprestimo})
}

// DeletarEmprestimoMongo - DELETE /mongo/emprestimos/:id
func DeletarEmprestimoMongo(c *gin.Context, emprestimoRepo repository.EmprestimoRepository) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := emprestimoRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar empréstimo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empréstimo deletado com sucesso"})
}
