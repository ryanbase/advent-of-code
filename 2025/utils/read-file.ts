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

  const filePath = args.file ?? args.test ? "test.txt" : "input.txt";

  return await readFile(filePath);
};
