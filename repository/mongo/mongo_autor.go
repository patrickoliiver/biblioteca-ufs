package mongo

import (
	"context"
	"crud-biblioteca/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AutorRepository struct {
	Collection *mongo.Collection
}

func NewAutorRepository(db *mongo.Database) *AutorRepository {
	return &AutorRepository{Collection: db.Collection("autores")}
}

func (r *AutorRepository) Create(ctx context.Context, autor model.Autor) error {
	_, err := r.Collection.InsertOne(ctx, autor)
	return err
}

func (r *AutorRepository) GetByID(ctx context.Context, id int) (*model.Autor, error) {
	var autor model.Autor
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&autor)
	return &autor, err
}
func (r *AutorRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
