import { parseArgs } from "util";

export const getArgs = () => {
  const { values, positionals } = parseArgs({
    args: Bun.argv,
    options: {
      file: {
        type: "string",
      },
    },
    strict: true,
    allowPositionals: true,
  });

  return values;
};
