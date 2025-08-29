package main

import (
	"bufio"
	"context"
	"crud-biblioteca/database"
	"crud-biblioteca/model"
	"crud-biblioteca/repository"
	mongoRepo "crud-biblioteca/repository/mongo"
	postgresRepo "crud-biblioteca/repository/postgres"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis do arquivo .env automaticamente
	_ = godotenv.Load()

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	// escolha do banco de dados
	var userRepo repository.UsuarioRepository
	var livroRepo repository.LivroRepository
	var autorRepo repository.AutorRepository
	var emprestimoRepo repository.EmprestimoRepository

	fmt.Println("Bem-vindo ao sistema de gerenciamento da biblioteca!")
	fmt.Println("Qual banco de dados você deseja usar?")
	fmt.Println("1: PostgreSQL")
	fmt.Println("2: MongoDB")
	fmt.Print("Escolha uma opção: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		log.Println("Conectando ao PostgreSQL...")
		pgConn, err := database.ConnectPostgres()
		if err != nil {
			os.Exit(1)
		}
		defer pgConn.Close(ctx)
		userRepo = postgresRepo.NewUsuarioRepository(pgConn)
		livroRepo = postgresRepo.NewLivroRepository(pgConn)
		autorRepo = postgresRepo.NewAutorRepository(pgConn)
		emprestimoRepo = postgresRepo.NewEmprestimoRepository(pgConn)
	case "2":
		log.Println("Conectando ao MongoDB...")
		mongoClient, err := database.ConnectMongoDB()
		if err != nil {
			os.Exit(1)
		}
		defer mongoClient.Disconnect(ctx)
		db := mongoClient.Database("bibliotecaDB")
		userRepo = mongoRepo.NewUsuarioRepository(db)
		livroRepo = mongoRepo.NewLivroRepository(db)
		autorRepo = mongoRepo.NewAutorRepository(db)
		emprestimoRepo = mongoRepo.NewEmprestimoRepository(db)
	default:
		log.Fatal("Opção inválida. Saindo.")
		return
	}

	// menu principal
	for {
		fmt.Println("\n===== MENU DE OPERAÇÕES =====")
		fmt.Println("--- Entidade: Usuário ---")
		fmt.Println("1: Criar Usuário")
		fmt.Println("2: Ler Usuário por CPF")
		fmt.Println("3: Atualizar Usuário (completo)")
		fmt.Println("4: Deletar Usuário")
		fmt.Println("--- Entidade: Livro e Relacionamento ---")
		fmt.Println("5: Criar Livro (com dados de teste fixos)")
		fmt.Println("6: Ler Livro por ISBN")
		fmt.Println("7: Deletar Livro")
		fmt.Println("8: Adicionar Autor a um Livro (Criar Relacionamento)")
		fmt.Println("9: Remover Autor de um Livro (Deletar Relacionamento E Autor)")
		fmt.Println("--- Entidade: Empréstimo ---")
		fmt.Println("10: Criar Empréstimo")
		fmt.Println("11: Ler Empréstimo por ID")
		fmt.Println("12: Atualizar Empréstimo")
		fmt.Println("13: Deletar Empréstimo")
		fmt.Println("-------------------------------")
		fmt.Println("0: Sair")
		fmt.Print("Escolha uma opção: ")

		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		switch op {
		case "1":
			handleCreateUsuario(ctx, userRepo, reader)
		case "2":
			handleReadUsuario(ctx, userRepo, reader)
		case "3":
			handleUpdateUsuario(ctx, userRepo, reader)
		case "4":
			handleDeleteUsuario(ctx, userRepo, reader)
		case "5":
			handleCreateLivro(ctx, livroRepo, reader)
		case "6":
			handleReadLivro(ctx, livroRepo, reader)
		case "7":
			handleDeleteLivro(ctx, livroRepo, reader)
		case "8":
			handleAddAutorRelacionamento(ctx, livroRepo, autorRepo, reader)
		case "9":
			handleRemoveAutorRelacionamento(ctx, livroRepo, autorRepo, reader)
		case "10":
			handleCreateEmprestimo(ctx, emprestimoRepo, reader)
		case "11":
			handleReadEmprestimo(ctx, emprestimoRepo, reader)
		case "12":
			handleUpdateEmprestimo(ctx, emprestimoRepo, reader)
		case "13":
			handleDeleteEmprestimo(ctx, emprestimoRepo, reader)
		case "0":
			log.Println("Saindo do sistema. Até logo!")
			return
		default:
			log.Println("Opção inválida. Tente novamente.")
		}
	}
}

