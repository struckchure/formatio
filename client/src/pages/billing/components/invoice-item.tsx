import moment from "moment";

import { DbInvoiceModel } from "@/api";
import { formatCurrency } from "@/utils";

export function InvoiceItem(props: DbInvoiceModel) {
  return (
    <div className="w-full p-2 flex items-center justify-between bg-primary">
      <label className="w-full">{moment(props.to).calendar()}</label>
      <label className="w-full">{props.description}</label>
      <label className="w-full">{props.status}</label>
      <label className="w-[200px]">{formatCurrency(props.totalPrice)}</label>
    </div>
  );
}
