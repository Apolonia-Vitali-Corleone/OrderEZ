import { useMemo, useRef, useState } from "react";
import "./App.css";
import { ConfigPanel } from "./components/ConfigPanel";
import { TestCard } from "./components/TestCard";
import { decodeJwt, formatJson, nowLabel } from "./utils/format";
import { fetchJson } from "./utils/http";
import { SharedState, TestConfig, TestOutcome, TestResult, TestHelpers, VisualTest } from "./types";

const defaultOrderItems = `[
  {
    "item_id": 2001,
    "item_name": "signature-noodles",
    "item_price": 32,
    "item_count": 2
  },
  {
    "item_id": 2005,
    "item_name": "soy-milk",
    "item_price": 8,
    "item_count": 1
  }
]`;

const defaultConfig: TestConfig = {
  userServiceBase: "http://127.0.0.1:48482",
  orderServiceBase: "http://127.0.0.1:48481",
  defaultPassword: "OrderEz#123",
  pageSize: 10,
  useManualAccount: false,
  manualUsername: "",
  manualPassword: "",
  orderItemsText: defaultOrderItems
};

const tests: VisualTest[] = [
  {
    id: "register",
    title: "注册测试账号",
    description:
      "调用用户服务创建一个仅用于演示的临时账号，并立即获取登录令牌。若已启用手动账号，则跳过注册步骤。",
    run: async ({ config, updateShared }): Promise<TestOutcome> => {
      if (config.useManualAccount) {
        if (!config.manualUsername || !config.manualPassword) {
          return {
            success: false,
            message: "已选择使用已有账号，但用户名或密码为空，请在左侧填写完整。"
          };
        }
        updateShared({
          username: config.manualUsername,
          password: config.manualPassword,
          token: undefined,
          tokenPreview: undefined,
          userId: undefined
        });
        return {
          success: true,
          message: `已切换为使用已有账号 ${config.manualUsername}，后续测试将复用该账号。`,
          details: formatJson({ username: config.manualUsername })
        };
      }

      const username = `tester_${Date.now()}`;
      const response = await fetchJson(`${config.userServiceBase}/user/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          username,
          password: config.defaultPassword
        })
      });

      if (!response.ok) {
        const errorMessage = typeof response.json === "object" && response.json && "error" in response.json
          ? String((response.json as Record<string, unknown>).error)
          : response.statusText;
        return {
          success: false,
          message: `注册失败：${errorMessage}`,
          details: response.text || undefined
        };
      }

      const body = (response.json ?? {}) as Record<string, unknown>;
      const token = typeof body.token === "string" ? body.token : undefined;
      if (!token) {
        return {
          success: false,
          message: "注册成功但未获取到 token，请检查用户服务响应。",
          details: response.text || undefined
        };
      }

      const payload = decodeJwt(token) ?? {};
      const userIdValue = typeof payload.user_id === "number" ? payload.user_id : undefined;

      updateShared({
        username,
        password: config.defaultPassword,
        token,
        tokenPreview: `${token.slice(0, 12)}…${token.slice(-6)}`,
        userId: userIdValue
      });

      return {
        success: true,
        message: `已注册临时账号 ${username}`,
        details: response.text
      };
    }
  },
  {
    id: "login",
    title: "账号登录并刷新令牌",
    description: "使用指定账号调用 /user/login 接口，验证登录流程并刷新一次 JWT 令牌。",
    run: async ({ config, shared, updateShared }): Promise<TestOutcome> => {
      const username = config.useManualAccount ? config.manualUsername : shared.username;
      const password = config.useManualAccount ? config.manualPassword : shared.password;

      if (!username || !password) {
        return {
          success: false,
          message: "无法登录：缺少账号或密码。请先执行注册测试，或在左侧填写已有账号信息。"
        };
      }

      const response = await fetchJson(`${config.userServiceBase}/user/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
      });

      if (!response.ok) {
        const message = typeof response.json === "object" && response.json && "error" in response.json
          ? String((response.json as Record<string, unknown>).error)
          : response.statusText;
        return {
          success: false,
          message: `登录失败：${message}`,
          details: response.text || undefined
        };
      }

      const body = (response.json ?? {}) as Record<string, unknown>;
      const token = typeof body.token === "string" ? body.token : undefined;
      if (!token) {
        return {
          success: false,
          message: "登录成功但响应缺少 token。",
          details: response.text || undefined
        };
      }

      const payload = decodeJwt(token) ?? {};
      const userIdValue = typeof payload.user_id === "number" ? payload.user_id : shared.userId;

      updateShared({
        username,
        password,
        token,
        tokenPreview: `${token.slice(0, 12)}…${token.slice(-6)}`,
        userId: userIdValue
      });

      return {
        success: true,
        message: `登录成功，已刷新 JWT（预览：${token.slice(0, 12)}…）`,
        details: response.text
      };
    }
  },
  {
    id: "list-users",
    title: "查询用户列表",
    description: "访问 /user/ 列表接口，验证分页参数与 JWT 鉴权（若提供）是否生效。",
    run: async ({ config, shared }): Promise<TestOutcome> => {
      const url = new URL("/user/", config.userServiceBase);
      url.searchParams.set("page", "1");
      url.searchParams.set("pageSize", String(config.pageSize));

      const headers: Record<string, string> = {};
      if (shared.token) {
        headers.Authorization = `Bearer ${shared.token}`;
      }

      const response = await fetchJson(url.toString(), {
        method: "GET",
        headers
      });

      if (!response.ok) {
        const message = typeof response.json === "object" && response.json && "error" in response.json
          ? String((response.json as Record<string, unknown>).error)
          : response.statusText;
        return {
          success: false,
          message: `查询失败：${message}`,
          details: response.text || undefined
        };
      }

      const body = (response.json ?? {}) as Record<string, unknown>;
      const users = Array.isArray(body.users) ? body.users : [];

      return {
        success: true,
        message: `成功获取 ${users.length} 条用户数据。`,
        details: response.text
      };
    }
  },
  {
    id: "create-order",
    title: "创建示例订单",
    description: "调用 /order/ 接口创建一笔测试订单，验证 JWT、库存与订单写入流程。",
    run: async ({ config, shared }): Promise<TestOutcome> => {
      if (!shared.token) {
        return {
          success: false,
          message: "缺少登录令牌，无法创建订单。请先执行登录测试。"
        };
      }

      let items: unknown;
      try {
        items = JSON.parse(config.orderItemsText);
      } catch (error) {
        return {
          success: false,
          message: "订单明细 JSON 解析失败，请检查格式是否正确。"
        };
      }

      if (!Array.isArray(items) || items.length === 0) {
        return {
          success: false,
          message: "订单明细必须是非空数组。"
        };
      }

      const response = await fetchJson(`${config.orderServiceBase}/order/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${shared.token}`
        },
        body: JSON.stringify({
          user_id: shared.userId ?? 0,
          create_order_items: items
        })
      });

      if (!response.ok) {
        const message = typeof response.json === "object" && response.json && "error" in response.json
          ? String((response.json as Record<string, unknown>).error)
          : response.statusText;
        return {
          success: false,
          message: `创建订单失败：${message}`,
          details: response.text || undefined
        };
      }

      return {
        success: true,
        message: "订单创建成功，响应结果已记录。",
        details: response.text
      };
    }
  }
];

