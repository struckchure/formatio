import { defineConfig } from "@kubb/core";
import { pluginOas } from "@kubb/plugin-oas";
import { pluginTanstackQuery } from "@kubb/swagger-tanstack-query";
import { pluginTs } from "@kubb/swagger-ts";

// import * as queryKey from "./.kubb/query-key.template";

export default defineConfig(() => {
  return {
    root: ".",
    input: { path: "../docs/swagger.yaml" },
    output: { path: "./src/api" },
    hooks: { done: ["npx prettier --write ./src/api"] },
    plugins: [
      pluginOas({}),
      pluginTs({ dateType: "date", unknownType: "unknown" }),
      pluginTanstackQuery({
        output: { path: "./hooks" },
        group: { type: "tag", output: "./hooks/{{tag}}" },
        transformers: { name: (name: string) => name },
        pathParamsType: "object",
        client: { importPath: "@/api/client.ts" },
        // templates: { queryKey: queryKey.templates },
        mutate: {
          methods: ["post", "put", "patch", "delete"],
          variablesType: "mutate",
        },
      }),
    ],
  };
});
