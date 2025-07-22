package types

import (
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
)

type ListTransactionArgs struct {
	lib.BaseListFilterArgs
}

type CreateTransactionArgs struct {
	UserId    string
	Amount    db.Decimal
	Type      db.TransactionType
	Status    db.TransactionStatus
	Reference string
}

type GetTransactionArgs struct {
	Id string
}

type UpdateTransactionArgs struct {
	Id        string
	UserId    *string
	Amount    *db.Decimal
	Type      *db.TransactionType
	Status    *db.TransactionStatus
	Reference *string
}

type DeleteTransactionArgs struct {
	Id string
}
