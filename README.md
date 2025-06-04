# OrderEZ
《点餐爽》——开箱即用的点餐系统






用户对应一个商品池子，这个池子记录商品，对应的类别，以及一个分数（这个分数根据用户查看商品、收藏商品进行加权计算）。商品作为key是不能重复的。而且这个商品本身是有热度排行的。用户的id指向一个hash，hash的key是商品的id，value是商品的具体信息的key。商品具体信息也是一个hash，

string:		UserID:ProductPoolHasnID

hash:		ProductPool:ProductID,ProductDetailID

hash：		ProductDetail：ProductCategory，销量

sorted set:    ProductID,score



string:		UserID:ProductPoolHasnID  记录用户的商品池

hash:		ProductPool:ProductID,ProductDetailID  记录收藏的商品以及商品详细信息的id

hash：		ProductDetail：ProductCategory，销量    记录收藏的商品的类别和销量

sorted set:    ProductID,score      用户查看和收藏商品会加权计算分数，得到score。然后根据sorted set的排序推荐商品。







用户发布动态，每条动态有多个评论。评论包括评论内容、评论者、点赞数等信息。

每个动态的评论数和点赞数可能会变化，并且需要支持按时间顺序获取最新评论。

动态本身需要能够按发布时间、点赞数等进行排序。

还需要支持用户通过点赞数来过滤出热门动态。





```
用户{
	动态1{
		发布时间（需要排序）   
		点赞数（需要排序）
		点赞{
			点赞者1
			点赞者2...
		}
		评论数
		评论{
			评论1{
				评论时间（需要排序）
				评论者
				评论内容
			}
			评论2...
		}
	}
	动态2...
}
```





```
玩家{
	实时状态
	积分
	任务列表{
		任务1{
			完成情况
			对应积分
		}
		任务2...
	}
}

排行榜{
	第一名
	第二名...
}
```





```
消息队列{
	用户id
	消息{
		消息1{
			消息类型
			内容
			发送时间
			状态
		}
	}
}
```





```
活动{
	状态
	商品列表{
		商品1{
			库存
		}
		商品2...
	}
}
用户{
	秒杀商品列表{
		商品1{
			数量
		}
		商品2...
	}
}

统计{
	秒杀商品列表{
		商品1{
			秒杀人数
			购买者{
				用户1
				用户2...
			}
		}
		商品2...
	}
}
```









目录结构：

tset

​	-originm3

​	-am3

​	-bm3

在目录origin里

git init --bare



在目录tset

git clone originm3 am3

cd am3

创建user.go并且内容是：

type User struct {
ID   int
Name string
}

git add *

git commit -m "aaa commit user.go origin"

git push



修改user.go为：

type User struct {
ID   int
Name string
Age	 int
}

git add *

git commit -m "aaa commit user.go"

git push



在目录tset

git clone originm3 bm3

cd bm3

修改user.go为：

type User struct {
ID   int
Name string
Email string
}

git add *

git commit -m "aaa commit user.go"

git push

push发现出错

git pull

查看user.go发下如下内容：

```
type User struct {
    ID   int
    Name string
<<<<<<< HEAD
        Email string
=======
        Age      int
>>>>>>> e3d33ecb91aec52bdac8dddd022fca545543b56e
}
```

修改为

```
type User struct {
    ID   int
    Name string
	Age	 int
	Email string
}
```

git add *

git commit -m "aaa commit user.go"

git push

push成功

















```
git config --global core.editor "\"C:\\develop\\notepad++\\notepad++.exe\" -multiInst -notabbar -nosession -noPlugin"
```





```
service=new_service(mysql,redis,mq)
(new_service自动new_repo)

handler=new_handler(service)
```



# Cart

## /cart

### 前端请求

```http
GET /cart
Authorization：token
```

### CartHandler

```go
type CartHandler struct {
	cartService       *service.CartService
	cartDetailService *service.CartDetailService
}
```

## Method

```
GetCart(){
	userID = getUserIDByToken(token)
	cartID=CartService.GetCartIDByUserID(userid)
	cartDetailList=cartDetailService.GetCartDetailListByCartID(cartid)
	return cartDetailList
}
```

### CartService

```

```

### CartDetailService

```

```

### CartRepository

```

```

### CartDetailRepository

```
s,ts/.Hc/Gs$8Ln
```





