export const dbMachineStatus = {
  MachineStatusCreating: "CREATING",
  MachineStatusRunning: "RUNNING",
  MachineStatusRebooting: "REBOOTING",
  MachineStatusShuttingDown: "SHUTTING_DOWN",
  MachineStatusShutdown: "SHUTDOWN",
} as const;
export type DbMachineStatus =
  (typeof dbMachineStatus)[keyof typeof dbMachineStatus];
