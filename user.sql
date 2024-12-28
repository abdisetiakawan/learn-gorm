-- Active: 1724149523236@@127.0.0.1@3306@perpustakaan
create table users
(
    id         varchar(100) not null,
    password   varchar(100) not null,
    name       varchar(100) not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;

select *
from users;

alter table users
    rename column name to first_name;

alter table users
    add column middle_name varchar(100) null after first_name;

alter table users
    add column last_name varchar(100) null after middle_name;

select *
from users; 

create table user_logs
(
    id          INT          AUTO_INCREMENT,
    user_id     varchar(100) not null,
    action      varchar(100) not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;

SELECT * FROM user_logs;
SHOW DATABASES;

DELETE FROM user_logs;

DELETE FROM user_logs;

ALTER TABLE user_logs
    MODIFY updated_at BIGINT not null;
ALTER TABLE user_logs
    MODIFY created_at BIGINT not null;

SELECT * FROM user_logs;

DESC user_logs;

CREATE TABLE wallets
(
    id VARCHAR(100) NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    Foreign Key (user_id) REFERENCES users(id)
)

SELECT * FROM wallets;

create table addresses
(
    id         bigint       not null auto_increment,
    user_id    varchar(100) not null,
    address    varchar(100) not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id) references users (id)
) engine = innodb;

desc addresses;

create table products
(
    id         varchar(100) not null,
    name       varchar(100) not null,
    price      bigint       not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = innodb;

desc products;

create table user_like_product
(
    user_id    varchar(100) not null,
    product_id varchar(100) not null,
    primary key (user_id, product_id),
    foreign key (user_id) references users (id),
    foreign key (product_id) references products (id)
) engine = innodb;

desc user_like_product;

select * from addresses;

SELECT * FROM `wallets` WHERE id = '2' LIMIT 1;

SELECT * FROM `users` WHERE `users`.`id` = '2';

SELECT * FROM `addresses` WHERE `addresses`.`user_id` = '2';

select * from wallets;

select count(id) from users;