```bash
docker run -dit --restart=always \
  --name dst-server \
  -v /root/dst-docker:/root/.klei/DoNotStarveTogether \
  -p 10999:10999/udp \
  -p 11000:11000/udp \
  -p 12346:12346 \
  swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/jamesits/dst-server:nightly

docker run -it \
  --name dst-server \
  -v /root/.klei/DoNotStarveTogether:/data \
  -p 10999-11000:10999-11000/udp \
  -p 12346-12347:12346-12347/udp \
  -e "DST_SERVER_ARCH=amd64"  \
  swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/jamesits/dst-server:nightly

docker pull swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/jamesits/dst-server:latest

docker inspect docker.m.daocloud.io/library/jamesits/dst-server


docker exec -it dst-server bash

steamcmd + login anonymous + app_update 343050 validate + quit

find / -type f -name "dontstarve_dedicated_server_nullrenderer_x64" 2>/dev/null

/opt/dst_server/bin64/dontstarve_dedicated_server_nullrenderer_x64 -cluster Cluster_4 -shard Master
```



```
find / -type f -name "main.lua" 2>/dev/null
```



```
/opt/dst_server/bin64/dontstarve_dedicated_server_nullrenderer_x64 -cluster Cluster_1 -shard Caves
```



```
find / -type f -name "dedicated_server_mods_setup.lua" 2>/dev/null
```





```shell

root@b9b4dd351c14:/opt/dst_server# find / -type f -name "dedicated_server_mods_setup.lua" 2>/dev/null
/opt/dst_default_config/DoNotStarveTogether/Cluster_1/mods/dedicated_server_mods_setup.lua
/data/DoNotStarveTogether/Cluster_1/mods/dedicated_server_mods_setup.lua

root@b9b4dd351c14:/opt/dst_server# find / -type f -name "dontstarve_dedicated_server_nullrenderer_x64" 2>/dev/null
/opt/dst_server/bin64/dontstarve_dedicated_server_nullrenderer_x64


```





```
docker tag \
  swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/jamesits/dst-server:nightly \
  dst-server:stable
```









非常赞！你这种**一步一步来、稳扎稳打**的方式非常适合服务器搭建 👍
下面是你所说步骤的**详细执行指南**，确保每一步都不出问题：

------

## ✅ 第一步：安装必要依赖（在 root 下执行）

```bash
apt update && apt upgrade -y
apt install software-properties-common -y
dpkg --add-architecture i386
apt update
apt install lib32gcc-s1 lib32stdc++6 libc6-i386 libncurses5:i386 libstdc++6:i386 wget screen -y
```

> 💡 `screen` 是后面运行 DST 服务常用的工具，提前装上。

------

## ✅ 第二步：创建 `steam` 用户并设置密码（root 下）

```bash
useradd -m steam -s /bin/bash
passwd steam
```

> 🔐 会提示你两次输入密码。设置一个你记得住的。密码：steam

------

## ✅ 第三步：切换到 steam 用户

```bash
su - steam
```

> 👤 此后你就在 `steam` 用户环境里了，所有路径都以 `~/` 为根。

------

## ✅ 第四步：安装 SteamCMD（非 root，使用 steam 用户）

```bash
mkdir ~/steamcmd && cd ~/steamcmd
wget https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz
tar -xvzf steamcmd_linux.tar.gz
cd ~
/home/steam/steamcmd/steamcmd.sh
```

> 🚀 会进入 SteamCMD 的终端界面，第一次运行时会更新。

------

## ✅ 第五步：安装 Don't Starve Together Dedicated Server

在 SteamCMD 命令行中输入：

```text
force_install_dir dstserver
login anonymous
app_update 343050 validate
quit
```

- `343050` 是 DST 专用服务器的 App ID
- 安装完成后会退出 `steamcmd`

------



```
mkdir -p /home/steam/.klei/DoNotStarveTogether/
cp -r /home/steam/backup/Cluster_4 /home/steam/.klei/DoNotStarveTogether/
chown -R steam:steam /home/steam/.klei/DoNotStarveTogether/Cluster_4
```

```
connect 47.117.155.156:10888


connect 47.117.155.156:10998 +password "97543007"
```





```
systemctl restart dst_master
journalctl -u dst_master -f

systemctl restart dst_caves
journalctl -u dst_caves -f
```







# 点餐系统后端架构设计

## 项目概述

这是一个基于Go语言的点餐系统后端设计，采用分布式微服务架构。项目使用go-micro作为微服务框架，gin作为HTTP框架，gorm作为ORM工具。

## 系统架构

系统采用微服务架构，主要包含以下服务:

1. **用户服务(User Service)** - 处理用户注册、登录、认证等
2. **菜品服务(Dish Service)** - 管理菜品信息、分类等
3. **订单服务(Order Service)** - 处理订单创建、支付、状态管理等
4. **购物车服务(Cart Service)** - 管理用户购物车
5. **评价服务(Review Service)** - 处理用户对菜品和服务的评价

## 技术栈

- **语言**: Golang
- **微服务框架**: go-micro
- **HTTP框架**: gin
- **ORM工具**: gorm
- **数据库**: MySQL
- **缓存**: Redis
- **消息队列**: RabbitMQ
- **服务发现**: Consul
- **配置中心**: go-micro config
- **链路追踪**: Jaeger

