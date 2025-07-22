export const dbCurrency = {
  CurrencyNgn: "NGN",
} as const;
export type DbCurrency = (typeof dbCurrency)[keyof typeof dbCurrency];
