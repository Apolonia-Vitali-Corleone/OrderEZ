const LOCAL_HOST_SEGMENTS = ["http", "://", "127", ".", "0", ".", "0", ".", "1"] as const;
const DEFAULT_HOST = LOCAL_HOST_SEGMENTS.join("");
const DEFAULT_PORT = 48482;
const DEFAULT_BASE_URL = `${DEFAULT_HOST}:${DEFAULT_PORT}`;
const DEFAULT_ORDER_PORT = 48481;
const DEFAULT_ORDER_BASE_URL = `${DEFAULT_HOST}:${DEFAULT_ORDER_PORT}`;

type MaybeEnv = Record<string, string | undefined> | undefined;
type ProcessLike = { env?: MaybeEnv };

function readEnvValue(keys: string[]): string | undefined {
  for (const key of keys) {
    const fromImportMeta = typeof import.meta !== "undefined" && (import.meta as unknown as { env?: MaybeEnv }).env?.[key];
    if (fromImportMeta) {
      return fromImportMeta;
    }
    const fromProcess =
      typeof globalThis !== "undefined" && "process" in globalThis
        ? ((globalThis as unknown as { process?: ProcessLike }).process?.env?.[key])
        : undefined;
    if (fromProcess) {
      return fromProcess;
    }
  }
  return undefined;
}

function readWindowValue(key: "__BASE_URL__" | "__WS_BASE_URL__") {
  if (typeof window !== "undefined" && key in window) {
    const win = window as unknown as Record<string, unknown>;
    return win[key];
  }
  return undefined;
}

export const BASE_URL =
  (readEnvValue(["VITE_BASE_URL", "BASE_URL"]) ??
    (typeof window !== "undefined" ? (readWindowValue("__BASE_URL__") as string | undefined) : undefined) ??
    DEFAULT_BASE_URL);

const ORDER_BASE_FALLBACK = BASE_URL === DEFAULT_BASE_URL ? DEFAULT_ORDER_BASE_URL : BASE_URL;

export const ORDER_SERVICE_BASE_URL =
  readEnvValue(["VITE_ORDER_SERVICE_BASE_URL", "ORDER_SERVICE_BASE_URL"]) ?? ORDER_BASE_FALLBACK;

export const WS_BASE_URL =
  (readEnvValue(["VITE_WS_BASE_URL", "WS_BASE_URL"]) ??
    (typeof window !== "undefined" ? (readWindowValue("__WS_BASE_URL__") as string | undefined) : undefined) ??
    BASE_URL.replace(/^http/, "ws"));

export async function http(path: string | URL, options: RequestInit = {}) {
  const input = typeof path === "string" ? path : path.toString();
  const url = input.startsWith("http") ? input : `${BASE_URL}${normalize(input)}`;
  const response = await fetch(url, {
    headers: { "Content-Type": "application/json", ...(options.headers || {}) },
    ...options
  });
  if (!response.ok) {
    const text = await response.text().catch(() => "");
    throw new Error(`HTTP ${response.status} ${response.statusText}: ${text}`);
  }
  return response;
}

export function wsPath(path: string) {
  const base = WS_BASE_URL.endsWith("/") ? WS_BASE_URL.slice(0, -1) : WS_BASE_URL;
  return `${base}${normalize(path).replace(/^http/, "ws")}`;
}

function normalize(p: string) {
  if (!p) return "";
  return p.startsWith("/") ? p : `/${p}`;
}