## 代码结构设计

每个微服务都遵循类似的代码结构:

```
service-name/
├── cmd/                 # 程序入口
├── internal/
│   ├── config/          # 配置相关
│   ├── handler/         # HTTP/gRPC 处理器
│   ├── model/           # 数据模型
│   ├── repository/      # 数据访问层
│   ├── service/         # 业务逻辑层
│   └── utils/           # 工具函数
├── proto/               # Proto 文件(用于服务间通信)
├── Dockerfile           # Docker 构建文件
├── go.mod               # Go 模块文件
└── go.sum               # 依赖版本锁定文件
```

## 服务详细设计

### 1. 用户服务 (User Service)

#### Model

```go
// internal/model/user.go
package model

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    gorm.Model
    Username     string `gorm:"size:50;not null;unique"`
    Password     string `gorm:"size:100;not null"`
    Nickname     string `gorm:"size:50"`
    Mobile       string `gorm:"size:20;index"`
    Email        string `gorm:"size:100;index"`
    Avatar       string `gorm:"size:255"`
    Gender       int    `gorm:"default:0"` // 0-未知 1-男 2-女
    Status       int    `gorm:"default:1"` // 0-禁用 1-正常
    LastLoginAt  *time.Time
    LastLoginIP  string `gorm:"size:50"`
}
```

#### Repository

```go
// internal/repository/user_repository.go
package repository

import (
    "context"
    "github.com/your-org/user-service/internal/model"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(ctx context.Context, user *model.User) error
    GetByID(ctx context.Context, id uint) (*model.User, error)
    GetByUsername(ctx context.Context, username string) (*model.User, error)
    Update(ctx context.Context, user *model.User) error
    Delete(ctx context.Context, id uint) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

// 实现各种方法...
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
    var user model.User
    if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// 其他方法实现
```

#### Service

```go
// internal/service/user_service.go
package service

import (
    "context"
    "errors"
    "golang.org/x/crypto/bcrypt"
    
    "github.com/your-org/user-service/internal/model"
    "github.com/your-org/user-service/internal/repository"
)

type UserService interface {
    Register(ctx context.Context, username, password, nickname, mobile, email string) (*model.User, error)
    Login(ctx context.Context, username, password string) (*model.User, error)
    GetUser(ctx context.Context, id uint) (*model.User, error)
    UpdateUser(ctx context.Context, user *model.User) error
    DeleteUser(ctx context.Context, id uint) error
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, username, password, nickname, mobile, email string) (*model.User, error) {
    // 检查用户名是否已存在
    existingUser, _ := s.userRepo.GetByUsername(ctx, username)
    if existingUser != nil {
        return nil, errors.New("username already exists")
    }
    
    // 密码加密
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    
    user := &model.User{
        Username: username,
        Password: string(hashedPassword),
        Nickname: nickname,
        Mobile:   mobile,
        Email:    email,
    }
    
    if err := s.userRepo.Create(ctx, user); err != nil {
        return nil, err
    }
    
    return user, nil
}

// 其他方法实现
```

#### Handler

```go
// internal/handler/user_handler.go
package handler

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/your-org/user-service/internal/service"
)

type UserHandler struct {
    userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
    var req struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
        Nickname string `json:"nickname"`
        Mobile   string `json:"mobile"`
        Email    string `json:"email"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user, err := h.userService.Register(c.Request.Context(), req.Username, req.Password, req.Nickname, req.Mobile, req.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "user registered successfully",
        "data": gin.H{
            "id":       user.ID,
            "username": user.Username,
            "nickname": user.Nickname,
        },
    })
}

// 其他处理方法
```

### 2. 菜品服务 (Dish Service)

#### Model

```go
// internal/model/dish.go
package model

import (
    "gorm.io/gorm"
)

type Category struct {
    gorm.Model
    Name        string `gorm:"size:50;not null"`
    Description string `gorm:"size:255"`
    Sort        int    `gorm:"default:0"`
    Status      int    `gorm:"default:1"` // 0-禁用 1-启用
}

type Dish struct {
    gorm.Model
    Name        string  `gorm:"size:100;not null"`
    CategoryID  uint    `gorm:"index"`
    Price       float64 `gorm:"type:decimal(10,2);not null"`
    OriginalPrice float64 `gorm:"type:decimal(10,2)"`
    Description string  `gorm:"size:500"`
    Image       string  `gorm:"size:255"`
    Status      int     `gorm:"default:1"` // 0-下架 1-上架
    Sort        int     `gorm:"default:0"`
    Sales       int     `gorm:"default:0"`
    Rating      float32 `gorm:"default:5.0"`
    Tags        string  `gorm:"size:255"` // 标签,逗号分隔
}
```

#### Repository

```go
// internal/repository/dish_repository.go
package repository

import (
    "context"
    "github.com/your-org/dish-service/internal/model"
    "gorm.io/gorm"
)