// funções auxiliares
// CRUD de Empréstimo
func handleCreateEmprestimo(ctx context.Context, repo repository.EmprestimoRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ID do empréstimo (número inteiro): ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ERRO: ID inválido. %v\n", err)
		return
	}

	dataEmprestimo := time.Now()

	fmt.Print("Digite o status do empréstimo: ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	fmt.Print("Digite a quantidade de livros: ")
	quantLivrosStr, _ := reader.ReadString('\n')
	quantLivrosStr = strings.TrimSpace(quantLivrosStr)
	quantLivros, err := strconv.Atoi(quantLivrosStr)
	if err != nil {
		log.Printf("ERRO: Quantidade de livros inválida. %v\n", err)
		return
	}

	fmt.Print("Digite o CPF do cliente/usuário: ")
	clienteCPF, _ := reader.ReadString('\n')
	clienteCPF = strings.TrimSpace(clienteCPF)

	novoEmprestimo := model.Emprestimo{
		ID:                id,
		DataEmprestimo:    dataEmprestimo,
		Status:            status,
		QuantLivros:       quantLivros,
		ClienteUsuarioCPF: clienteCPF,
	}

	if err := repo.Create(ctx, novoEmprestimo); err != nil {
		log.Printf("ERRO: Não foi possível criar o empréstimo. %v\n", err)
	} else {
		log.Println("SUCESSO: Empréstimo criado. Verifique o banco de dados.")
	}
}

func handleReadEmprestimo(ctx context.Context, repo repository.EmprestimoRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ID do empréstimo a ser lido (número inteiro): ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ERRO: ID inválido. %v\n", err)
		return
	}

	emprestimo, err := repo.GetByID(ctx, id)
	if err != nil {
		log.Printf("ERRO: Empréstimo com ID '%d' não encontrado. %v\n", id, err)
	} else {
		log.Printf("SUCESSO: Empréstimo encontrado: %+v\n", *emprestimo)
	}
}

func handleUpdateEmprestimo(ctx context.Context, repo repository.EmprestimoRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ID do empréstimo a ser atualizado (número inteiro): ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ERRO: ID inválido. %v\n", err)
		return
	}

	emprestimo, err := repo.GetByID(ctx, id)
	if err != nil {
		log.Printf("ERRO: Empréstimo com ID '%d' não encontrado para atualizar. %v\n", id, err)
		return
	}
	log.Printf("Atualizando empréstimo: %+v\n", *emprestimo)

	fmt.Printf("Digite o novo status (atual: %s): ", emprestimo.Status)
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)
	if status != "" {
		emprestimo.Status = status
	}

	fmt.Printf("Digite a nova quantidade de livros (atual: %d): ", emprestimo.QuantLivros)
	quantLivrosStr, _ := reader.ReadString('\n')
	quantLivrosStr = strings.TrimSpace(quantLivrosStr)
	if quantLivrosStr != "" {
		quantLivros, err := strconv.Atoi(quantLivrosStr)
		if err == nil {
			emprestimo.QuantLivros = quantLivros
		}
	}

	fmt.Printf("Digite o novo CPF do cliente/usuário (atual: %s): ", emprestimo.ClienteUsuarioCPF)
	clienteCPF, _ := reader.ReadString('\n')
	clienteCPF = strings.TrimSpace(clienteCPF)
	if clienteCPF != "" {
		emprestimo.ClienteUsuarioCPF = clienteCPF
	}

	// Atualiza data do empréstimo para agora
	emprestimo.DataEmprestimo = time.Now()

	if err := repo.Update(ctx, *emprestimo); err != nil {
		log.Printf("ERRO: Não foi possível atualizar o empréstimo. %v\n", err)
	} else {
		log.Println("SUCESSO: Empréstimo atualizado. Verifique o banco de dados.")
	}
}

func handleDeleteEmprestimo(ctx context.Context, repo repository.EmprestimoRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ID do empréstimo a ser deletado (número inteiro): ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("ERRO: ID inválido. %v\n", err)
		return
	}

	if err := repo.Delete(ctx, id); err != nil {
		log.Printf("ERRO: Não foi possível deletar o empréstimo. %v\n", err)
	} else {
		log.Println("SUCESSO: Empréstimo deletado. Verifique o banco de dados.")
	}
}