function createInitialResults(): Record<string, TestResult> {
  return tests.reduce<Record<string, TestResult>>((acc, test) => {
    acc[test.id] = { status: "idle" };
    return acc;
  }, {});
}

export default function App() {
  const [config, setConfig] = useState<TestConfig>(defaultConfig);
  const [sharedState, setSharedState] = useState<SharedState>({});
  const sharedRef = useRef<SharedState>({});
  const [results, setResults] = useState<Record<string, TestResult>>(createInitialResults);
  const [isRunningAll, setIsRunningAll] = useState(false);

  const updateShared = (updates: Partial<SharedState>) => {
    setSharedState((prev) => {
      const next = { ...prev, ...updates };
      sharedRef.current = next;
      return next;
    });
  };

  const runTest = async (test: VisualTest) => {
    setResults((prev) => ({
      ...prev,
      [test.id]: { status: "running", message: "正在执行..." }
    }));

    const startedAt = performance.now();
    try {
      const helpers: TestHelpers = {
        config,
        shared: sharedRef.current,
        updateShared
      };
      const outcome = await test.run(helpers);
      const duration = Math.round(performance.now() - startedAt);
      const result: TestResult = {
        status: outcome.success ? "success" : "error",
        message: outcome.message,
        details: outcome.details,
        durationMs: duration,
        finishedAt: nowLabel()
      };
      setResults((prev) => ({ ...prev, [test.id]: result }));
    } catch (error) {
      const duration = Math.round(performance.now() - startedAt);
      setResults((prev) => ({
        ...prev,
        [test.id]: {
          status: "error",
          message: error instanceof Error ? error.message : "执行过程中出现未知错误。",
          details: error instanceof Error ? error.stack : undefined,
          durationMs: duration,
          finishedAt: nowLabel()
        }
      }));
    }
  };

  const runAll = async () => {
    setIsRunningAll(true);
    for (const test of tests) {
      // eslint-disable-next-line no-await-in-loop
      await runTest(test);
      if (!sharedRef.current.token && test.id === "login") {
        break;
      }
    }
    setIsRunningAll(false);
  };

  const resetResults = () => {
    setResults(createInitialResults());
  };

  const summary = useMemo(() => {
    const values = Object.values(results);
    const success = values.filter((item) => item.status === "success").length;
    const failed = values.filter((item) => item.status === "error").length;
    const running = values.filter((item) => item.status === "running").length;
    return { total: tests.length, success, failed, running };
  }, [results]);

  return (
    <div className="app-shell">
      <header className="app-header">
        <h1>OrderEZ 服务可视化测试面板</h1>
        <p>
          该面板提供一套开箱即用的可视化测试流程，帮助你在没有单元测试的情况下快速验证用户服务与订单服务是否按预期运作。
          配置环境参数后，可单独执行或一键运行所有测试。
        </p>
      </header>

      <section className="summary-card">
        <h2>执行总览</h2>
        <div className="summary-metrics">
          <div className="summary-metric">
            <h3>全部测试</h3>
            <p>{summary.total}</p>
          </div>
          <div className="summary-metric">
            <h3>已通过</h3>
            <p>{summary.success}</p>
          </div>
          <div className="summary-metric">
            <h3>失败</h3>
            <p>{summary.failed}</p>
          </div>
          <div className="summary-metric">
            <h3>运行中</h3>
            <p>{summary.running}</p>
          </div>
        </div>
      </section>

      <div className="actions-row">
        <button className="primary-button" onClick={runAll} disabled={isRunningAll}>
          {isRunningAll ? "正在执行所有测试..." : "一键运行全部测试"}
        </button>
        <button className="ghost-button" onClick={resetResults} disabled={isRunningAll}>
          重置结果
        </button>
        {sharedState.tokenPreview && (
          <span>
            当前令牌：<code>{sharedState.tokenPreview}</code>
          </span>
        )}
      </div>

      <div className="layout">
        <ConfigPanel config={config} onChange={setConfig} disabled={isRunningAll} />

        <div className="tests-grid">
          {tests.map((test) => (
            <TestCard
              key={test.id}
              test={test}
              result={results[test.id]}
              onRun={() => runTest(test)}
              disabled={isRunningAll}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
