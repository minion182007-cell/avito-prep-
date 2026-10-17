CREATE TABLE GoodCategory
(
    id     int primary key check (id > 0),
    name   text not null,
    ageMin int check ( ageMin > 0 )
);
CREATE TABLE Good
(
    id              int primary key check (id > 0),
    name            text not null,
    productionDate  date,
    price           numeric(10, 2),
    weight          float,
    id_goodCategory int  references GoodCategory (id)
                             ON UPDATE CASCADE
                             ON DELETE SET NULL

);
CREATE TABLE Shop
(
    id      int primary key check (id > 0),
    name    text not null,
    address text
);
CREATE TABLE Good_Shop
(
    id_good int references Good (id)
        ON DELETE CASCADE,
    id_shop int references Shop (id)
        ON DELETE CASCADE,
    count   int check (count >= 0),
    primary key (id_good, id_shop)
);
INSERT INTO GoodCategory(id, name, ageMin)
VALUES (1, 'Электроника', 16),
       (2, 'Мебель', 16),
       (3, 'Аксессуары', NULL);

INSERT INTO Good(id, name, productionDate, price, weight, id_goodCategory)
VALUES (1, 'Apple iPhone 14', '2022-09-01', 140000, 0.7, 1),
       (2, 'Чехол на телефон 6.5', '2022-09-01', 999.99, 0.1, 3),
       (3, 'Стол компьютерный', '2022-06-02', 4500, 14.7, 2);
INSERT INTO Shop(id, name, address)
VALUES (1, 'МВидео', 'ул. Горького, 165'),
       (2, 'Эльдорадо', 'ул. Тимирязева, 22');
INSERT INTO Good_Shop(id_good, id_shop, count)
VALUES (1, 1, 3),
       (1, 2, 1),
       (2, 1, 10),
       (2, 2, 4),
       (3, 1, 2);

SELECT * FROM Good_Shop;

SELECT * FROM Good_Shop WHERE id_shop=2;
SELECT * FROM Good_Shop WHERE count>3;
SELECT * FROM good ORDER BY price DESC ;
SELECT Good.name,Good.price FROM good;
SELECT ID,NAME FROM goodcategory WHERE ageMin IS NULL ;
UPDATE good SET price=price * 1.1 WHERE id=2;
SELECT Good.price FROM good WHERE id=2;
DELETE FROM goodcategory WHERE id=2;
SELECT * FROM Good;
DELETE FROM good WHERE id=1;
SELECT * FROM Good_Shop;
DELETE FROM shop where id=1;
SELECT * FROM Good_Shop;

DROP TABLE Good_Shop;
DROP TABLE Good;
DROP TABLE Shop;
DROP TABLE GoodCategory;

insert into good(id, name,price) values  (4, 'Новый товар', 500);
SELECT * FROM Good;
SELECT s.name, gs.id_good, gs.count
FROM Good_Shop gs
         JOIN Shop s ON s.id = gs.id_shop;

select g.name, gc.name from goodcategory gc  join Good g on gc.id = g.id_goodCategory;
select g.name, gc.name from goodcategory gc  right join   Good g on gc.id = g.id_goodCategory;
select g.name, gc.name from goodcategory gc  join Good g on gc.id = g.id_goodCategory;
select g.name, gc.name from Good g  left join goodcategory gc  on gc.id = g.id_goodCategory;

select g.name, s.name, gs.count
from Good_Shop gs
         join Good g on g.id = gs.id_good
         join Shop s on s.id = gs.id_shop;

select g.name, s.name, gs.count, gc.name
from Good_Shop gs
         join goodcategory gc on gc.id = gs.id_good
         join Good g on g.id = gs.id_good

         join Shop s on s.id = gs.id_shop;
select g.name,gc.name, s.name, gs.count
from Good_Shop gs
         join Good g on g.id = gs.id_good
         join goodcategory gc on gc.id = g.id_goodCategory
         join Shop s on s.id = gs.id_shop;
select g.name, s.name, gs.count
from good g
         left join  Good_Shop GS on g.id = GS.id_good and gs.id_shop=1
         left join  Shop S on S.id = GS.id_shop;

select g.name, gs.count
from good g
         left join  Good_Shop GS on g.id = GS.id_good where gs.id_good is null