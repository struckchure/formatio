import classNames from "classnames";
import { FiCheck as CheckIcon } from "react-icons/fi";
import { LiaTimesSolid as DeleteIcon } from "react-icons/lia";
import { PiArrowsCounterClockwise as RunningIcon } from "react-icons/pi";

import { DbDeploymentStatus } from "@/api";

export function StatusIcon({ status }: { status: DbDeploymentStatus }) {
  return (
    <div
      className={classNames(
        "p-1 rounded-full",
        status == "SUCCESSFUL" && "bg-success",
        status == "FAILED" && "bg-danger text-primary",
        status == "IN_PROGRESS" && "bg-warning text-primary animate-spin"
      )}
    >
      {status == "SUCCESSFUL" && <CheckIcon size={14} />}
      {status == "FAILED" && <DeleteIcon size={14} />}
      {status == "IN_PROGRESS" && <RunningIcon size={14} />}
    </div>
  );
}
