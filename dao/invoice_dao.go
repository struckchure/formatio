package dao

import (
	"context"

	"github.com/samber/lo"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IInvoiceDao interface {
	CreateInvoice(args types.CreateInvoiceArgs) (*db.InvoiceModel, error)
	DeleteInvoice(args types.DeleteInvoiceArgs) error
	GetInvoice(args types.GetInvoiceArgs) (*db.InvoiceModel, error)
	LastInvoice(args types.LastInvoiceArgs) (*db.InvoiceModel, error)
	ListInvoice(args types.ListInvoiceArgs) ([]db.InvoiceModel, error)
	UpdateInvoice(args types.UpdateInvoiceArgs) (*db.InvoiceModel, error)
}

type InvoiceDao struct {
	client *db.PrismaClient
	ctx    context.Context
}

func (d *InvoiceDao) ListInvoice(args types.ListInvoiceArgs) ([]db.InvoiceModel, error) {
	return d.client.Invoice.
		FindMany(db.Invoice.User.Where(db.User.ID.EqualsIfPresent(args.UserId))).
		Exec(d.ctx)
}

func (d *InvoiceDao) CreateInvoice(args types.CreateInvoiceArgs) (*db.InvoiceModel, error) {
	return d.client.Invoice.
		CreateOne(
			db.Invoice.User.Link(db.User.ID.Equals(args.UserId)),
			db.Invoice.UnitPrice.Set(args.UnitPrice),
			db.Invoice.TotalPrice.Set(args.TotalPrice),
			db.Invoice.ProductID.Set(args.ProductId),
			db.Invoice.Reference.Set(args.Reference),
			db.Invoice.Description.Set(args.Description),
			db.Invoice.From.Set(args.From),
			db.Invoice.To.Set(args.To),
			db.Invoice.Quantity.Set(args.Quantity),
			db.Invoice.Currency.Set(args.Currency),
			db.Invoice.Status.SetIfPresent(lo.ToPtr(args.Status)),
		).
		Exec(d.ctx)
}

func (d *InvoiceDao) GetInvoice(args types.GetInvoiceArgs) (*db.InvoiceModel, error) {
	return d.client.Invoice.
		FindFirst(db.Invoice.ID.Equals(args.Id)).
		Exec(d.ctx)
}

func (d *InvoiceDao) UpdateInvoice(args types.UpdateInvoiceArgs) (*db.InvoiceModel, error) {
	return d.client.Invoice.
		FindUnique(db.Invoice.ID.Equals(args.Id)).
		With(db.Invoice.User.Fetch()).
		Update(
			db.Invoice.ProductID.SetIfPresent(args.ProductId),
			db.Invoice.Reference.SetIfPresent(args.Reference),
			db.Invoice.Description.SetIfPresent(args.Description),
			db.Invoice.From.SetIfPresent(args.From),
			db.Invoice.To.SetIfPresent(args.To),
			db.Invoice.Quantity.SetIfPresent(args.Quantity),
			db.Invoice.UnitPrice.SetIfPresent(args.UnitPrice),
			db.Invoice.TotalPrice.SetIfPresent(args.TotalPrice),
			db.Invoice.Currency.SetIfPresent(args.Currency),
			db.Invoice.Status.SetIfPresent(args.Status),
		).
		Exec(d.ctx)
}

func (d *InvoiceDao) DeleteInvoice(args types.DeleteInvoiceArgs) error {
	_, err := d.client.Invoice.
		FindUnique(db.Invoice.ID.Equals(args.Id)).
		Delete().
		Exec(d.ctx)

	return err
}

func (d *InvoiceDao) LastInvoice(args types.LastInvoiceArgs) (*db.InvoiceModel, error) {
	return d.client.Invoice.
		FindFirst(db.Invoice.ProductID.Equals(args.ProductId)).
		OrderBy(db.Invoice.CreatedAt.Order(db.DESC)).
		Exec(d.ctx)
}

func NewInvoiceDao(connection *lib.DatabaseConnection) IInvoiceDao {
	return &InvoiceDao{
		client: connection.Client,
		ctx:    context.Background(),
	}
}
