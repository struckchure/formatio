export const dbTransactionStatus = {
  TransactionStatusPending: "PENDING",
  TransactionStatusSuccess: "SUCCESS",
  TransactionStatusFailed: "FAILED",
} as const;
export type DbTransactionStatus =
  (typeof dbTransactionStatus)[keyof typeof dbTransactionStatus];
