import { VisualTest, TestResult } from "../types";
import "./TestCard.css";

interface TestCardProps {
  test: VisualTest;
  result: TestResult;
  onRun: () => void;
  disabled?: boolean;
}

const statusLabel: Record<TestResult["status"], string> = {
  idle: "等待执行",
  running: "正在执行",
  success: "执行成功",
  error: "执行失败"
};

export function TestCard({ test, result, onRun, disabled }: TestCardProps) {
  const status = result.status;
  return (
    <article className={`test-card test-card--${status}`}>
      <header className="test-card__header">
        <div>
          <h3>{test.title}</h3>
          <p>{test.description}</p>
        </div>
        <div className="test-card__meta">
          <span className={`status status--${status}`}>{statusLabel[status]}</span>
          <button className="ghost-button" onClick={onRun} disabled={disabled || status === "running"}>
            {status === "running" ? "执行中..." : "重新执行"}
          </button>
        </div>
      </header>
      <div className="test-card__body">
        {status === "idle" && <p className="test-card__placeholder">尚未执行</p>}
        {status !== "idle" && (
          <>
            <p className="test-card__message">{result.message}</p>
            <dl className="test-card__stats">
              {typeof result.durationMs === "number" && (
                <div>
                  <dt>耗时</dt>
                  <dd>{result.durationMs} ms</dd>
                </div>
              )}
              {result.finishedAt && (
                <div>
                  <dt>完成时间</dt>
                  <dd>{result.finishedAt}</dd>
                </div>
              )}
            </dl>
            {result.details && (
              <details>
                <summary>查看响应详情</summary>
                <pre>{result.details}</pre>
              </details>
            )}
          </>
        )}
      </div>
    </article>
  );
}
