import { AxiosError } from "axios";
import React from "react";

export interface ChildrenProps {
  children: React.ReactNode;
}

export type FieldErrors<T = unknown> = AxiosError<{
  errors: Record<keyof T, string[]>;
}>;
