import { type URLObject, URLPath } from "@kubb/core/utils";
import { useOperation, useOperationManager } from "@kubb/plugin-oas/hooks";
import { Function, Type } from "@kubb/react";
import { QueryKey } from "@kubb/swagger-tanstack-query/components";
import { pluginTsName } from "@kubb/swagger-ts";
import React from "react";

export const templates = {
  ...QueryKey.templates,
  react: function ({
    name,
    typeName,
    params,
    generics,
    returnType,
    JSDoc,
  }: React.ComponentProps<typeof QueryKey.templates.react>) {
    const operation = useOperation();
    const { getSchemas } = useOperationManager();

    const schemas = getSchemas(operation, {
      pluginKey: [pluginTsName],
      type: "type",
    });
    const path = new URLPath(operation.path);
    const withQueryParams = !!schemas.queryParams?.name;

    const pathObject = path.toObject({
      type: "path",
    }) as URLObject;

    // Filter out empty segments and replace dynamic parameters with their respective variable names
    const pathSegments = pathObject.url
      .split("/")
      .filter((segment) => segment !== "" && segment.startsWith(":"))
      .map((segment) => segment.startsWith(":") && segment.slice(1));

    const keys = [
      `"${operation.getOperationId()}"`,
      ...pathSegments,
      withQueryParams ? "...(params ? [params] : [])" : undefined,
    ].filter(Boolean);

    return (
      <>
        <Function.Arrow
          name={name}
          export
          generics={generics}
          params={params}
          returnType={returnType}
          singleLine
          JSDoc={JSDoc}
        >
          {`[${keys}] as const`}
        </Function.Arrow>

        <Type name={typeName} export>
          {`ReturnType<typeof ${name}>`}
        </Type>
      </>
    );
  },
} as const;
