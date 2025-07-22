package services

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type CardService struct {
	cardDao        dao.ICardDao
	userDao        dao.IUserDao
	transactionDao dao.ITransactionDao
	payment        lib.Payment
}

func (s *CardService) PreAuthorizeCard(args types.PreAuthorizeCardArgs) (*types.PreAuthorizeCardResult, error) {
	card, err := s.cardDao.CreateCard(types.CreateCardArgs{
		UserId:      args.UserId,
		CardType:    "unknown", // TODO: fix card type
		LastDigits:  args.CardNumber[len(args.CardNumber)-4:],
		ExpiryMonth: args.ExpiryMonth,
		ExpiryYear:  args.ExpiryYear,
	})
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	user, err := s.userDao.GetUser(types.GetUserArgs{ID: &args.UserId})
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	transaction, err := s.transactionDao.CreateTransaction(types.CreateTransactionArgs{
		UserId: args.UserId,
		Amount: decimal.NewFromInt(100),
		Type:   db.TransactionTypeDebit,
		Status: db.TransactionStatusPending,
	})
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	transactionAmount, _ := strconv.Atoi(transaction.Amount.String())
	userEmail, _ := user.Email()

	cardAuth, err := s.payment.DirectCharge(lib.DirectChargePayload{
		CardNumber:  args.CardNumber,
		Cvv:         args.Cvv,
		ExpiryMonth: args.ExpiryMonth,
		ExpiryYear:  args.ExpiryYear,
		Amount:      int32(transactionAmount),
		TxRef:       transaction.Reference,
		Currency:    string(db.CurrencyNgn),
		Country:     "NG",
		Email:       userEmail,
		Authorization: lib.DirectChargePayloadAuthorization{
			Mode: "pin",
			Pin:  args.Pin,
		},
		RedirectURL: fmt.Sprintf("%s/api/v1/callbacks/flutterwave", "example.com"),
		Meta: lib.DirectChargePayloadMeta{
			CardId: card.ID,
		},
	})

	if err != nil {
		s.cardDao.DeleteCard(types.DeleteCardArgs{Id: card.ID})
		s.transactionDao.DeleteTransaction(types.DeleteTransactionArgs{Id: transaction.ID})

		return nil, lib.HttpError{Message: err.Error(), StatusCode: fiber.StatusBadRequest}
	}

	s.cardDao.UpdateCard(types.UpdateCardArgs{
		Id:       card.ID,
		CardType: &cardAuth.Data.Card.Issuer,
	})

	return &types.PreAuthorizeCardResult{Reference: cardAuth.Data.FlwRef}, err
}

func (s *CardService) AuthorizeCard(args types.AuthorizeCardArgs) (card *db.CardModel, err error) {
	validateCharge, err := s.payment.ValidateCharge(lib.ValidateChargePayload{
		FlwRef: args.Reference,
		Otp:    args.Otp,
		Type:   "card",
	})
	if err != nil {
		return nil, lib.HttpError{Message: err.Error(), StatusCode: fiber.StatusBadRequest}
	}

	paymentTransaction, err := s.payment.FetchTransaction(lib.FetchTransactionPayload{ID: validateCharge.Data.ID})
	if err != nil {
		return nil, lib.HttpError{Message: err.Error(), StatusCode: fiber.StatusBadRequest}
	}

	card, err = s.cardDao.UpdateCard(types.UpdateCardArgs{
		Id:         paymentTransaction.Data.Meta.CardId,
		AuthToken:  &paymentTransaction.Data.Card.Token,
		IsApproved: lo.ToPtr(true),
	})
	if err != nil {
		return nil, err
	}

	_, err = s.payment.RefundCharge(
		lib.RefundChargePayload{
			FlwRef: args.Reference,
			Amount: paymentTransaction.Data.ChargedAmount,
		},
	)
	if err != nil {
		return nil, err
	}

	_, err = s.transactionDao.CreateTransaction(types.CreateTransactionArgs{
		UserId: card.UserID,
		Amount: decimal.NewFromInt32(int32(paymentTransaction.Data.ChargedAmount)),
		Type:   db.TransactionTypeCredit,
		Status: db.TransactionStatusPending,
	})
	if err != nil {
		return nil, lib.TranslateDAOError(err)
	}

	card = lib.RemoveField(card, lib.RemoveFieldOptionFunc("authToken"))

	return card, err
}

func (s *CardService) ListCard(args types.ListCardArgs) ([]db.CardModel, error) {
	args.IsApproved = lo.Ternary(args.IsApproved != nil, args.IsApproved, lo.ToPtr(true))

	cards, err := s.cardDao.ListCard(args)
	cards = lo.Map(cards, func(card db.CardModel, index int) db.CardModel {
		return lib.RemoveField(card, lib.RemoveFieldOptionFunc("authToken"))
	})

	return cards, err
}

func (s *CardService) UpdateCard(args types.UpdateCardArgs) (*db.CardModel, error) {
	if args.IsDefault != nil {
		card, err := s.cardDao.GetCard(types.GetCardArgs{Id: args.Id})
		if err != nil {
			return nil, lib.TranslateDAOError(err)
		}

		allUserCards, err := s.cardDao.ListCard(types.ListCardArgs{UserId: &card.UserID})
		if err != nil {
			return nil, lib.TranslateDAOError(err)
		}

		for _, userCard := range allUserCards {
			if userCard.ID != args.Id {
				s.cardDao.UpdateCard(types.UpdateCardArgs{
					Id:        userCard.ID,
					IsDefault: lo.ToPtr(false),
				})
			}
		}
	}

	card, err := s.cardDao.UpdateCard(args)

	card = lib.RemoveField(card, lib.RemoveFieldOptionFunc("authToken"))

	return card, err
}

func (s *CardService) DeleteCard(args types.DeleteCardArgs) error {
	return s.cardDao.DeleteCard(args)
}

func NewCardService(
	cardDao dao.ICardDao,
	userDao dao.IUserDao,
	transactionDao dao.ITransactionDao,
	payment lib.Payment,
) CardService {
	return CardService{
		cardDao:        cardDao,
		userDao:        userDao,
		transactionDao: transactionDao,
		payment:        payment,
	}
}
