package mongo

import (
	"context"
	"crud-biblioteca/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmprestimoRepository struct {
	Collection *mongo.Collection
}

func NewEmprestimoRepository(db *mongo.Database) *EmprestimoRepository {
	return &EmprestimoRepository{Collection: db.Collection("emprestimos")}
}

func (r *EmprestimoRepository) Create(ctx context.Context, emprestimo model.Emprestimo) error {
	_, err := r.Collection.InsertOne(ctx, emprestimo)
	return err
}

func (r *EmprestimoRepository) GetByID(ctx context.Context, id int) (*model.Emprestimo, error) {
	var emprestimo model.Emprestimo
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&emprestimo)
	if err != nil {
		return nil, err
	}
	return &emprestimo, nil
}

func (r *EmprestimoRepository) GetAll(ctx context.Context) ([]model.Emprestimo, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var emprestimos []model.Emprestimo
	if err = cursor.All(ctx, &emprestimos); err != nil {
		return nil, err
	}
	return emprestimos, nil
}

func (r *EmprestimoRepository) Update(ctx context.Context, emprestimo model.Emprestimo) error {
	filter := bson.M{"_id": emprestimo.ID}
	update := bson.M{"$set": bson.M{
		"data_emprestimo":    emprestimo.DataEmprestimo,
		"status":              emprestimo.Status,
		"quant_livros":        emprestimo.QuantLivros,
		"cliente_usuario_cpf": emprestimo.ClienteUsuarioCPF,
	}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *EmprestimoRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
