package mongo

import (
	"context"
	"crud-biblioteca/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioRepository struct {
	Collection *mongo.Collection
}

func NewUsuarioRepository(db *mongo.Database) *UsuarioRepository {
	return &UsuarioRepository{Collection: db.Collection("usuarios")}
}

func (r *UsuarioRepository) Create(ctx context.Context, usuario model.Usuario) error {
	_, err := r.Collection.InsertOne(ctx, usuario)
	return err
}

func (r *UsuarioRepository) GetByCPF(ctx context.Context, cpf string) (*model.Usuario, error) {
	var usuario model.Usuario
	err := r.Collection.FindOne(ctx, bson.M{"_id": cpf}).Decode(&usuario)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) GetAll(ctx context.Context) ([]model.Usuario, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var usuarios []model.Usuario
	if err = cursor.All(ctx, &usuarios); err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (r *UsuarioRepository) Update(ctx context.Context, usuario model.Usuario) error {
	filter := bson.M{"_id": usuario.CPF}
	update := bson.M{"$set": bson.M{
		"data_nascimento": usuario.DataNascimento,
		"sobrenome":       usuario.Sobrenome,
		"primeiro_nome":   usuario.PrimeiroNome,
	}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UsuarioRepository) Delete(ctx context.Context, cpf string) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": cpf})
	return err
}
