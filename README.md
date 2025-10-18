# OrderEZ 后端服务

OrderEZ 是一个围绕餐饮点餐流程的多服务后端。项目包含独立的用户服务与订单服务，并提供商品秒杀、购物车等扩展模块所需的基础设施（MySQL、Redis、RabbitMQ）。代码基于 **Golang + Gin + GORM** 架构，支持 JWT 登录认证、RabbitMQ 异步下单、Redis 预减库存等场景。

## 目录结构

```
.
├── docker-compose.yml           # 本地依赖环境（MySQL、Redis、RabbitMQ）
├── deploy/local/mysql/init.sql  # MySQL 初始化脚本
├── order-service/               # 订单微服务
├── user-service/                # 用户微服务
└── Web/                         # 可视化测试面板（React + Vite，可选）
```

## 运行前准备

| 组件      | 版本/说明                          | 默认端口 |
|-----------|------------------------------------|----------|
| Go        | 1.23+（模块声明在 go.mod 中）       | -        |
| MySQL     | 8.0，默认账号 `root/123`           | 3306     |
| Redis     | 7.x，无密码                         | 6379     |
| RabbitMQ  | 3.x，默认账号 `guest/guest`        | 5672/15672 |

> 如果你安装了 Docker 与 Docker Compose，可以直接使用项目提供的 `docker-compose.yml` 快速拉起所有依赖服务。

```bash
# 1. 安装 Go 依赖
go mod download

# 2. 启动依赖服务（后台运行）
docker compose up -d

# 3. 首次启动会自动初始化数据库表，脚本位于 deploy/local/mysql/init.sql
```

数据库会初始化下列表结构：

- `oe_user`：存储用户账号（用户名唯一，密码为加密存储）。
- `oe_order`：存储订单主表（雪花算法生成的订单号，按用户唯一约束）。
- `oe_order_item`：存储订单明细（商品、数量、价格等）。

## 环境变量

服务在启动时会读取以下环境变量，未显式设置时将使用默认值：

| 变量名         | 说明                                  | 默认值 |
|----------------|---------------------------------------|--------|
| `MYSQL_DSN`    | MySQL 连接串                          | `root:123@tcp(127.0.0.1:3306)/order_ez?charset=utf8mb4&parseTime=True&loc=Local` |
| `REDIS_ADDR`   | Redis 地址                            | `127.0.0.1:6379` |
| `REDIS_PASSWORD` | Redis 密码                          | 空字符串 |
| `REDIS_DB`     | Redis DB 序号                         | `0` |
| `RABBITMQ_URL` | RabbitMQ 连接串                       | `amqp://guest:guest@127.0.0.1:5672/` |
| `JWT_SECRET`   | JWT 签名密钥（所有服务需保持一致）     | `order-ez-development-secret` |

> 本地开发无需修改即可使用 docker compose 提供的默认配置。

## 启动微服务

在依赖服务就绪后，可分别启动用户服务与订单服务：

```bash
# 启动用户服务（默认监听 127.0.0.1:48482）
cd user-service
go run ./...

# 启动订单服务（默认监听 127.0.0.1:48481）
cd ../order-service
go run ./...
```

服务启动时会输出监听地址，确认日志后即可通过 HTTP 接口访问。常用路由如下：

### 用户服务 API（`http://127.0.0.1:48482`）

| 方法 | 路径          | 说明         |
|------|---------------|--------------|
| POST | `/user/register` | 注册账号，返回 JWT |
| POST | `/user/login`    | 登录并返回 JWT |
| POST | `/user/logout`   | JWT 失效处理 |
| GET  | `/user/`         | 查询所有用户（需有效 JWT） |

### 订单服务 API（`http://127.0.0.1:48481`）

| 方法 | 路径     | 说明                 |
|------|----------|----------------------|
| POST | `/order/` | 创建订单并写入明细 |

订单创建请求示例：

```json
{
  "user_id": 10001,
  "items": [
    {"item_id": 2001, "item_name": "signature-noodles", "item_price": 32, "item_count": 2},
    {"item_id": 2005, "item_name": "soy-milk", "item_price": 8, "item_count": 1}
  ]
}
```

> 订单服务会使用 Redis 进行库存预减、使用 RabbitMQ 异步投递下单消息，并将最终结果写入 MySQL。请确保 Redis 与 RabbitMQ 均已启动。

## 运行测试

仓库内的所有 Go 单元测试可以直接运行：

```bash
go test ./...
cd order-service && go test ./...
cd ../user-service && go test ./...
```

## 可视化测试面板

如果你希望通过 Web 界面快速验证接口是否可用，可使用 `Web/` 目录下提供的 React 应用：

```bash
cd web
npm install
npm run dev
```

面板默认调用本地 `127.0.0.1` 上的用户服务与订单服务，你可以在左侧配置面板中修改地址、账号与测试订单明细，并按需执行各项测试。

## 常见问题

1. **数据库连接失败**：确认 MySQL 是否启动、端口是否被占用，或通过 `MYSQL_DSN` 调整连接串。
2. **JWT 鉴权失败**：确保所有服务使用相同的 `JWT_SECRET` 值，并在前端请求中通过 `Authorization: Bearer <token>` 传递。
3. **RabbitMQ 队列不存在**：代码会在发布消息时自动声明队列，若仍失败，请检查 `RABBITMQ_URL` 与服务网络是否互通。

