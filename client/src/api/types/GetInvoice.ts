import type { DbInvoiceModel } from "./db/InvoiceModel";

export type GetInvoicePathParams = {
  /**
   * @description Invoice Id
   * @type string
   */
  invoiceId: string;
};
/**
 * @description OK
 */
export type GetInvoice200 = DbInvoiceModel;
/**
 * @description OK
 */
export type GetInvoiceQueryResponse = DbInvoiceModel;
export type GetInvoiceQuery = {
  Response: GetInvoiceQueryResponse;
  PathParams: GetInvoicePathParams;
};
