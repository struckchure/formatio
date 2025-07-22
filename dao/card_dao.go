package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type ICardDao interface {
	CreateCard(args types.CreateCardArgs) (*db.CardModel, error)
	DeleteCard(args types.DeleteCardArgs) error
	GetCard(args types.GetCardArgs) (*db.CardModel, error)
	ListCard(args types.ListCardArgs) ([]db.CardModel, error)
	UpdateCard(args types.UpdateCardArgs) (*db.CardModel, error)
}

type CardDao struct {
	client *db.PrismaClient
	ctx    context.Context
}

func (d *CardDao) ListCard(args types.ListCardArgs) ([]db.CardModel, error) {
	args.Skip = lib.UseDefaultValueIf(0, args.Skip, 0)
	args.Take = lib.UseDefaultValueIf(0, args.Take, 10)

	return d.client.Card.
		FindMany(
			db.Card.UserID.EqualsIfPresent(args.UserId),
			db.Card.IsApproved.EqualsIfPresent(args.IsApproved),
		).
		Skip(args.Skip).
		Take(args.Take).
		OrderBy(db.Card.UpdatedAt.Order(db.DESC)).
		Exec(d.ctx)
}

func (d *CardDao) CreateCard(args types.CreateCardArgs) (*db.CardModel, error) {
	return d.client.Card.
		CreateOne(
			db.Card.User.Link(db.User.ID.Equals(args.UserId)),
			db.Card.AuthToken.Set(args.AuthToken),
			db.Card.IsDefault.Set(args.IsDefault),
			db.Card.IsApproved.Set(args.IsApproved),
			db.Card.CardType.Set(args.CardType),
			db.Card.LastDigits.Set(args.LastDigits),
			db.Card.ExpiryMonth.Set(args.ExpiryMonth),
			db.Card.ExpiryYear.Set(args.ExpiryYear),
		).
		Exec(d.ctx)
}

func (d *CardDao) GetCard(args types.GetCardArgs) (*db.CardModel, error) {
	return d.client.Card.
		FindFirst(
			db.Card.ID.Equals(args.Id),
			db.Card.UserID.EqualsIfPresent(args.UserId),
			db.Card.IsDefault.EqualsIfPresent(args.IsDefault),
			db.Card.IsApproved.EqualsIfPresent(args.IsApproved),
		).
		Exec(d.ctx)
}

func (d *CardDao) UpdateCard(args types.UpdateCardArgs) (*db.CardModel, error) {
	return d.client.Card.
		FindUnique(db.Card.ID.Equals(args.Id)).
		With(db.Card.User.Fetch()).
		Update(
			db.Card.AuthToken.SetIfPresent(args.AuthToken),
			db.Card.IsDefault.SetIfPresent(args.IsDefault),
			db.Card.IsApproved.SetIfPresent(args.IsApproved),
			db.Card.CardType.SetIfPresent(args.CardType),
			db.Card.LastDigits.SetIfPresent(args.LastDigits),
			db.Card.ExpiryMonth.SetIfPresent(args.ExpiryMonth),
			db.Card.ExpiryYear.SetIfPresent(args.ExpiryYear),
		).
		Exec(d.ctx)
}

func (d *CardDao) DeleteCard(args types.DeleteCardArgs) error {
	_, err := d.client.Card.
		FindUnique(db.Card.ID.Equals(args.Id)).
		Exec(d.ctx)

	return err
}

func NewCardDao(connection *lib.DatabaseConnection) ICardDao {
	return &CardDao{
		client: connection.Client,
		ctx:    context.Background(),
	}
}
