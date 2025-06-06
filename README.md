

# user

```

```



# order

# 创建订单以及对应的订单物品

## json

```json
// 大量商品订单
{
  "user_id": 11111,
  "items": [
    {
      "item_id": 1003,
      "item_name": "fuck"
      "item_price": 10,
      "item_count": 5
    },
    {
      "item_id": 1003,
      "item_name": "fuck"
      "item_price": 10,
      "item_count": 5
    },
    {
      "item_id": 1003,
      "item_name": "fuck"
      "item_price": 10,
      "item_count": 5
    },
  ]
}
```

## CreateOrderRequest

```
UserID
OrderItem[]
```

## CreateOrderItem

```
ItemID
ItemName
ItemPrice
ItemCount
```

handler接收到CreateOrderRequest，然后分出UserID和OrderItem[]

雪花算法一个OrderID，userid，然后通过OrderItem的price*count，传递给OrderService

## OrderService

OrderService调用OrderRepository将这三个参数组装成一个Order，然后直接存入，返回是否err

## Order

```go
OrderID
UserID
TotalPrice#
```

OrderItemService需要OrderItem[],然后handler把这个OrderItem[]传递给OrderItemService。

OrderItemService遍历这个OrderItem[]然后依次存储进数据库，返回是否err







# cart

```

```



# cart_item

```

```

