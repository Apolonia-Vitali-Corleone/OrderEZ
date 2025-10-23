create database if not exists order_ez default character set utf8mb4 collate utf8mb4_unicode_ci default encryption = 'N';

use order_ez;

drop table if exists oe_order;
create table oe_order
(
    order_id    bigint not null,
    user_id     bigint not null,
    total_price int    not null,
    primary key (order_id)
) engine = innodb
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;

drop table if exists oe_order_item;
create table oe_order_item
(
    order_item_id bigint       not null,
    order_id      bigint       not null,
    user_id       bigint       not null,
    item_id       bigint       not null,
    item_name     varchar(255) not null,
    item_price    int          not null,
    item_count    int          not null,
    primary key (order_item_id)
) engine = innodb
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;

drop table if exists oe_user;
create table oe_user
(
    user_id  bigint       not null auto_increment,
    username varchar(64)  not null,
    password varchar(255) not null,
    primary key (user_id),
    unique key username (username)
) engine = innodb
  auto_increment = 31
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;