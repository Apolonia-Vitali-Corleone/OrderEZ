export function formatJson(input: unknown): string {
  try {
    if (typeof input === "string") {
      return JSON.stringify(JSON.parse(input), null, 2);
    }
    return JSON.stringify(input, null, 2);
  } catch (error) {
    return typeof input === "string" ? input : JSON.stringify(input);
  }
}

export function decodeJwt(token: string): Record<string, unknown> | null {
  try {
    const payload = token.split(".")[1];
    if (!payload) {
      return null;
    }
    const normalized = payload.replace(/-/g, "+").replace(/_/g, "/");
    const decoded = decodeURIComponent(
      atob(normalized)
        .split("")
        .map((char) => `%${char.charCodeAt(0).toString(16).padStart(2, "0")}`)
        .join("")
    );
    return JSON.parse(decoded);
  } catch (error) {
    return null;
  }
}

export function nowLabel(): string {
  const now = new Date();
  const pad = (value: number) => value.toString().padStart(2, "0");
  const date = `${now.getFullYear()}-${pad(now.getMonth() + 1)}-${pad(now.getDate())}`;
  const time = `${pad(now.getHours())}:${pad(now.getMinutes())}:${pad(now.getSeconds())}`;
  return `${date} ${time}`;
}
