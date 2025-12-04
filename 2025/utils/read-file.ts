import { getArgs } from "./args";

export const readFile = async (filePath: string) => {
  try {
    const file = Bun.file(filePath);
    const text = await file.text();
    return text;
  } catch (error) {
    console.error("Error reading file");
    process.exit(1);
  }
};

export const readFileFromArgs = async () => {
  const args = getArgs();

  const filePath = args.file;

  if (!filePath) {
    console.error("Must provide a file path");
    process.exit(1);
  }

  return await readFile(filePath);
};
