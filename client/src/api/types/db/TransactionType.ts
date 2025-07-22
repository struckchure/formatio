export const dbTransactionType = {
  TransactionTypeDebit: "DEBIT",
  TransactionTypeCredit: "CREDIT",
} as const;
export type DbTransactionType =
  (typeof dbTransactionType)[keyof typeof dbTransactionType];
