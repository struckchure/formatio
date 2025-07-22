export const dbDeploymentStatus = {
  DeploymentStatusIdle: "IDLE",
  DeploymentStatusInProgress: "IN_PROGRESS",
  DeploymentStatusSuccessful: "SUCCESSFUL",
  DeploymentStatusFailed: "FAILED",
} as const;
export type DbDeploymentStatus =
  (typeof dbDeploymentStatus)[keyof typeof dbDeploymentStatus];
