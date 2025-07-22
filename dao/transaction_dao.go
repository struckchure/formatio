package dao

import (
	"context"
	"strings"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type ITransactionDao interface {
	CreateTransaction(args types.CreateTransactionArgs) (*db.TransactionModel, error)
	DeleteTransaction(args types.DeleteTransactionArgs) error
	GetTransaction(args types.GetTransactionArgs) (*db.TransactionModel, error)
	ListTransaction(args types.ListTransactionArgs) ([]db.TransactionModel, error)
	UpdateTransaction(args types.UpdateTransactionArgs) (*db.TransactionModel, error)
}

type TransactionDao struct {
	client *db.PrismaClient
	ctx    context.Context
}

func (d TransactionDao) ListTransaction(args types.ListTransactionArgs) ([]db.TransactionModel, error) {
	args.Skip = lib.UseDefaultValueIf(0, args.Skip, 0)
	args.Take = lib.UseDefaultValueIf(0, args.Take, 10)

	return d.client.Transaction.
		FindMany().
		Skip(args.Skip).
		Take(args.Take).
		Exec(d.ctx)
}

func (d TransactionDao) CreateTransaction(args types.CreateTransactionArgs) (*db.TransactionModel, error) {
	if args.Reference == "" {
		args.Reference = strings.ToUpper(lib.RandomString(10))
	}

	return d.client.Transaction.
		CreateOne(
			db.Transaction.User.Link(db.User.ID.Equals(args.UserId)),
			db.Transaction.Amount.Set(args.Amount),
			db.Transaction.Status.Set(args.Status),
			db.Transaction.Type.Set(args.Type),
			db.Transaction.Reference.Set(args.Reference),
		).
		Exec(d.ctx)
}

func (d TransactionDao) GetTransaction(args types.GetTransactionArgs) (*db.TransactionModel, error) {
	return d.client.Transaction.
		FindFirst(db.Transaction.ID.Equals(args.Id)).
		Exec(d.ctx)
}

func (d TransactionDao) UpdateTransaction(args types.UpdateTransactionArgs) (*db.TransactionModel, error) {
	return d.client.Transaction.
		FindUnique(db.Transaction.ID.Equals(args.Id)).
		Update(
			db.Transaction.UserID.SetIfPresent(args.UserId),
			db.Transaction.Amount.SetIfPresent(args.Amount),
			db.Transaction.Type.SetIfPresent(args.Type),
			db.Transaction.Status.SetIfPresent(args.Status),
			db.Transaction.Reference.SetIfPresent(args.Reference),
		).
		Exec(d.ctx)
}

func (d TransactionDao) DeleteTransaction(args types.DeleteTransactionArgs) error {
	_, err := d.client.Transaction.
		FindUnique(db.Transaction.ID.Equals(args.Id)).
		Delete().
		Exec(d.ctx)

	return err
}

func NewTransactionDao(connection *lib.DatabaseConnection) ITransactionDao {
	return &TransactionDao{
		client: connection.Client,
		ctx:    context.Background(),
	}
}
