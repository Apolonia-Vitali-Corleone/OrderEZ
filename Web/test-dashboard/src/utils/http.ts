import { formatJson } from "./format";

export interface ParsedResponse {
  ok: boolean;
  status: number;
  statusText: string;
  json: unknown | null;
  text: string;
}

export async function parseResponse(response: Response): Promise<ParsedResponse> {
  const text = await response.text();
  try {
    const json = text ? JSON.parse(text) : null;
    return { ok: response.ok, status: response.status, statusText: response.statusText, json, text: formatJson(json ?? text) };
  } catch (error) {
    return { ok: response.ok, status: response.status, statusText: response.statusText, json: null, text: text || "" };
  }
}

export async function fetchJson(input: RequestInfo, init?: RequestInit): Promise<ParsedResponse> {
  const response = await fetch(input, init);
  return parseResponse(response);
}
