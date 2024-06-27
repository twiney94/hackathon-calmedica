import { config } from "dotenv";

config();

const API_BASE_URL = process.env.API_BASE_URL as string;

export const performHttpCall = async (
  path: string,
  method: string = "GET",
  body: any = undefined,
  file: boolean = false
): Promise<any> => {
  try {
    const headers: Record<string, string> = {
      Accept: file ? "*/*" : "application/json",
    };

    if (!(body instanceof FormData) && !file) {
      headers["Content-Type"] = "application/json;charset=UTF-8";
    }

    const options: RequestInit = {
      method,
      headers,
      body: body instanceof FormData ? body : JSON.stringify(body),
    };

    const response = await fetch(`${API_BASE_URL}${path}`, options);

    if (!response.ok) {
      const errorResponse = await response.json();
      throw new Error(errorResponse.message || response.statusText);
    }

    if (file) {
      const blob = await response.blob();
      return blob;
    }

    const data = await response.json();
    return data;
  } catch (error: any) {
    throw new Error(error.message);
  }
};
