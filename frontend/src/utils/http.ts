const env = import.meta.env;
const API_BASE_URL = env.VITE_API_BASE_URL as string;

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

    console.log("🚀 ~ `${API_BASE_URL}${path}`:", `${API_BASE_URL}${path}`);
    const response = await fetch(`${API_BASE_URL}${path}`, options);

    if (!response.ok) {
      const errorResponse = await response.json();
      throw new Error(errorResponse.message || response.statusText);
    }

    const data = await response.json();
    return data;
  } catch (error: any) {
    throw new Error(error.message);
  }
};
