"use client";

import "@/styles/globals.scss";

import { ErrorBoundary, FallbackProps } from "react-error-boundary";
import { RodelarProvider } from "rodelar/react";

import { Button } from "@/components/button";
import { AppLayout } from "@/layouts/applayout";
import { Auth0Provider } from "@/providers/auth0.provider";
import { QueryClientProvider } from "@/providers/query-client.provider";
import { ToastProvider } from "@/providers/toast.provider";
import { Router } from "@/router";
import { RODELAR_API_URL } from "./constants/env-vars";

function fallbackRender({ error, resetErrorBoundary }: FallbackProps) {
  // Call resetErrorBoundary() to reset the error boundary and retry the render.

  return (
    <div
      className="w-full h-screen grid place-content-center place-items-center"
      role="alert"
    >
      <p>Something went wrong:</p>
      <pre style={{ color: "red", width: "60%" }}>
        {JSON.stringify(error.message)}
      </pre>
      <Button
        onClick={resetErrorBoundary}
        className="rounded-none mx-auto my-2"
        size="sm"
      >
        Reload
      </Button>
    </div>
  );
}

export function App() {
  return (
    <QueryClientProvider>
      <RodelarProvider url={RODELAR_API_URL}>
        <AppLayout>
          <ErrorBoundary
            fallbackRender={fallbackRender}
            onReset={() => {
              // Reset the state of your app so the error doesn't happen again
              window.location.reload();
            }}
          >
            <Auth0Provider>
              <Router />
              <ToastProvider />
            </Auth0Provider>
          </ErrorBoundary>
        </AppLayout>
      </RodelarProvider>
    </QueryClientProvider>
  );
}