func handleCreateUsuario(ctx context.Context, repo repository.UsuarioRepository, reader *bufio.Reader) {
	fmt.Print("Digite o CPF: ")
	cpf, _ := reader.ReadString('\n')

	fmt.Print("Digite o Primeiro Nome: ")
	primeiroNome, _ := reader.ReadString('\n')

	fmt.Print("Digite o Sobrenome: ")
	sobrenome, _ := reader.ReadString('\n')

	fmt.Print("Digite a Data de Nascimento (AAAA-MM-DD): ")
	dataNascStr, _ := reader.ReadString('\n')
	dataNasc, err := time.Parse("2006-01-02", strings.TrimSpace(dataNascStr))
	if err != nil {
		log.Printf("Formato de data inválido: %v\n", err)
		return
	}

	novoUsuario := model.Usuario{
		CPF:            strings.TrimSpace(cpf),
		PrimeiroNome:   strings.TrimSpace(primeiroNome),
		Sobrenome:      strings.TrimSpace(sobrenome),
		DataNascimento: dataNasc,
	}

	if err := repo.Create(ctx, novoUsuario); err != nil {
		log.Printf("ERRO: Não foi possível criar o usuário. %v\n", err)
	} else {
		log.Println("SUCESSO: Usuário criado. Verifique o banco de dados.")
	}
}

func handleReadUsuario(ctx context.Context, repo repository.UsuarioRepository, reader *bufio.Reader) {
	fmt.Print("Digite o CPF do usuário a ser lido: ")
	cpf, _ := reader.ReadString('\n')
	cpf = strings.TrimSpace(cpf)

	usuario, err := repo.GetByCPF(ctx, cpf)
	if err != nil {
		log.Printf("ERRO: Usuário com CPF '%s' não encontrado. %v\n", cpf, err)
	} else {
		log.Printf("SUCESSO: Usuário encontrado: %+v\n", *usuario)
	}
}

func handleUpdateUsuario(ctx context.Context, repo repository.UsuarioRepository, reader *bufio.Reader) {
	fmt.Print("Digite o CPF do usuário a ser atualizado: ")
	cpf, _ := reader.ReadString('\n')
	cpf = strings.TrimSpace(cpf)

	// 1buscar o usuário existente para obter os dados atuais
	usuario, err := repo.GetByCPF(ctx, cpf)
	if err != nil {
		log.Printf("ERRO: Usuário com CPF '%s' não encontrado para atualizar.\n", cpf)
		return
	}
	log.Printf("Atualizando usuário: %+v\n", *usuario)
	log.Println("Deixe o campo em branco e pressione Enter para manter o valor atual.")

	// coletar novos dados para cada campo
	fmt.Printf("Digite o novo Primeiro Nome (atual: %s): ", usuario.PrimeiroNome)
	primeiroNome, _ := reader.ReadString('\n')
	primeiroNome = strings.TrimSpace(primeiroNome)
	if primeiroNome != "" { // Só atualiza se o usuário digitar algo
		usuario.PrimeiroNome = primeiroNome
	}

	fmt.Printf("Digite o novo Sobrenome (atual: %s): ", usuario.Sobrenome)
	sobrenome, _ := reader.ReadString('\n')
	sobrenome = strings.TrimSpace(sobrenome)
	if sobrenome != "" {
		usuario.Sobrenome = sobrenome
	}

	fmt.Printf("Digite a nova Data de Nascimento (AAAA-MM-DD) (atual: %s): ", usuario.DataNascimento.Format("2006-01-02"))
	dataNascStr, _ := reader.ReadString('\n')
	dataNascStr = strings.TrimSpace(dataNascStr)
	if dataNascStr != "" {
		dataNasc, err := time.Parse("2006-01-02", dataNascStr)
		if err != nil {
			log.Printf("ERRO: Formato de data inválido. A data de nascimento não foi alterada: %v\n", err)
		} else {
			usuario.DataNascimento = dataNasc
		}
	}

	// chamar o método de atualização do repositório com o objeto modificado
	if err := repo.Update(ctx, *usuario); err != nil {
		log.Printf("ERRO: Não foi possível atualizar o usuário. %v\n", err)
	} else {
		log.Println("SUCESSO: Usuário atualizado. Verifique o banco de dados.")
	}
}

func handleDeleteUsuario(ctx context.Context, repo repository.UsuarioRepository, reader *bufio.Reader) {
	fmt.Print("Digite o CPF do usuário a ser deletado: ")
	cpf, _ := reader.ReadString('\n')
	cpf = strings.TrimSpace(cpf)

	if err := repo.Delete(ctx, cpf); err != nil {
		log.Printf("ERRO: Não foi possível deletar o usuário. %v\n", err)
	} else {
		log.Println("SUCESSO: Usuário deletado. Verifique o banco de dados.")
	}
}