type DishRepository interface {
    Create(ctx context.Context, dish *model.Dish) error
    GetByID(ctx context.Context, id uint) (*model.Dish, error)
    ListByCategoryID(ctx context.Context, categoryID uint, page, pageSize int) ([]*model.Dish, int64, error)
    Update(ctx context.Context, dish *model.Dish) error
    Delete(ctx context.Context, id uint) error
}

// Repository实现...
```

### 3. 订单服务 (Order Service)

#### Model

```go
// internal/model/order.go
package model

import (
    "gorm.io/gorm"
    "time"
)

type Order struct {
    gorm.Model
    OrderNo     string    `gorm:"size:50;not null;unique"`
    UserID      uint      `gorm:"index;not null"`
    Amount      float64   `gorm:"type:decimal(10,2);not null"`
    Status      int       `gorm:"default:0"` // 0-待支付 1-已支付 2-已取消 3-已完成 4-已退款
    Address     string    `gorm:"size:255"`
    ContactName string    `gorm:"size:50"`
    ContactTel  string    `gorm:"size:20"`
    Remark      string    `gorm:"size:255"`
    PayMethod   int       `gorm:"default:1"` // 1-在线支付 2-货到付款
    PayTime     *time.Time
    DeliverTime *time.Time
    FinishTime  *time.Time
}

type OrderItem struct {
    gorm.Model
    OrderID   uint    `gorm:"index;not null"`
    DishID    uint    `gorm:"not null"`
    DishName  string  `gorm:"size:100;not null"`
    Price     float64 `gorm:"type:decimal(10,2);not null"`
    Quantity  int     `gorm:"not null"`
    Amount    float64 `gorm:"type:decimal(10,2);not null"`
}
```

#### Service

```go
// internal/service/order_service.go
package service

import (
    "context"
    "fmt"
    "time"
    
    "github.com/your-org/order-service/internal/model"
    "github.com/your-org/order-service/internal/repository"
)

type OrderService interface {
    CreateOrder(ctx context.Context, userID uint, items []OrderItemRequest, address, contactName, contactTel, remark string) (*model.Order, error)
    GetOrder(ctx context.Context, id uint) (*model.Order, []*model.OrderItem, error)
    ListUserOrders(ctx context.Context, userID uint, page, pageSize int) ([]*model.Order, int64, error)
    PayOrder(ctx context.Context, orderID uint, payMethod int) error
    CancelOrder(ctx context.Context, orderID uint) error
}

// 具体实现...
```

### 4. 购物车服务 (Cart Service)

#### Model

```go
// internal/model/cart.go
package model

import (
    "gorm.io/gorm"
)

type CartItem struct {
    gorm.Model
    UserID   uint    `gorm:"index;not null"`
    DishID   uint    `gorm:"not null"`
    Quantity int     `gorm:"default:1"`
    Selected bool    `gorm:"default:true"`
}
```

### 5. 评价服务 (Review Service)

#### Model

```go
// internal/model/review.go
package model

import (
    "gorm.io/gorm"
)

type Review struct {
    gorm.Model
    UserID      uint    `gorm:"index;not null"`
    DishID      uint    `gorm:"index;not null"`
    OrderID     uint    `gorm:"index"`
    Content     string  `gorm:"size:500"`
    Rating      int     `gorm:"not null;default:5"` // 1-5分
    Images      string  `gorm:"size:1000"` // 图片URL,逗号分隔
    ReplyContent string `gorm:"size:500"`
    ReplyTime    *time.Time
}
```

## 微服务集成

### Proto定义示例 (Order Service)

```protobuf
// proto/order/order.proto
syntax = "proto3";

package order;

option go_package = "github.com/your-org/order-service/proto/order";

service OrderService {
    rpc CreateOrder (CreateOrderRequest) returns (OrderResponse) {}
    rpc GetOrder (GetOrderRequest) returns (OrderDetailResponse) {}
    rpc ListUserOrders (ListUserOrdersRequest) returns (ListOrdersResponse) {}
    rpc PayOrder (PayOrderRequest) returns (EmptyResponse) {}
    rpc CancelOrder (CancelOrderRequest) returns (EmptyResponse) {}
}

message CreateOrderRequest {
    uint32 user_id = 1;
    repeated OrderItem items = 2;
    string address = 3;
    string contact_name = 4;
    string contact_tel = 5;
    string remark = 6;
}

message OrderItem {
    uint32 dish_id = 1;
    int32 quantity = 2;
}

// 其他消息定义...
```

### 服务注册与发现

```go
// cmd/main.go
package main

