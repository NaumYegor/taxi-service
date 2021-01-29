-- +migrate Up
create table roles (
    id bigserial primary key,
    name varchar(15) not null,
    description text
);

create table users (
    id bigserial primary key,
    nickname varchar(20) unique not null,
    password bytea not null,
    token varchar(64),
    role_id bigserial not null
);

create table cars (
    id bigserial primary key,
    driver bigserial references users(id),
    model varchar(25) not null,
    number varchar(25) unique not null
);
ALTER TABLE cars ALTER COLUMN driver DROP NOT NULL;

create table orders (
    id bigserial primary key,
    client_id bigserial references users(id) not null,
    car_id bigserial references cars(id),
    completed boolean not null DEFAULT false,
    info text not null
);
ALTER TABLE orders ALTER COLUMN car_id DROP NOT NULL;

insert into roles (name)
values
('client'),
('driver'),
('admin');
