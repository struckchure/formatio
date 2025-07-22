import { useState } from "react";

import { useListCards, useListInvoice } from "@/api/hooks/billing";
import { Button } from "@/components/button";
import { Modal } from "@/components/modal";
import { CardItem } from "@/pages/billing/components/card-item";
import { InvoiceItem } from "@/pages/billing/components/invoice-item";
import { NewCardForm } from "@/pages/billing/components/new-card-form";

export function BillingPage() {
  const listInvoice = useListInvoice();
  const listCards = useListCards();

  const [isNewCardModalVisible, setNewCardModalVisibility] =
    useState<boolean>(false);

  return (
    <>
      <Modal
        visibility={isNewCardModalVisible}
        setVisibility={setNewCardModalVisibility}
      >
        <NewCardForm />
      </Modal>

      <section className="w-full flex gap-4 items-start justify-start">
        <div className="w-[70%] p-4 bg-[rgba(0,0,0,0.125)]">
          <h4 className="text-xl">Invoice</h4>

          <div className="w-full py-2 flex gap-2 flex-col items-start justify-start">
            {listInvoice.data?.map((invoice, index) => (
              <InvoiceItem key={index} {...invoice} />
            ))}
          </div>
        </div>

        <div className="w-[30%] p-4 bg-[rgba(0,0,0,0.125)]">
          <div className="w-full flex items-center justify-between">
            <h4 className="text-xl">Cards</h4>

            <Button
              size="sm"
              className="rounded-none"
              onClick={() => setNewCardModalVisibility(!isNewCardModalVisible)}
            >
              New Card
            </Button>
          </div>

          <div className="w-full py-2 flex gap-2 flex-col items-start justify-start">
            {listCards.data?.map((card, index) => (
              <CardItem key={index} {...card} />
            ))}
          </div>
        </div>
      </section>
    </>
  );
}
