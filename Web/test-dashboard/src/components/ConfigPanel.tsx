import { ChangeEvent } from "react";
import { TestConfig } from "../types";
import "./ConfigPanel.css";

interface ConfigPanelProps {
  config: TestConfig;
  onChange: (config: TestConfig) => void;
  disabled?: boolean;
}

export function ConfigPanel({ config, onChange, disabled }: ConfigPanelProps) {
  const handleInput = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const target = event.target;
    const { name, value } = target;

    if (target instanceof HTMLInputElement && target.type === "checkbox") {
      onChange({ ...config, [name]: target.checked });
    } else if (name === "pageSize") {
      const parsed = Number.parseInt(value, 10);
      onChange({ ...config, pageSize: Number.isNaN(parsed) ? 1 : Math.max(parsed, 1) });
    } else {
      onChange({ ...config, [name]: value });
    }
  };

  return (
    <section className="panel">
      <header className="panel__header">
        <h2>测试环境配置</h2>
        <p>根据你的本地部署调整服务地址、账号以及测试订单数据。</p>
      </header>

      <div className="panel__group">
        <h3>服务地址</h3>
        <label className="panel__field">
          <span>用户服务基础地址</span>
          <input
            type="url"
            name="userServiceBase"
            value={config.userServiceBase}
            onChange={handleInput}
            disabled={disabled}
            placeholder="http://127.0.0.1:48482"
          />
        </label>
        <label className="panel__field">
          <span>订单服务基础地址</span>
          <input
            type="url"
            name="orderServiceBase"
            value={config.orderServiceBase}
            onChange={handleInput}
            disabled={disabled}
            placeholder="http://127.0.0.1:48481"
          />
        </label>
      </div>

      <div className="panel__group">
        <h3>账号策略</h3>
        <label className="panel__field">
          <span>自动注册密码</span>
          <input
            type="text"
            name="defaultPassword"
            value={config.defaultPassword}
            onChange={handleInput}
            disabled={disabled}
            placeholder="为自动注册的账号设置密码"
          />
        </label>

        <label className="panel__checkbox">
          <input
            type="checkbox"
            name="useManualAccount"
            checked={config.useManualAccount}
            onChange={handleInput}
            disabled={disabled}
          />
          <span>使用已有账号替换自动注册</span>
        </label>

        <label className="panel__field">
          <span>已有账号用户名</span>
          <input
            type="text"
            name="manualUsername"
            value={config.manualUsername}
            onChange={handleInput}
            disabled={disabled || !config.useManualAccount}
            placeholder="仅在勾选“使用已有账号”时生效"
          />
        </label>

        <label className="panel__field">
          <span>已有账号密码</span>
          <input
            type="password"
            name="manualPassword"
            value={config.manualPassword}
            onChange={handleInput}
            disabled={disabled || !config.useManualAccount}
            placeholder="输入已有账号的密码"
          />
        </label>
      </div>

      <div className="panel__group">
        <h3>测试订单明细</h3>
        <textarea
          name="orderItemsText"
          value={config.orderItemsText}
          onChange={handleInput}
          rows={10}
          disabled={disabled}
          spellCheck={false}
        />
        <p className="panel__hint">
          以上 JSON 会直接作为 <code>create_order_items</code> 字段发送给订单服务，你可以根据需要调整商品。字段包含
          <code>item_id</code>、<code>item_name</code>、<code>item_price</code> 与 <code>item_count</code>。
        </p>
      </div>

      <div className="panel__group panel__group--slim">
        <label className="panel__field">
          <span>用户列表分页大小</span>
          <input
            type="number"
            min={1}
            name="pageSize"
            value={config.pageSize}
            onChange={handleInput}
            disabled={disabled}
          />
        </label>
      </div>
    </section>
  );
}