import (
    "log"
    
    "github.com/gin-gonic/gin"
    "github.com/go-micro/plugins/v4/registry/consul"
    "go-micro.dev/v4"
    "go-micro.dev/v4/registry"
    
    "github.com/your-org/order-service/internal/config"
    "github.com/your-org/order-service/internal/handler"
    "github.com/your-org/order-service/internal/repository"
    "github.com/your-org/order-service/internal/service"
    orderProto "github.com/your-org/order-service/proto/order"
)

func main() {
    // 加载配置
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // 设置Consul注册中心
    reg := consul.NewRegistry(
        registry.Addrs(cfg.ConsulAddress),
    )
    
    // 创建服务
    svc := micro.NewService(
        micro.Name("order.service"),
        micro.Version("latest"),
        micro.Registry(reg),
    )
    
    // 初始化服务
    svc.Init()
    
    // 数据库初始化
    db, err := repository.InitDB(cfg.Database)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    
    // 依赖注入
    orderRepo := repository.NewOrderRepository(db)
    orderItemRepo := repository.NewOrderItemRepository(db)
    orderService := service.NewOrderService(orderRepo, orderItemRepo)
    
    // 注册Handler
    if err := orderProto.RegisterOrderServiceHandler(svc.Server(), handler.NewOrderHandler(orderService)); err != nil {
        log.Fatalf("Failed to register handler: %v", err)
    }
    
    // HTTP服务
    router := gin.Default()
    httpHandler := handler.NewHTTPHandler(orderService)
    httpHandler.RegisterRoutes(router)
    
    // 启动HTTP服务
    go func() {
        if err := router.Run(cfg.HTTPAddress); err != nil {
            log.Fatalf("Failed to run HTTP server: %v", err)
        }
    }()
    
    // 启动RPC服务
    if err := svc.Run(); err != nil {
        log.Fatalf("Failed to run service: %v", err)
    }
}
```

## API 设计

### RESTful API 示例 (订单服务)

| 方法   | 路径                     | 描述               |
|------|--------------------------|-------------------|
| POST | /api/orders              | 创建新订单          |
| GET  | /api/orders/:id          | 获取订单详情        |
| GET  | /api/orders              | 获取用户所有订单     |
| PUT  | /api/orders/:id/pay      | 支付订单           |
| PUT  | /api/orders/:id/cancel   | 取消订单           |

### API 处理示例

```go
// internal/handler/http_handler.go
package handler

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "github.com/your-org/order-service/internal/service"
)

type HTTPHandler struct {
    orderService service.OrderService
}

func NewHTTPHandler(orderService service.OrderService) *HTTPHandler {
    return &HTTPHandler{orderService: orderService}
}

func (h *HTTPHandler) RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        orders := api.Group("/orders")
        {
            orders.POST("", h.CreateOrder)
            orders.GET("/:id", h.GetOrder)
            orders.GET("", h.ListUserOrders)  
            orders.PUT("/:id/pay", h.PayOrder)
            orders.PUT("/:id/cancel", h.CancelOrder)
        }
    }
}

