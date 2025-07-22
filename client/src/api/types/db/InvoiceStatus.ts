export const dbInvoiceStatus = {
  InvoiceStatusPaid: "PAID",
  InvoiceStatusUnpaid: "UNPAID",
} as const;
export type DbInvoiceStatus =
  (typeof dbInvoiceStatus)[keyof typeof dbInvoiceStatus];
