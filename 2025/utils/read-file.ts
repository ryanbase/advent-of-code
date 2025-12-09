import { getArgs } from "./args";

export const getFileName = () => {
  const { file, test } = getArgs();
  return file ?? test ? "test.txt" : "input.txt";
};

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
  return await readFile(getFileName());
};
