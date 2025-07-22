import { DbDeploymentLogModel, dbDeploymentStatus } from "@/api";
import {
  useGetDeploymentById,
  useListDeploymentLogsById,
} from "@/api/hooks/deployments";
import { useParams } from "@/api/hooks/router";
import { Button } from "@/components/button";
import { formatLogsToJobs, Job, Step } from "@/utils";
import { useEffect, useState } from "react";
import { useRodelar } from "rodelar/react";
import { OverrideProperties } from "type-fest";
import { Accordion } from "./accordion";

export function DeploymentLogs() {
  const { deploymentId } = useParams().params as { deploymentId: string };

  const [activeJob, setActiveJob] = useState<string | undefined>();
  const [formattedLogs, setFormattedLogs] = useState<Job[]>([]);

  const getDeployment = useGetDeploymentById({ deploymentId });
  const listDeploymentLogs = useListDeploymentLogsById({ deploymentId });

  const { subscribe, unsubscribe } = useRodelar();

  // Update logs when the API call for existing logs succeeds
  useEffect(() => {
    if (listDeploymentLogs.isSuccess) {
      setFormattedLogs(formatLogsToJobs(listDeploymentLogs.data));
    }
  }, [listDeploymentLogs.isSuccess, listDeploymentLogs.data]);

  useEffect(() => {
    if (
      getDeployment.data?.status ===
      dbDeploymentStatus.DeploymentStatusInProgress
    ) {
      subscribe({
        event: `deployment-log-stream-event/${deploymentId}`,
        callback(data) {
          if (data.message) {
            const log = JSON.parse(data.message) as DbDeploymentLogModel;

            setFormattedLogs((prev) => [...prev, ...formatLogsToJobs([log])]);
          }
        },
      });
    }
  }, [deploymentId, getDeployment.data?.status, subscribe]);

  useEffect(() => {
    // Clean-up: Unsubscribe when the component unmounts or deployment status changes
    return () => {
      unsubscribe({ event: `deployment-log-stream-event/${deploymentId}` });
    };
  }, [deploymentId, unsubscribe]);

  // Format logs when deploymentLogs are updated
  useEffect(() => {
    // Auto-select the first job if activeJob is not set
    if (formattedLogs.length > 0 && !activeJob) {
      setActiveJob(
        formattedLogs.at(
          getDeployment.data?.status ===
            dbDeploymentStatus.DeploymentStatusInProgress
            ? -1
            : 0
        )?.name
      );
    }
  }, [activeJob, formattedLogs, getDeployment.data?.status]);

  const currentJobLogs: OverrideProperties<Step, { message: string[] }>[] = [];

  // const currentJobLogs = useCallback(() => {
  //   const job = formattedLogs.find((job) => job.name === activeJob);
  //   if (!job) return [];

  //   // Merge steps with the same name and accumulate their messages
  //   const stepsWithMessages: OverrideProperties<Step, { message: string[] }>[] =
  //     [];
  //   job.steps.forEach((step) => {
  //     const existingStepIndex = stepsWithMessages.findIndex(
  //       (s) => s.name === step.name
  //     );
  //     if (existingStepIndex > -1) {
  //       stepsWithMessages[existingStepIndex].message.push(step.message);
  //     } else {
  //       stepsWithMessages.push({ name: step.name, message: [step.message] });
  //     }
  //   });

  //   return stepsWithMessages;
  // }, [activeJob, formattedLogs]);

  return (
    <div className="w-full h-full flex justify-start">
      {/* Jobs navigation */}
      <div className="w-[10%] h-full flex flex-col sticky top-36 divide-y-2 divide-white/5">
        {formattedLogs.map((job, index) => (
          <Button
            key={index}
            size="md"
            onClick={() => setActiveJob(job.name)}
            className={`first:rounded-t-md rounded-none last:rounded-b-md ${
              job.name === activeJob ? "bg-white/20" : ""
            }`}
          >
            {job.name}
          </Button>
        ))}
      </div>

      {/* Job logs display */}
      <div className="w-[90%] h-full flex flex-col items-center divide-y-2 divide-white/5 px-4">
        {currentJobLogs.map((step, index) => (
          <Accordion key={index} title={step.name}>
            <div className="bg-black rounded-md overflow-auto my-2 p-4 w-full">
              {step.message.map((message, msgIndex) => (
                <pre
                  key={msgIndex}
                  className="hover:bg-white/5 hover:cursor-pointer"
                >
                  {msgIndex + 1}. {message}
                </pre>
              ))}
            </div>
          </Accordion>
        ))}
      </div>
    </div>
  );
}
