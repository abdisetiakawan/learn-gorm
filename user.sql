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