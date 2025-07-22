import { TypesListInvoiceArgs } from "./types/ListInvoiceArgs";
import type { DbInvoiceModel } from "./db/InvoiceModel";

/**
 * @description Accepted
 */
export type ListInvoice202 = DbInvoiceModel[];
/**
 * @description List Invoice
 */
export type ListInvoiceQueryRequest = TypesListInvoiceArgs;
/**
 * @description Accepted
 */
export type ListInvoiceQueryResponse = DbInvoiceModel[];
export type ListInvoiceQuery = {
  Response: ListInvoiceQueryResponse;
  Request: ListInvoiceQueryRequest;
};
