export interface TestConfig {
  userServiceBase: string;
  orderServiceBase: string;
  defaultPassword: string;
  pageSize: number;
  useManualAccount: boolean;
  manualUsername: string;
  manualPassword: string;
  orderItemsText: string;
}

export interface SharedState {
  username?: string;
  password?: string;
  token?: string;
  tokenPreview?: string;
  userId?: number;
}

export interface TestOutcome {
  success: boolean;
  message: string;
  details?: string;
}

export interface TestHelpers {
  config: TestConfig;
  shared: SharedState;
  updateShared: (updates: Partial<SharedState>) => void;
}

export interface VisualTest {
  id: string;
  title: string;
  description: string;
  run: (helpers: TestHelpers) => Promise<TestOutcome>;
}

export type TestStatus = "idle" | "running" | "success" | "error";

export interface TestResult {
  status: TestStatus;
  message?: string;
  details?: string;
  durationMs?: number;
  finishedAt?: string;
}
