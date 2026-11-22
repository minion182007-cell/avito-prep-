insert into Goodcategory (name)
SELECT 'category' || g  from generate_series(1,20) g;

SELECT id, name FROM goodcategory ORDER BY id;

INSERT INTO good(name, price, id_goodcategory)
select 'good' || g,
       (random() * 1000)::numeric(10,2),
       1+(g%20)
from generate_series(1,200000) g;

SELECT count(*) FROM good;

SELECT id_goodCategory, count(*) FROM good GROUP BY id_goodCategory ORDER BY 1;

ANALYZE good;

EXPLAIN ANALYSE SELECT * FROM good where name='good150000';

create index good_name_idx on good(name);

Select pg_size_pretty(pg_relation_size('good_name_idx'));
Select pg_size_pretty(pg_relation_size('good'));

EXPLAIN ANALYSE SELECT * from good where id_goodcategory=5 and price>900;

CREATE INDEX good_cat_price_idx on good(id_goodcategory,price);

EXPLAIN ANALYZE SELECT * FROM good WHERE price > 900;

DROP INDEX good_name_idx;
DROP INDEX good_cat_price_idx;

SELECT indexname FROM pg_indexes WHERE tablename = 'good';

