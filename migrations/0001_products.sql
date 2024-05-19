-- +migrate Up
CREATE TABLE IF NOT EXISTS `products`
(
    id          serial primary key,
    name        varchar(100)    not null,
    description text            not null,
    price       bigint unsigned not null,
    online      tinyint         not null default 0,
    created_at   datetime        not null default now(),
    modified_at  datetime                 default null
);

-- +migrate StatementBegin
DROP PROCEDURE IF EXISTS `createProduct`;
CREATE PROCEDURE `createProduct`(
    IN _name varchar(50),
    IN _description text,
    IN _price bigint unsigned
)
BEGIN
    START TRANSACTION;

    insert into `products` (name, description, price) values (_name, _description, _price);
    select id, name, description, price, online, created_at, modified_at from `products` where id = LAST_INSERT_ID() limit 1;
    COMMIT;
END;
-- +migrate StatementEnd

