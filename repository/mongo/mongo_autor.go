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

func (r *AutorRepository) GetAll(ctx context.Context) ([]model.Autor, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var autores []model.Autor
	if err = cursor.All(ctx, &autores); err != nil {
		return nil, err
	}
	return autores, nil
}

func (r *AutorRepository) Update(ctx context.Context, autor model.Autor) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": autor.ID}, bson.M{"$set": autor})
	return err
}

func (r *AutorRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
