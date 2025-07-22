package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"

	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type InvoiceService struct {
	rmq            lib.RabbitMQ
	invoiceDao     dao.IInvoiceDao
	cardDao        dao.ICardDao
	userDao        dao.IUserDao
	transactionDao dao.ITransactionDao
	payment        lib.Payment
	scheduler      *lib.Scheduler
}

func (s *InvoiceService) ScheduleInvoice(machine db.MachineModel) error {
	var (
		lastChargedInvoice time.Time
		nextInvoiceDate    time.Time
	)

	lastInvoice, _ := s.invoiceDao.LastInvoice(types.LastInvoiceArgs{ProductId: machine.ID})
	if lastInvoice != nil {
		lastChargedInvoice = lastInvoice.CreatedAt
	} else {
		lastChargedInvoice = machine.CreatedAt
	}

	nextInvoiceDate = lastChargedInvoice.AddDate(0, 0, 30)

	machinePlan := machine.Plan()
	machinePlanMonthlyRate, _ := machinePlan.MonthlyRate()
	machineName, _ := machine.MachineName()

	return s.scheduler.Schedule(
		lib.WithFrequency("@every"),
		lib.WithDuration(nextInvoiceDate.Sub(lastChargedInvoice)),
		lib.WithCallback(func() error {
			invoice, err := s.invoiceDao.CreateInvoice(types.CreateInvoiceArgs{
				UserId:      machine.OwnerID,
				ProductId:   machine.ID,
				Description: fmt.Sprintf("%s payment for %s", nextInvoiceDate.Month(), machineName),
				From:        lastChargedInvoice,
				To:          nextInvoiceDate,
				Status:      db.InvoiceStatusUnpaid,
				Quantity:    1,
				UnitPrice:   machinePlanMonthlyRate,
				TotalPrice:  machinePlanMonthlyRate,
				Currency:    db.CurrencyNgn,
			})
			if err != nil {
				log.Println("[MachineInvoiceCronService] ", err)

				return err
			}

			payload, err := json.Marshal(invoice)
			if err != nil {
				log.Println("[MachineInvoiceCronService] ", err)

				return err
			}

			err = s.rmq.Publish(lib.PublishArgs{
				Queue:   types.CREATE_INVOICE_QUEUE,
				Content: string(payload),
			})
			if err != nil {
				log.Println("[MachineInvoiceCronService] ", err)

				return err
			}

			return nil
		}),
	)
}

func (s *InvoiceService) ProcessInvoice(invoice db.InvoiceModel) error {
	card, err := s.cardDao.GetCard(types.GetCardArgs{
		UserId:     &invoice.UserID,
		IsDefault:  lo.ToPtr(true),
		IsApproved: lo.ToPtr(true),
	})
	if err != nil {
		log.Printf("[ProcessInvoice]: %s", err)

		return nil
	}

	invoiceTotalPrice, _ := strconv.Atoi(invoice.TotalPrice.String())

	if invoiceTotalPrice > 0 {
		user, err := s.userDao.GetUser(types.GetUserArgs{ID: &invoice.UserID})
		if err != nil {
			log.Printf("[ProcessInvoice]: %s", err)

			return err
		}

		cardAuthToken, _ := card.AuthToken()
		userEmail, _ := user.Email()
		invoiceDescription, _ := invoice.Description()

		_, err = s.payment.RecurringCharge(lib.RecurringChargePayload{
			Token:     cardAuthToken,
			Currency:  string(db.CurrencyNgn),
			Country:   "NG",
			Amount:    invoiceTotalPrice,
			Email:     userEmail,
			TxRef:     strings.ToUpper(lib.RandomString(10)),
			Narration: invoiceDescription,
		})
		if err != nil {
			log.Printf("[ProcessInvoice]: %s", err)

			return err
		}
	}

	_, err = s.invoiceDao.UpdateInvoice(types.UpdateInvoiceArgs{
		Id:     invoice.ID,
		Status: lo.ToPtr(db.InvoiceStatusPaid),
	})
	if err != nil {
		log.Printf("[ProcessInvoice]: %s", err)

		return err
	}

	return nil
}

func (s *InvoiceService) ListInvoice(args types.ListInvoiceArgs) ([]db.InvoiceModel, error) {
	return s.invoiceDao.ListInvoice(args)
}

func (s *InvoiceService) GetInvoice(args types.GetInvoiceArgs) (*db.InvoiceModel, error) {
	return s.invoiceDao.GetInvoice(args)
}

func NewInvoiceService(
	rmq lib.RabbitMQ,
	invoiceDao dao.IInvoiceDao,
	cardDao dao.ICardDao,
	userDao dao.IUserDao,
	transactionDao dao.ITransactionDao,

	payment lib.Payment,
	scheduler *lib.Scheduler,
) InvoiceService {
	return InvoiceService{
		rmq:            rmq,
		invoiceDao:     invoiceDao,
		cardDao:        cardDao,
		userDao:        userDao,
		transactionDao: transactionDao,

		payment:   payment,
		scheduler: scheduler,
	}
}
