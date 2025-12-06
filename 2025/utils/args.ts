import { parseArgs } from "util";

export const getArgs = () => {
  const { values, positionals } = parseArgs({
    args: Bun.argv,
    options: {
      file: {
        type: "string",
      },
      test: {
        type: "boolean",
      },
    },
    strict: true,
    allowPositionals: true,
  });

  return values;
};
