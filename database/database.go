package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectPostgres() (*pgx.Conn, error) {
	// string de conexão agora vem da variável de ambiente
	connStr := os.Getenv("POSTGRES_CONN")
	if connStr == "" {
		fmt.Fprintf(os.Stderr, "Variável de ambiente POSTGRES_CONN não definida.\n")
		return nil, fmt.Errorf("variável de ambiente POSTGRES_CONN não definida")
	}
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Não foi possível conectar ao PostgreSQL: %v\n", err)
		return nil, err
	}
	fmt.Println("Conectado ao PostgreSQL com sucesso!")
	return conn, nil
}

func ConnectMongoDB() (*mongo.Client, error) {
	// substituir pela sua string de conexão do mongodb
	connStr := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Não foi possível conectar ao MongoDB: %v\n", err)
		return nil, err
	}

	// ping para verificar a conexão
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Conectado ao MongoDB com sucesso!")
	return client, nil
}
