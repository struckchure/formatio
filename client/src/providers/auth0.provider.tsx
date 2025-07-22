import { Auth0Provider as BaseProvider } from "@auth0/auth0-react";

import { AUTH0_REDIRECT_URL } from "@/constants";
import { AUTH0_CLIENT_ID, AUTH0_DOMAIN } from "@/constants/env-vars";
import { ChildrenProps } from "@/types";

export function Auth0Provider({ children }: ChildrenProps) {
  return (
    <BaseProvider
      domain={AUTH0_DOMAIN}
      clientId={AUTH0_CLIENT_ID}
      authorizationParams={{
        redirect_uri: AUTH0_REDIRECT_URL,
        audience: `https://${AUTH0_DOMAIN}/api/v2/`,
        scope: "openid profile email",
      }}
    >
      {children}
    </BaseProvider>
  );
}
