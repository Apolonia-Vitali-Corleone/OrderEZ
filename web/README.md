# OrderEZ 可视化测试面板

一个基于 React + Vite 构建的轻量级可视化工具，用于在缺乏自动化测试的场景下快速验证 OrderEZ 用户服务与订单服务的核心流程。

## 功能概览

- ✅ 自动注册临时账号或复用现有账号
- ✅ 登录并刷新 JWT，展示令牌预览
- ✅ 带分页参数的用户列表查询
- ✅ 使用配置化订单明细创建测试订单
- ✅ 汇总执行耗时、状态与响应详情

## 快速开始

```bash
cd web
npm install
npm run dev
```

Vite 会在本机的 5173 端口启动开发服务器。页面默认请求由 `src/api/client.ts` 中 `BASE_URL` 与 `WS_BASE_URL` 指定的地址。

> 如果你的后端部署在其它地址或端口，可在左侧配置面板或环境变量中修改。

## 覆盖 BASE_URL / WS_BASE_URL

所有发往本地服务的 HTTP/WebSocket 请求都会通过 `src/api/client.ts` 暴露的客户端完成。你可以通过以下方式覆盖默认地址：

- 构建或启动前设置环境变量：

  ```bash
  # HTTP 基础地址
  BASE_URL="https://your-api.example.com" npm run dev

  # 可选：独立覆盖订单服务或 WebSocket 地址
  ORDER_SERVICE_BASE_URL="https://orders.example.com" npm run dev
  WS_BASE_URL="wss://your-api.example.com" npm run dev
  ```

- 在部署的 HTML 中注入运行时变量（若需要在不重新构建的情况下调整）：

  ```html
  <script>
    window.__BASE_URL__ = window.__BASE_URL__ || "https://your-api.example.com";
    window.__WS_BASE_URL__ = window.__WS_BASE_URL__ || null;
  </script>
  ```

默认情况下，`BASE_URL` 指向 127.0.0.1 的 48482 端口，你仍可以在页面左侧的“服务地址”区域手动调整为其它值。

## 使用建议

1. 启动用户服务与订单服务，确保数据库、Redis、RabbitMQ 等依赖就绪。
2. 打开测试面板，根据需要调整服务地址、账号策略及订单明细。
3. 点击“**一键运行全部测试**”或针对某个接口单独执行。
4. 展开每个测试卡片下方的“查看响应详情”即可查看完整的返回数据。

## 构建产物

```bash
npm run build
```

构建后的静态资源会输出到 `dist/` 目录，可交由任意静态资源服务托管。