func handleCreateLivro(ctx context.Context, repo repository.LivroRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ISBN do livro: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	fmt.Print("Digite o Título: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	fmt.Print("Digite a Edição: ")
	edicao, _ := reader.ReadString('\n')
	edicao = strings.TrimSpace(edicao)

	fmt.Print("Digite o Número de Páginas: ")
	numPaginasStr, _ := reader.ReadString('\n')
	numPaginas, err := strconv.Atoi(strings.TrimSpace(numPaginasStr))
	if err != nil {
		log.Printf("ERRO: Número de páginas inválido: %v\n", err)
		return
	}

	const editoraCNPJFixo = "11222333000144"
	const funcMatriculaFixo = 100
	log.Printf("Usando valores fixos para teste: CNPJ da Editora=%s, Matrícula do Funcionário=%d\n", editoraCNPJFixo, funcMatriculaFixo)

	novoLivro := model.Livro{
		ISBN:                 isbn,
		Titulo:               titulo,
		Edicao:               edicao,
		NumPaginas:           numPaginas,
		EditoraCNPJ:          editoraCNPJFixo,   // valor fixo para teste
		FuncionarioMatricula: funcMatriculaFixo, // valor fixo para teste
		Autores:              []model.Autor{},
	}

	if err := repo.Create(ctx, novoLivro); err != nil {
		log.Printf("ERRO: Não foi possível criar o livro. %v\n", err)
	} else {
		log.Println("SUCESSO: Livro criado. Verifique o banco de dados.")
	}
}

func handleReadLivro(ctx context.Context, repo repository.LivroRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ISBN do livro a ser lido: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	livro, err := repo.GetByISBN(ctx, isbn)
	if err != nil {
		log.Printf("ERRO: Livro com ISBN '%s' não encontrado. %v\n", isbn, err)
	} else {
		log.Printf("SUCESSO: Livro encontrado: %+v\n", *livro)
	}
}

func handleDeleteLivro(ctx context.Context, repo repository.LivroRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ISBN do livro a ser deletado: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	if err := repo.Delete(ctx, isbn); err != nil {
		log.Printf("ERRO: Não foi possível deletar o livro. %v\n", err)
	} else {
		log.Println("SUCESSO: Livro deletado. Verifique o banco de dados.")
	}
}

func handleAddAutorRelacionamento(ctx context.Context, livroRepo repository.LivroRepository, autorRepo repository.AutorRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ISBN do livro para adicionar um autor: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	fmt.Print("Digite o ID do autor: ")
	autorIDStr, _ := reader.ReadString('\n')
	autorID, _ := strconv.Atoi(strings.TrimSpace(autorIDStr))

	fmt.Print("Digite o primeiro nome do autor: ")
	autorNome, _ := reader.ReadString('\n')

	fmt.Print("Digite o sobrenome do autor: ")
	autorSobrenome, _ := reader.ReadString('\n')

	autor := model.Autor{
		ID:           autorID,
		PrimeiroNome: strings.TrimSpace(autorNome),
		Sobrenome:    strings.TrimSpace(autorSobrenome),
	}

	if err := autorRepo.Create(ctx, autor); err != nil {
		// ignora o erro se o autor já existir
		log.Printf("AVISO: Autor pode já existir. Continuando... (%v)\n", err)
	}

	if err := livroRepo.AddAutor(ctx, isbn, autor); err != nil {
		log.Printf("ERRO: Não foi possível adicionar o relacionamento. %v\n", err)
	} else {
		log.Println("SUCESSO: Relacionamento criado. Verifique o banco de dados.")
	}
}

func handleRemoveAutorRelacionamento(ctx context.Context, livroRepo repository.LivroRepository, autorRepo repository.AutorRepository, reader *bufio.Reader) {
	fmt.Print("Digite o ISBN do livro para remover um autor: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	fmt.Print("Digite o ID do autor a ser removido: ")
	autorIDStr, _ := reader.ReadString('\n')
	autorID, _ := strconv.Atoi(strings.TrimSpace(autorIDStr))

	// remove o relacionamento do livro.
	if err := livroRepo.RemoveAutor(ctx, isbn, autorID); err != nil {
		log.Printf("ERRO: Não foi possível remover o relacionamento. %v\n", err)
		return // se não conseguir remover a relação, não deleta o autor.
	}
	log.Println("SUCESSO: Relacionamento removido. Verifique a tabela/coleção de relacionamento.")

	// deleta o autor da tabela/coleção principal 'Autor'.
	log.Printf("Deletando autor com ID %d da tabela principal 'Autor'...", autorID)
	if err := autorRepo.Delete(ctx, autorID); err != nil {
		log.Printf("ERRO: Não foi possível deletar o autor da tabela principal. %v\n", err)
	} else {
		log.Println("SUCESSO: Autor deletado da tabela principal. Verifique o banco de dados.")
	}
}
