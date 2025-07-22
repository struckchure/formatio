package tasks

import (
	"encoding/json"
	"log"
	"time"

	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type BillingTasks struct {
	rmq            lib.RabbitMQ
	invoiceDao     dao.IInvoiceDao
	machineDao     dao.IMachineDao
	invoiceService services.InvoiceService
}

func (t *BillingTasks) SleepUntilNextMidnight() {
	now := time.Now()
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	duration := nextMidnight.Sub(now)
	time.Sleep(duration)
}

func (t *BillingTasks) ScheduleMachineInvoicesTask() {
	for {
		log.Println("[ScheduleMachineInvoicesTask]: Running ...")

		t.SleepUntilNextMidnight()

		machines, _ := t.machineDao.ListMachines(types.ListMachineArgs{})
		for _, machine := range machines {
			err := t.invoiceService.ScheduleInvoice(machine)
			if err != nil {
				log.Println("[ScheduleMachineInvoicesTask]: ", err)
			}
		}
	}
}

func (t *BillingTasks) ProcessInvoiceTask() {
	err := t.rmq.Subscribe(
		lib.SubscribeArgs{
			Queue: types.CREATE_INVOICE_QUEUE,
			Callback: func(content string) error {
				var invoice db.InvoiceModel
				err := json.Unmarshal([]byte(content), &invoice)
				if err != nil {
					log.Fatalln("[CreateInvoice] ", err)
				}

				return t.invoiceService.ProcessInvoice(invoice)
			},
		},
	)
	if err != nil {
		log.Println("[ProcessInvoiceTask]: ", err)
	}
}

func NewBillingTasks(
	rmq lib.RabbitMQ,
	invoiceDao dao.IInvoiceDao,
	machineDao dao.IMachineDao,
	invoiceService services.InvoiceService,
) BillingTasks {
	return BillingTasks{
		rmq:            rmq,
		invoiceDao:     invoiceDao,
		machineDao:     machineDao,
		invoiceService: invoiceService,
	}
}
