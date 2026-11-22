create database goodsapi_bench owner prep;

insert into goodcategory(name) select 'cat' || g from generate_series(1,100) g;

insert into good(name,price,id_goodcategory)
select 'good' || g,
       (random() * 1000 + 1)::numeric(10,2),
    (random() * 99 + 1)::int
from generate_series(1,1000000) g;

ANALYSE good;

select count(*) from good;

explain analyse select * from good where id_goodcategory=5;

explain analyse select * from good where name='good500000';
create index good_name_idx on good(name);

explain analyse select * from good where lower(name)='good500000';
create index good_lower_name_idx on good(lower(name));

explain analyse select * from good where name LIKE 'good500000%';

CREATE INDEX good_name_pattern_idx ON good(name text_pattern_ops);
EXPLAIN ANALYZE SELECT * FROM good WHERE name LIKE 'good50000%';  -- теперь индекс
EXPLAIN ANALYZE SELECT * FROM good WHERE name LIKE '%50000';      -- Seq Scan при любом индексе
CREATE INDEX GOOD_CAT_Price_idx on good(id_goodcategory,price);
Explain analyse select * from good where id_goodcategory=5 and price>900;
Explain analyse select * from good where id_goodcategory=5 order by price;
VACUUM good;
EXPLAIN ANALYZE SELECT id_goodcategory, price FROM good WHERE id_goodcategory = 5;

CREATE INDEX good_price_idx ON good(price);
EXPLAIN ANALYZE SELECT * FROM good WHERE price > 995;  -- около 0.5% строк: индекс
EXPLAIN ANALYZE SELECT * FROM good WHERE price > 10;