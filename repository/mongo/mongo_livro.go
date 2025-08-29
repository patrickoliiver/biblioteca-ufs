package mongo

import (
	"context"
	"crud-biblioteca/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LivroRepository struct {
	Collection *mongo.Collection
}

func NewLivroRepository(db *mongo.Database) *LivroRepository {
	return &LivroRepository{Collection: db.Collection("livros")}
}

func (r *LivroRepository) Create(ctx context.Context, livro model.Livro) error {
	_, err := r.Collection.InsertOne(ctx, livro)
	return err
}

func (r *LivroRepository) GetByISBN(ctx context.Context, isbn string) (*model.Livro, error) {
	var livro model.Livro
	err := r.Collection.FindOne(ctx, bson.M{"_id": isbn}).Decode(&livro)
	return &livro, err
}

func (r *LivroRepository) GetAll(ctx context.Context) ([]model.Livro, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var livros []model.Livro
	if err = cursor.All(ctx, &livros); err != nil {
		return nil, err
	}
	return livros, nil
}

func (r *LivroRepository) Update(ctx context.Context, livro model.Livro) error {
	filter := bson.M{"_id": livro.ISBN}
	update := bson.M{"$set": bson.M{
		"titulo": livro.Titulo,
		"edicao": livro.Edicao,
	}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *LivroRepository) Delete(ctx context.Context, isbn string) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": isbn})
	return err
}

// CRUD do relacionamento embutido
func (r *LivroRepository) AddAutor(ctx context.Context, isbn string, autor model.Autor) error {
	filter := bson.M{"_id": isbn}
	update := bson.M{"$push": bson.M{"autores": autor}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *LivroRepository) RemoveAutor(ctx context.Context, isbn string, autorID int) error {
	filter := bson.M{"_id": isbn}
	update := bson.M{"$pull": bson.M{"autores": bson.M{"_id": autorID}}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}