func (h *HTTPHandler) CreateOrder(c *gin.Context) {
    var req struct {
        Items       []struct {
            DishID   uint `json:"dishId" binding:"required"`
            Quantity int  `json:"quantity" binding:"required,min=1"`
        } `json:"items" binding:"required,dive"`
        Address     string `json:"address" binding:"required"`
        ContactName string `json:"contactName" binding:"required"`
        ContactTel  string `json:"contactTel" binding:"required"`
        Remark      string `json:"remark"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 从JWT token中获取用户ID
    userID := getUserIDFromContext(c)
    
    // 转换请求格式
    items := make([]service.OrderItemRequest, 0, len(req.Items))
    for _, item := range req.Items {
        items = append(items, service.OrderItemRequest{
            DishID:   item.DishID,
            Quantity: item.Quantity,
        })
    }
    
    order, err := h.orderService.CreateOrder(c.Request.Context(), userID, items, req.Address, req.ContactName, req.ContactTel, req.Remark)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "order created successfully",
        "data": order,
    })
}

// 其他处理方法...
```

## 中间件设计

### JWT认证中间件

```go
// internal/middleware/jwt.go
package middleware

import (
    "net/http"
    "strings"
    
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

func JWTAuth(secretKey string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
            c.Abort()
            return
        }
        
        parts := strings.SplitN(authHeader, " ", 2)
        if !(len(parts) == 2 && parts[0] == "Bearer") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
            c.Abort()
            return
        }
        
        token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
            return []byte(secretKey), nil
        })
        
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }
        
        // 提取claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to parse token claims"})
            c.Abort()
            return
        }
        
        // 设置用户ID到上下文
        userID := uint(claims["user_id"].(float64))
        c.Set("userID", userID)
        
        c.Next()
    }
}
```

## 数据库设计

主要包括以下几个表:

1. `users` - 用户信息表
2. `categories` - 菜品分类表
3. `dishes` - 菜品信息表
4. `orders` - 订单主表
5. `order_items` - 订单明细表
6. `cart_items` - 购物车表
7. `reviews` - 评价表

## 缓存设计

使用Redis缓存以下数据:

1. 菜品分类列表
2. 热门菜品
3. 用户购物车
4. 用户会话信息

```go
// internal/cache/dish_cache.go
package cache

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
    
    "github.com/go-redis/redis/v8"
    "github.com/your-org/dish-service/internal/model"
)

type DishCache interface {
    SetHotDishes(ctx context.Context, dishes []*model.Dish) error
    GetHotDishes(ctx context.Context) ([]*model.Dish, error)
}

type dishCache struct {
    rdb *redis.Client
}

func NewDishCache(rdb *redis.Client) DishCache {
    return &dishCache{rdb: rdb}
}

func (c *dishCache) SetHotDishes(ctx context.Context, dishes []*model.Dish) error {
    data, err := json.Marshal(dishes)
    if err != nil {
        return err
    }
    
    return c.rdb.Set(ctx, "hot_dishes", data, 30*time.Minute).Err()
}

func (c *dishCache) GetHotDishes(ctx context.Context) ([]*model.Dish, error) {
    data, err := c.rdb.Get(ctx, "hot_dishes").Bytes()
    if err != nil {
        if err == redis.Nil {
            return nil, nil
        }
        return nil, err
    }
    
    var dishes []*model.Dish
    err = json.Unmarshal(data, &dishes)
    return dishes, err
}
```

## 消息队列设计

使用RabbitMQ处理以下异步任务:

1. 订单状态变更通知
2. 库存更新
3. 积分变更
4. 推送通知

```go
// internal/mq/order_publisher.go
package mq

import (
    "context"
    "encoding/json"
    
    "github.com/streadway/amqp"
    "github.com/your-org/order-service/internal/model"
)

type OrderStatusChangedEvent struct {
    OrderID    uint   `json:"orderId"`
    UserID     uint   `json:"userId"`
    OldStatus  int    `json:"oldStatus"`
    NewStatus  int    `json:"newStatus"`
    OrderNo    string `json:"orderNo"`
    Amount     float64 `json:"amount"`
    UpdateTime int64  `json:"updateTime"`
}

type OrderPublisher interface {
    PublishOrderStatusChanged(ctx context.Context, event OrderStatusChangedEvent) error
}

type orderPublisher struct {
    channel *amqp.Channel
}

func NewOrderPublisher(channel *amqp.Channel) (OrderPublisher, error) {
    // 声明交换机
    err := channel.ExchangeDeclare(
        "order_events", // 交换机名称
        "topic",        // 交换机类型
        true,           // 持久化
        false,          // 自动删除
        false,          // 内部交换机
        false,          // no-wait
        nil,            // 参数
    )
    if err != nil {
        return nil, err
    }
    
    return &orderPublisher{channel: channel}, nil
}

func (p *orderPublisher) PublishOrderStatusChanged(ctx context.Context, event OrderStatusChangedEvent) error {
    data, err := json.Marshal(event)
    if err != nil {
        return err
    }
    
    return p.channel.Publish(
        "order_events",             // 交换机
        "order.status.changed",     // 路由键
        false,                      // mandatory
        false,                      // immediate
        amqp.Publishing{
            ContentType:  "application/json",
            Body:         data,
            DeliveryMode: amqp.Persistent,
        },
    )
}
```

## 配置管理

使用go-micro的配置管理功能:

```go
// internal/config/config.go
package config

import (
    "github.com/go-micro/plugins/v4/config/source/consul"
    "go-micro.dev/v4/config"
)

type Config struct {
    Service struct {
        Name    string `json:"name"`
        Version string `json:"version"`
    } `json:"service"`
    
    Database struct {
        Driver   string `json:"driver"`
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Username string `json:"username"`
        Password string `json:"password"`
        DBName   string `json:"dbname"`
    } `json:"database"`
    
    Redis struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Password string `json:"password"`
        DB       int    `json:"db"`
    } `json:"redis"`
    
    RabbitMQ struct {
        URL string `json:"url"`
    } `json:"rabbitmq"`
    
    ConsulAddress string `json:"consulAddress"`
    HTTPAddress   string `json:"httpAddress"`
    JWTSecret     string `json:"jwtSecret"`
}

func Load() (*Config, error) {
    // 创建配置源
    consulSource := consul.NewSource(
        consul.WithAddress("localhost:8500"),
        consul.WithPrefix("/config/order-service"),
        consul.StripPrefix(true),
    )
    
    // 创建配置
    conf := config.NewConfig()
    
    // 加载配置
    if err := conf.Load(consulSource); err != nil {
        return nil, err
    }
    
    var cfg Config
    if err := conf.Scan(&cfg); err != nil {
        return nil, err
    }
    
    return &cfg, nil
}
```

## 链路追踪

使用Jaeger进行分布式链路追踪:

```go
// internal/tracer/tracer.go
package tracer

import (
    "io"
    
    "github.com/opentracing/opentracing-go"
    "github.com/uber/jaeger-client-go"
    jaegercfg "github.com/uber/jaeger-client-go/config"
)

// InitTracer 初始化Jaeger追踪器
func InitTracer(serviceName string, jaegerAgentHost string) (opentracing.Tracer, io.Closer, error) {
    cfg := jaegercfg.Configuration{
        ServiceName: serviceName,
        Sampler: &jaegercfg.SamplerConfig{
            Type:  jaeger.SamplerTypeConst,
            Param: 1,
        },
        Reporter: &jaegercfg.ReporterConfig{
            LogSpans:           true,
            LocalAgentHostPort: jaegerAgentHost,
        },
    }
    
    return cfg.NewTracer()
}
```

## 部署架构

系统可以部署在Kubernetes集群中,每个微服务都有独立的部署配置:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: order-service
  template:
    metadata:
      labels:
        app: order-service
    spec:
      containers:
      - name: order-service
        image: your-registry/order-service:latest
        ports:
        - containerPort: 8080
        envFrom:
        - configMapRef:
            name: order-service-config
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        resources:
          limits:
            cpu: "500m"
            memory: "512Mi"
          requests:
            cpu: "100m"
            memory: "128Mi"
        volumeMounts:
        - name: config-volume
          mountPath: /app/config
      volumes:
      - name: config-volume
        configMap:
          name: order-service-config
---
apiVersion: v1
kind: Service
metadata:
  name: order-service
spec:
  selector:
    app: order-service
  ports:
  - port: 80
    targetPort: 8080
  type: ClusterIP
```

## 安全设计

### 数据加密

1. 用户密码使用bcrypt算法加密存储
2. 敏感数据传输使用HTTPS
3. API接口使用JWT进行认证授权

```go
// internal/utils/password.go
package utils

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword 对密码进行加密
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash 验证密码
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

### JWT认证

```go
// internal/utils/jwt.go
package utils

import (
    "time"
    
    "github.com/golang-jwt/jwt/v4"
)

// GenerateToken 生成JWT token
func GenerateToken(userID uint, username string, secret string, expireDuration time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "user_id":  userID,
        "username": username,
        "exp":      time.Now().Add(expireDuration).Unix(),
        "iat":      time.Now().Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

// ParseToken 解析JWT token
func ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, jwt.ErrSignatureInvalid
}
```

### 权限控制

```go
// internal/middleware/rbac.go
package middleware

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

const (
    RoleCustomer = 1
    RoleStaff    = 2
    RoleAdmin    = 3
)

// RequireRole 检查用户角色
func RequireRole(requiredRole int) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("userRole")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }
        
        role := userRole.(int)
        if role < requiredRole {
            c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

## 错误处理

统一错误处理机制:

```go
// internal/utils/error.go
package utils

import (
    "fmt"
    "net/http"
)

// AppError 应用错误结构
type AppError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Err     error  `json:"-"`
}

func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

// NewAppError 创建新的应用错误
func NewAppError(code int, message string, err error) *AppError {
    return &AppError{
        Code:    code,
        Message: message,
        Err:     err,
    }
}

// ErrorMiddleware 统一错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        if len(c.Errors) > 0 {
            for _, e := range c.Errors {
                // 处理自定义错误
                if appErr, ok := e.Err.(*AppError); ok {
                    c.JSON(appErr.Code, gin.H{
                        "error": appErr.Message,
                    })
                    return
                }
            }
            
            // 处理其他错误
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Internal Server Error",
            })
        }
    }
}
```

### 错误使用示例

```go
// 服务层使用错误
func (s *userService) Login(ctx context.Context, username, password string) (*model.User, error) {
    user, err := s.userRepo.GetByUsername(ctx, username)
    if err != nil {
        return nil, utils.NewAppError(http.StatusInternalServerError, "Failed to get user", err)
    }
    
    if user == nil {
        return nil, utils.NewAppError(http.StatusUnauthorized, "Invalid username or password", nil)
    }
    
    if !utils.CheckPasswordHash(password, user.Password) {
        return nil, utils.NewAppError(http.StatusUnauthorized, "Invalid username or password", nil)
    }
    
    return user, nil
}

