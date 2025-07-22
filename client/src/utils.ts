import { DbDeploymentLogModel } from "@/api";
import { FieldErrors } from "@/types";
import _ from "lodash";

export function convertToCamelCase(str: string): string {
  return str
    .replace(/([-_][a-z])/gi, ($1) => {
      return $1.toUpperCase().replace("-", "").replace("_", "");
    })
    .replace(str[0], str[0].toLowerCase());
}

export function joinPaths(...params: string[]): string {
  // Filter and clean up the path segments
  const cleanedSegments = params.map((segment) =>
    segment.replace(/^\/+|\/+$/g, "")
  );

  // Join the cleaned segments with a single slash between them
  return cleanedSegments.join("/");
}

export function randomString(length: number = 12) {
  let result = "";
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  const charactersLength = characters.length;

  for (let i = 0; i < length; i++)
    result += characters.charAt(Math.floor(Math.random() * charactersLength));

  return result;
}

export function formatDate(date: string | Date) {
  const newDate = new Date(date);

  return newDate.toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });
}

export function replaceWith(text: string, target = " ", seperator = "_") {
  return text.toLowerCase().replace(target, seperator);
}

export function popKeys<T extends object, K extends keyof T>(
  object: T,
  ...keys: K[]
): Partial<T> {
  _.omit(object, keys);

  return object;
}

export function formatCurrency(price: number) {
  return `₦ ${new Intl.NumberFormat().format(price)}`;
}

export function formatFieldErrors<T = unknown>(error: Error) {
  return (error as unknown as FieldErrors<T>)?.response?.data?.errors;
}

export interface Step {
  name: string;
  message: string;
}

export interface Job {
  name: string;
  steps: Step[];
}

export function formatLogsToJobs(logs: DbDeploymentLogModel[]): Job[] {
  const jobs: Job[] = [];

  const logExists = (jobName: string) => {
    if (jobs.find((job) => job.name === jobName)) {
      return true;
    }

    return false;
  };

  for (let i = 0; i < logs.length; i++) {
    const log = logs[i];
    const jobName = log.jobId.split(".")[0];
    if (!logExists(jobName)) {
      jobs.push({ name: jobName, steps: [] });
    }

    const jobIndex = jobs.findIndex((job) => job.name === jobName);
    if (!jobs[jobIndex].steps.find((step) => step.message === log.message)) {
      jobs[jobIndex].steps.push({
        name: log.jobId.split(".")[1],
        message: log.message,
      });
    }
  }

  return jobs;
}
