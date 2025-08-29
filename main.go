package main

import (
	"context"
	"crud-biblioteca/database"
	"crud-biblioteca/handlers"
	"crud-biblioteca/repository"
	mongoRepo "crud-biblioteca/repository/mongo"
	postgresRepo "crud-biblioteca/repository/postgres"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis do arquivo .env automaticamente
	_ = godotenv.Load()

	ctx := context.Background()

	// Inicializa o Gin
	router := gin.Default()

	// Adiciona CORS para permitir requisições do frontend
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Conecta aos bancos de dados
		pgConn, err := database.ConnectPostgres()
		if err != nil {
		log.Fatal("Erro ao conectar ao PostgreSQL:", err)
		}
		defer pgConn.Close(ctx)

		mongoClient, err := database.ConnectMongoDB()
		if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB:", err)
		}
		defer mongoClient.Disconnect(ctx)

	// Inicializa repositórios PostgreSQL
	pgUserRepo := postgresRepo.NewUsuarioRepository(pgConn)
	pgLivroRepo := postgresRepo.NewLivroRepository(pgConn)
	pgAutorRepo := postgresRepo.NewAutorRepository(pgConn)
	pgEmprestimoRepo := postgresRepo.NewEmprestimoRepository(pgConn)

	// Inicializa repositórios MongoDB
	db := mongoClient.Database("bibliotecaDB")
	mongoUserRepo := mongoRepo.NewUsuarioRepository(db)
	mongoLivroRepo := mongoRepo.NewLivroRepository(db)
	mongoAutorRepo := mongoRepo.NewAutorRepository(db)
	mongoEmprestimoRepo := mongoRepo.NewEmprestimoRepository(db)

	// ===== ROTAS PARA POSTGRESQL =====

	// --- Rotas para Usuários (PostgreSQL) ---
	router.GET("/postgres/usuarios", func(c *gin.Context) {
		handlers.ListarUsuariosPostgres(c, pgUserRepo)
	})
	router.POST("/postgres/usuarios", func(c *gin.Context) {
		handlers.CriarUsuarioPostgres(c, pgUserRepo)
	})
	router.PUT("/postgres/usuarios/:cpf", func(c *gin.Context) {
		handlers.AtualizarUsuarioPostgres(c, pgUserRepo)
	})
	router.DELETE("/postgres/usuarios/:cpf", func(c *gin.Context) {
		handlers.DeletarUsuarioPostgres(c, pgUserRepo)
	})

	// --- Rotas para Autores (PostgreSQL) ---
	router.GET("/postgres/autores", func(c *gin.Context) {
		handlers.ListarAutoresPostgres(c, pgAutorRepo)
	})
	router.POST("/postgres/autores", func(c *gin.Context) {
		handlers.CriarAutorPostgres(c, pgAutorRepo)
	})
	router.PUT("/postgres/autores/:id", func(c *gin.Context) {
		handlers.AtualizarAutorPostgres(c, pgAutorRepo)
	})
	router.DELETE("/postgres/autores/:id", func(c *gin.Context) {
		handlers.DeletarAutorPostgres(c, pgAutorRepo)
	})

	// --- Rotas para Livros (PostgreSQL) ---
	router.GET("/postgres/livros", func(c *gin.Context) {
		handlers.ListarLivrosPostgres(c, pgLivroRepo)
	})
	router.POST("/postgres/livros", func(c *gin.Context) {
		handlers.CriarLivroPostgres(c, pgLivroRepo)
	})
	router.PUT("/postgres/livros/:isbn", func(c *gin.Context) {
		handlers.AtualizarLivroPostgres(c, pgLivroRepo)
	})
	router.DELETE("/postgres/livros/:isbn", func(c *gin.Context) {
		handlers.DeletarLivroPostgres(c, pgLivroRepo)
	})

	// --- Rotas para Empréstimos (PostgreSQL) ---
	router.GET("/postgres/emprestimos", func(c *gin.Context) {
		handlers.ListarEmprestimosPostgres(c, pgEmprestimoRepo)
	})
	router.POST("/postgres/emprestimos", func(c *gin.Context) {
		handlers.CriarEmprestimoPostgres(c, pgEmprestimoRepo)
	})
	router.PUT("/postgres/emprestimos/:id", func(c *gin.Context) {
		handlers.AtualizarEmprestimoPostgres(c, pgEmprestimoRepo)
	})
	router.DELETE("/postgres/emprestimos/:id", func(c *gin.Context) {
		handlers.DeletarEmprestimoPostgres(c, pgEmprestimoRepo)
	})

	// ===== ROTAS PARA MONGODB =====

	// --- Rotas para Usuários (MongoDB) ---
	router.GET("/mongo/usuarios", func(c *gin.Context) {
		handlers.ListarUsuariosMongo(c, mongoUserRepo)
	})
	router.POST("/mongo/usuarios", func(c *gin.Context) {
		handlers.CriarUsuarioMongo(c, mongoUserRepo)
	})
	router.PUT("/mongo/usuarios/:id", func(c *gin.Context) {
		handlers.AtualizarUsuarioMongo(c, mongoUserRepo)
	})
	router.DELETE("/mongo/usuarios/:id", func(c *gin.Context) {
		handlers.DeletarUsuarioMongo(c, mongoUserRepo)
	})

	// --- Rotas para Autores (MongoDB) ---
	router.GET("/mongo/autores", func(c *gin.Context) {
		handlers.ListarAutoresMongo(c, mongoAutorRepo)
	})
	router.POST("/mongo/autores", func(c *gin.Context) {
		handlers.CriarAutorMongo(c, mongoAutorRepo)
	})
	router.PUT("/mongo/autores/:id", func(c *gin.Context) {
		handlers.AtualizarAutorMongo(c, mongoAutorRepo)
	})
	router.DELETE("/mongo/autores/:id", func(c *gin.Context) {
		handlers.DeletarAutorMongo(c, mongoAutorRepo)
	})

	// --- Rotas para Livros (MongoDB) ---
	router.GET("/mongo/livros", func(c *gin.Context) {
		handlers.ListarLivrosMongo(c, mongoLivroRepo)
	})
	router.POST("/mongo/livros", func(c *gin.Context) {
		handlers.CriarLivroMongo(c, mongoLivroRepo)
	})
	router.PUT("/mongo/livros/:id", func(c *gin.Context) {
		handlers.AtualizarLivroMongo(c, mongoLivroRepo)
	})
	router.DELETE("/mongo/livros/:id", func(c *gin.Context) {
		handlers.DeletarLivroMongo(c, mongoLivroRepo)
	})

	// --- Rotas para Empréstimos (MongoDB) ---
	router.GET("/mongo/emprestimos", func(c *gin.Context) {
		handlers.ListarEmprestimosMongo(c, mongoEmprestimoRepo)
	})
	router.POST("/mongo/emprestimos", func(c *gin.Context) {
		handlers.CriarEmprestimoMongo(c, mongoEmprestimoRepo)
	})
	router.PUT("/mongo/emprestimos/:id", func(c *gin.Context) {
		handlers.AtualizarEmprestimoMongo(c, mongoEmprestimoRepo)
	})
	router.DELETE("/mongo/emprestimos/:id", func(c *gin.Context) {
		handlers.DeletarEmprestimoMongo(c, mongoEmprestimoRepo)
	})

	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"message": "API da Biblioteca funcionando!",
			"endpoints": gin.H{
				"postgres": gin.H{
					"usuarios": "/postgres/usuarios",
					"autores": "/postgres/autores", 
					"livros": "/postgres/livros",
					"emprestimos": "/postgres/emprestimos",
				},
				"mongo": gin.H{
					"usuarios": "/mongo/usuarios",
					"autores": "/mongo/autores",
					"livros": "/mongo/livros", 
					"emprestimos": "/mongo/emprestimos",
				},
			},
		})
	})

	// Obtém a porta do ambiente ou usa 8080 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor iniciado na porta %s", port)
	log.Printf("📚 API da Biblioteca disponível em: http://localhost:%s", port)
	log.Printf("🏥 Health check: http://localhost:%s/health", port)
	log.Printf("📖 Documentação das rotas:")
	log.Printf("   PostgreSQL: http://localhost:%s/postgres/*", port)
	log.Printf("   MongoDB: http://localhost:%s/mongo/*", port)

	// Inicia o servidor
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