// 控制器层使用错误处理
func (h *UserHandler) Login(c *gin.Context) {
    var req struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.Error(utils.NewAppError(http.StatusBadRequest, "Invalid request", err))
        return
    }
    
    user, err := h.userService.Login(c.Request.Context(), req.Username, req.Password)
    if err != nil {
        c.Error(err)
        return
    }
    
    // 生成JWT token
    token, err := utils.GenerateToken(user.ID, user.Username, h.jwtSecret, time.Hour*24*7)
    if err != nil {
        c.Error(utils.NewAppError(http.StatusInternalServerError, "Failed to generate token", err))
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user": gin.H{
            "id":       user.ID,
            "username": user.Username,
            "nickname": user.Nickname,
        },
    })
}
```

## 日志管理

使用结构化日志:

```go
// internal/logger/logger.go
package logger

import (
    "os"
    
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger 初始化日志
func InitLogger(serviceName, level string) error {
    // 解析日志级别
    var logLevel zapcore.Level
    if err := logLevel.UnmarshalText([]byte(level)); err != nil {
        logLevel = zapcore.InfoLevel
    }
    
    // 创建编码器配置
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.TimeKey = "timestamp"
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    
    // 创建编码器
    consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)
    
    // 设置输出
    core := zapcore.NewCore(
        consoleEncoder,
        zapcore.AddSync(os.Stdout),
        logLevel,
    )
    
    // 创建日志
    Log = zap.New(core, zap.AddCaller(), zap.Fields(
        zap.String("service", serviceName),
    ))
    
    return nil
}
```

### 日志使用示例

```go
// 初始化日志
if err := logger.InitLogger("order-service", "info"); err != nil {
    log.Fatalf("Failed to initialize logger: %v", err)
}

// 使用日志
logger.Log.Info("Order created",
    zap.Uint("orderID", order.ID),
    zap.Uint("userID", order.UserID),
    zap.String("orderNo", order.OrderNo),
    zap.Float64("amount", order.Amount),
)

// 记录错误
logger.Log.Error("Failed to create order",
    zap.Uint("userID", userID),
    zap.Error(err),
)
```

## 服务监控

使用Prometheus和Grafana进行监控:

```go
// internal/metrics/metrics.go
package metrics

import (
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "time"
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "path"},
    )
)

// InitMetrics 初始化指标
func InitMetrics() {
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

// MetricsMiddleware Gin中间件收集HTTP指标
func MetricsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        c.Next()
        
        status := c.Writer.Status()
        path := c.FullPath()
        method := c.Request.Method
        
        httpRequestsTotal.WithLabelValues(method, path, string(rune(status))).Inc()
        httpRequestDuration.WithLabelValues(method, path).Observe(time.Since(start).Seconds())
    }
}

// RegisterMetricsEndpoint 注册/metrics端点
func RegisterMetricsEndpoint(r *gin.Engine) {
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
```

## 优雅关闭

处理服务优雅关闭:

```go
// cmd/main.go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    
    "github.com/gin-gonic/gin"
    "go-micro.dev/v4"
)

func main() {
    // ... 初始化代码 ...
    
    // 创建HTTP服务器
    srv := &http.Server{
        Addr:    cfg.HTTPAddress,
        Handler: router,
    }
    
    // 启动HTTP服务
    go func() {
        logger.Log.Info("Starting HTTP server", zap.String("address", cfg.HTTPAddress))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Log.Fatal("Failed to start HTTP server", zap.Error(err))
        }
    }()
    
    // 启动微服务
    go func() {
        logger.Log.Info("Starting RPC server")
        if err := service.Run(); err != nil {
            logger.Log.Fatal("Failed to start RPC server", zap.Error(err))
        }
    }()
    
    // 等待中断信号优雅关闭
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    logger.Log.Info("Shutting down servers...")
    
    // 设置超时上下文
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // 关闭HTTP服务器
    if err := srv.Shutdown(ctx); err != nil {
        logger.Log.Fatal("HTTP server shutdown failed", zap.Error(err))
    }
    
    // 关闭微服务
    if err := service.Stop(); err != nil {
        logger.Log.Fatal("RPC server shutdown failed", zap.Error(err))
    }
    
    logger.Log.Info("Servers gracefully stopped")
}
```

## 总结

本文档详细描述了一个基于Go语言的点餐系统后端架构设计，采用了现代化的微服务架构，使用go-micro作为微服务框架，gin作为HTTP框架，gorm作为ORM工具。系统分为用户服务、菜品服务、订单服务、购物车服务和评价服务五个主要的微服务。

系统实现了以下特点:

1. **微服务架构**: 使用go-micro实现服务注册、发现、负载均衡等功能
2. **RESTful API**: 使用gin框架实现HTTP API
3. **数据持久化**: 使用gorm操作MySQL数据库
4. **缓存**: 使用Redis进行数据缓存
5. **消息队列**: 使用RabbitMQ处理异步任务
6. **链路追踪**: 使用Jaeger进行分布式追踪
7. **配置管理**: 使用Consul进行配置管理
8. **监控**: 使用Prometheus和Grafana进行系统监控
9. **安全**: 实现了JWT认证和RBAC权限控制
10. **容器化部署**: 提供了Kubernetes部署配置

该架构具有良好的可扩展性和可维护性，适合中小型餐饮企业使用。
